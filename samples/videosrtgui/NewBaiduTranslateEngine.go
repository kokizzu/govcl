// Automatically generated by the res2go, do not edit.
package main

import (
    "github.com/ying32/govcl/vcl"
)

type TNewBaiduTranslateEngine struct {
    *vcl.TForm
    Label1        *vcl.TLabel
    Label2        *vcl.TLabel
    Label3        *vcl.TLabel
    Label4        *vcl.TLabel
    Label5        *vcl.TLabel
    EdtName       *vcl.TEdit
    EdtAppId      *vcl.TEdit
    EditAppSecret *vcl.TEdit
    CbbType       *vcl.TComboBox
    BtnOK         *vcl.TButton
    Button2       *vcl.TButton

    //::private::
    TNewBaiduTranslateEngineFields
}

var NewBaiduTranslateEngine *TNewBaiduTranslateEngine




// vcl.Application.CreateForm(&NewBaiduTranslateEngine)

func NewNewBaiduTranslateEngine(owner vcl.IComponent) (root *TNewBaiduTranslateEngine)  {
    vcl.CreateResForm(owner, &root)
    return
}

var newBaiduTranslateEngineBytes = []byte("\x54\x50\x46\x30\x18\x54\x4E\x65\x77\x42\x61\x69\x64\x75\x54\x72\x61\x6E\x73\x6C\x61\x74\x65\x45\x6E\x67\x69\x6E\x65\x17\x4E\x65\x77\x42\x61\x69\x64\x75\x54\x72\x61\x6E\x73\x6C\x61\x74\x65\x45\x6E\x67\x69\x6E\x65\x04\x4C\x65\x66\x74\x03\xBD\x03\x06\x48\x65\x69\x67\x68\x74\x03\x07\x01\x03\x54\x6F\x70\x03\xC7\x01\x05\x57\x69\x64\x74\x68\x03\xDF\x01\x0B\x42\x6F\x72\x64\x65\x72\x49\x63\x6F\x6E\x73\x0B\x0C\x62\x69\x53\x79\x73\x74\x65\x6D\x4D\x65\x6E\x75\x00\x0B\x42\x6F\x72\x64\x65\x72\x53\x74\x79\x6C\x65\x07\x08\x62\x73\x53\x69\x6E\x67\x6C\x65\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x24\xE6\x96\xB0\xE5\xBB\xBA\xE7\xBF\xBB\xE8\xAF\x91\xE5\xBC\x95\xE6\x93\x8E\xEF\xBC\x88\xE7\x99\xBE\xE5\xBA\xA6\xE7\xBF\xBB\xE8\xAF\x91\xEF\xBC\x89\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x03\x07\x01\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\xDF\x01\x08\x50\x6F\x73\x69\x74\x69\x6F\x6E\x07\x0E\x70\x6F\x53\x63\x72\x65\x65\x6E\x43\x65\x6E\x74\x65\x72\x0A\x4C\x43\x4C\x56\x65\x72\x73\x69\x6F\x6E\x06\x08\x32\x2E\x30\x2E\x31\x30\x2E\x30\x00\x06\x54\x4C\x61\x62\x65\x6C\x06\x4C\x61\x62\x65\x6C\x31\x04\x4C\x65\x66\x74\x02\x16\x06\x48\x65\x69\x67\x68\x74\x02\x11\x03\x54\x6F\x70\x02\x18\x05\x57\x69\x64\x74\x68\x02\x41\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x10\xE5\x90\x8D\xE7\xA7\xB0\x2F\xE5\x88\xAB\xE5\x90\x8D\xEF\xBC\x9A\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x00\x00\x06\x54\x4C\x61\x62\x65\x6C\x06\x4C\x61\x62\x65\x6C\x32\x04\x4C\x65\x66\x74\x02\x16\x06\x48\x65\x69\x67\x68\x74\x02\x11\x03\x54\x6F\x70\x02\x39\x05\x57\x69\x64\x74\x68\x02\x30\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x08\x41\x70\x70\x49\x64\xEF\xBC\x9A\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x00\x00\x06\x54\x4C\x61\x62\x65\x6C\x06\x4C\x61\x62\x65\x6C\x33\x04\x4C\x65\x66\x74\x02\x16\x06\x48\x65\x69\x67\x68\x74\x02\x11\x03\x54\x6F\x70\x02\x5B\x05\x57\x69\x64\x74\x68\x02\x48\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x0C\x41\x70\x70\x53\x65\x63\x72\x65\x74\xEF\xBC\x9A\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x00\x00\x06\x54\x4C\x61\x62\x65\x6C\x06\x4C\x61\x62\x65\x6C\x34\x04\x4C\x65\x66\x74\x02\x16\x06\x48\x65\x69\x67\x68\x74\x02\x11\x03\x54\x6F\x70\x02\x7E\x05\x57\x69\x64\x74\x68\x02\x54\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x15\xE5\xB8\x90\xE5\x8F\xB7\xE8\xAE\xA4\xE8\xAF\x81\xE7\xB1\xBB\xE5\x9E\x8B\xEF\xBC\x9A\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x00\x00\x06\x54\x4C\x61\x62\x65\x6C\x06\x4C\x61\x62\x65\x6C\x35\x04\x4C\x65\x66\x74\x02\x16\x06\x48\x65\x69\x67\x68\x74\x02\x22\x03\x54\x6F\x70\x03\x9B\x00\x05\x57\x69\x64\x74\x68\x03\x92\x01\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x7A\xE8\xAF\xB4\xE6\x98\x8E\xEF\xBC\x9A\x0D\x0A\xE8\xAF\xB7\xE5\xA1\xAB\xE5\x86\x99\xE5\x9C\xA8\xE2\x80\x9C\xE7\x99\xBE\xE5\xBA\xA6\xE7\xBF\xBB\xE8\xAF\x91\xE5\xBC\x80\xE6\x94\xBE\xE5\xB9\xB3\xE5\x8F\xB0\xE2\x80\x9D\xE7\x94\xB3\xE8\xAF\xB7\xE7\x9A\x84\xE5\xAF\x86\xE9\x92\xA5\xEF\xBC\x8C\xE6\xB3\xA8\xE6\x84\x8F\xE2\x80\x9C\xE6\xA0\x87\xE5\x87\x86\xE7\x89\x88\xE2\x80\x9D\xE5\x92\x8C\xE2\x80\x9C\xE9\xAB\x98\xE7\xBA\xA7\xE7\x89\x88\xE2\x80\x9D\xE7\x9A\x84\xE9\x80\x89\xE6\x8B\xA9\xE3\x80\x82\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x00\x00\x05\x54\x45\x64\x69\x74\x07\x45\x64\x74\x4E\x61\x6D\x65\x04\x4C\x65\x66\x74\x02\x6F\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x02\x15\x05\x57\x69\x64\x74\x68\x03\x61\x01\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x00\x00\x05\x54\x45\x64\x69\x74\x08\x45\x64\x74\x41\x70\x70\x49\x64\x04\x4C\x65\x66\x74\x02\x6F\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x02\x34\x05\x57\x69\x64\x74\x68\x03\x61\x01\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x01\x00\x00\x05\x54\x45\x64\x69\x74\x0D\x45\x64\x69\x74\x41\x70\x70\x53\x65\x63\x72\x65\x74\x04\x4C\x65\x66\x74\x02\x6F\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x02\x56\x05\x57\x69\x64\x74\x68\x03\x61\x01\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x02\x00\x00\x09\x54\x43\x6F\x6D\x62\x6F\x42\x6F\x78\x07\x43\x62\x62\x54\x79\x70\x65\x04\x4C\x65\x66\x74\x02\x6F\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x02\x79\x05\x57\x69\x64\x74\x68\x03\x61\x01\x0A\x49\x74\x65\x6D\x48\x65\x69\x67\x68\x74\x02\x11\x09\x49\x74\x65\x6D\x49\x6E\x64\x65\x78\x02\x00\x0D\x49\x74\x65\x6D\x73\x2E\x53\x74\x72\x69\x6E\x67\x73\x01\x06\x09\xE6\xA0\x87\xE5\x87\x86\xE7\x89\x88\x06\x09\xE9\xAB\x98\xE7\xBA\xA7\xE7\x89\x88\x00\x05\x53\x74\x79\x6C\x65\x07\x0E\x63\x73\x44\x72\x6F\x70\x44\x6F\x77\x6E\x4C\x69\x73\x74\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x03\x04\x54\x65\x78\x74\x06\x09\xE6\xA0\x87\xE5\x87\x86\xE7\x89\x88\x00\x00\x07\x54\x42\x75\x74\x74\x6F\x6E\x05\x42\x74\x6E\x4F\x4B\x04\x4C\x65\x66\x74\x03\x31\x01\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x03\xDE\x00\x05\x57\x69\x64\x74\x68\x02\x4B\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x0C\xE7\xA1\xAE\xE5\xAE\x9A\xE6\x96\xB0\xE5\xA2\x9E\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x04\x00\x00\x07\x54\x42\x75\x74\x74\x6F\x6E\x07\x42\x75\x74\x74\x6F\x6E\x32\x04\x4C\x65\x66\x74\x03\x84\x01\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x03\xDE\x00\x05\x57\x69\x64\x74\x68\x02\x4B\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE5\x8F\x96\xE6\xB6\x88\x0B\x4D\x6F\x64\x61\x6C\x52\x65\x73\x75\x6C\x74\x02\x02\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x05\x00\x00\x00")

// Register Form Resources
var _ = vcl.RegisterFormResource(NewBaiduTranslateEngine, &newBaiduTranslateEngineBytes)
