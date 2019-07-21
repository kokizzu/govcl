//----------------------------------------
// 加载文件或者内存中的窗口资源文件功能
// 需要配合窗口设计器使用。
// 设计器是独立于govcl的，设计器的目的是用于简化窗口的创建，设计器不开源。
//
// 例：
//    type TMainForm struct {
//        *vcl.TForm  // 不要写名称，使用隐式的
//        Button1 *vcl.TButton // 注意，设计器中一定要首字母大写，否则不能使用反射填充
//    }
//
//	  var mainForm *TMainForm
//
//    func main() {
//        ...
//        // 文件加载方式
//        vcl.Application.CreateFormFromFile(filename, &mainForm)
//        // 流加载方式
//        // mem := vcl.NewMemoryStream()
//        // mem.Write([]byte(...))
//        // vcl.Application.CreateFormFromStream(mem, &mainForm)
//        // mem.Free()
//        // 资源加载方式
//        //vcl.Application.CreateFormFromResourceName("ResName", &mainForm)
//        ...
//     }

// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------

package vcl

import (
	"fmt"
	"reflect"
	"unsafe"

	. "github.com/ying32/govcl/vcl/api"
)

func setFiledVal(name string, instance uintptr, v reflect.Value) {
	field := v.FieldByName(name)
	if field.IsValid() && field.CanSet() {
		// 获取这个字段的类型  field.Type() 为指针类，后面再使用Elem()后返回的为非指针类型的
		fv := reflect.New(field.Type().Elem())
		// 这里用循环去找 instance实例是因为防止以后加字段用，一般情况索引为1，先不使用循环吧。
		//for j := 0; j < fv.Elem().Type().NumField(); j++ {
		//	if fv.Elem().Type().Field(j).Name == "instance" {
		// 因为反射不能设置未导出的，所以直接使用指针来设置
		setVal := func(idx int, value uintptr) {
			*(*uintptr)(unsafe.Pointer(fv.Elem().Field(idx).UnsafeAddr())) = value
		}
		// idx = 0 = TForm
		setVal(1, instance) // idx = 1 = instance
		setVal(2, instance) // idx = 2 = ptr
		// instance ord = 1
		//*(*uintptr)(unsafe.Pointer(fv.Elem().Field(1).UnsafeAddr())) = instance
		// ptr ord = 2
		//*(*uintptr)(unsafe.Pointer(fv.Elem().Field(2).UnsafeAddr())) = instance

		//		break
		//	}
		//}
		// 将这个实例填充到字段中
		field.Set(fv)
	}
}

// fullFiledVal 动态创设置字段值
func fullFiledVal(f IComponent, out interface{}, fullSubComponent, afterBindSubComponentsEvents bool) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("err:", err)
		}
	}()
	// out是一个 **TXXForm的变量指针，未进行分配内存，表现形式为 **TXXX，每使用一个Elem()减少一个
	vt := reflect.TypeOf(out).Elem()
	v := reflect.New(vt.Elem())
	// 将实例化后的值填充到out指针变量中，这里要能修改的需要使用Elem()方法获取
	reflect.ValueOf(out).Elem().Set(v)
	// 获取指针类型
	vPtr := v.Elem()
	// 检查是否有效，并且可以被设置
	if vPtr.IsValid() && vPtr.CanSet() {
		// 如果没有名字，就指定一个名字，名字以当前类，如果首个为T则去除
		if f.Name() == "" {
			newName := vPtr.Type().Name()
			if newName[0] == 'T' {
				newName = newName[1:]
			}
			f.SetName(newName)
		}
		// TForm，默认的, 使用隐式嵌入
		setFiledVal("TForm", f.Instance(), vPtr)
		// fullComponent 只有当是从资源加载的才进行填充操作。
		if fullSubComponent {
			var ci int32
			for ci = 0; ci < f.ComponentCount(); ci++ {
				// 通过资源文件中的组件名字查找
				setFiledVal(f.Components(ci).Name(), f.Components(ci).Instance(), vPtr)
			}
		}
		// 关联事件
		autoBindEvents(v, f, fullSubComponent, afterBindSubComponentsEvents)
	}
}

// 共用的一个从资源中加载构建对象
func resObjtBuild(typ int, owner IComponent, appInst uintptr, fields ...interface{}) IComponent {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("resCreateForm Error: ", err)
		}
	}()

	var fullSubComponent bool
	var afterBindSubComponentsEvents bool
	var field1 interface{}

	// 初始创建时是否使用缩放
	initScale := len(fields) != 2
	if len(fields) >= 2 {
		switch fields[1].(type) {
		case bool:
			initScale = true
		default:
			initScale = false
		}
	}

	var resObj IComponent

	switch typ {
	case 0:
		// 由参数的个数决定，创建窗口时是否使用缩放，此值需要 vcl.Application.SetFormScaled(true) 后才能生效。
		resObj = FormFromInst(Application_CreateForm(appInst, initScale))
	case 1:
		resObj = FormFromInst(Form_Create2(CheckPtr(owner), initScale))
	case 2:
		// TFrame
		resObj = NewFrame(owner)
	}

	switch len(fields) {
	case 1:
		field1 = fields[0]
		fullSubComponent = false
		afterBindSubComponentsEvents = false

	case 2:
		switch fields[1].(type) {
		// 当第二个参数为bool时，表示不填充子组件，为true表示之后绑定事件
		case bool:
			field1 = fields[0]
			fullSubComponent = false
			afterBindSubComponentsEvents = fields[1].(bool)
		default:
			// 第二个参数类型不为bool时，填充子组件为true，之后绑定事件为false
			field1 = fields[1]
			fullSubComponent = true
			afterBindSubComponentsEvents = false
			switch fields[0].(type) {
			case string:
				ResFormLoadFromFile(fields[0].(string), CheckPtr(resObj))
			case []byte:
				mem := NewMemoryStreamFromBytes(fields[0].([]byte))
				defer mem.Free()
				ResFormLoadFromStream(CheckPtr(mem), CheckPtr(resObj))
			}
		}
	default:
		return resObj
	}
	fullFiledVal(resObj, field1, fullSubComponent, afterBindSubComponentsEvents)
	return nil
}