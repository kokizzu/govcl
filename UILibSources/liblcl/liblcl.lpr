
//----------------------------------------
// The code is automatically generated by the GenlibLcl tool.
// Copyright © ying32. All Rights Reserved.
// 
// Licensed under Apache License 2.0
//
//----------------------------------------

library liblcl;

//{$mode objfpc}{$H+}
{$mode delphi}

{$ifndef windows}
  {$define UseCThreads}
{$endif}

uses
{$IFDEF UNIX}{$IFDEF UseCThreads}
  cthreads,
{$ENDIF}{$ENDIF}
{$I UseAll.inc}
  ,Classes,
  SysUtils,
  uFormDesignerFile,
{$IFDEF LCLCocoa}
  uMacOSPatchs,
{$ENDIF}
{$IFDEF LINUX}
  uLinuxPatchs,
{$ENDIF}
  uComponents,
  uControlPatchs,
  uExport1,
  uExport2;

{$IFDEF WINDOWS}
  {$R *.res}
{$ENDIF}

{$I ExtDecl.inc}
{$I LazarusDef.inc}

// 用户自己定义的组件
{$I UserDefineComponents.inc}

begin
  InitLazarusDef;
end.