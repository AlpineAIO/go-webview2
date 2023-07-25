//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

type ICoreWebView2FrameInfoVtbl struct {
	IUnknownVtbl
	GetName   ComProc
	GetSource ComProc
}

type ICoreWebView2FrameInfo struct {
	Vtbl *ICoreWebView2FrameInfoVtbl
}

func (i *ICoreWebView2FrameInfo) AddRef() uintptr {
	refCounter, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return refCounter
}

func (i *ICoreWebView2FrameInfo) GetName() (*string, error) {
	// Create *uint16 to hold result
	var _name *uint16

	hr, _, err := i.Vtbl.GetName.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_name)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return nil, syscall.Errno(hr)
	} // Get result and cleanup
	name := UTF16PtrToString(_name)
	CoTaskMemFree(unsafe.Pointer(_name))
	return &name, err
}

func (i *ICoreWebView2FrameInfo) GetSource() (*string, error) {
	// Create *uint16 to hold result
	var _source *uint16

	hr, _, err := i.Vtbl.GetSource.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_source)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return nil, syscall.Errno(hr)
	} // Get result and cleanup
	source := UTF16PtrToString(_source)
	CoTaskMemFree(unsafe.Pointer(_source))
	return &source, err
}
