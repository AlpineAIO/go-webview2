//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

type ICoreWebView2ContextMenuTargetVtbl struct {
	IUnknownVtbl
	GetKind                    ComProc
	GetIsEditable              ComProc
	GetIsRequestedForMainFrame ComProc
	GetPageUri                 ComProc
	GetFrameUri                ComProc
	GetHasLinkUri              ComProc
	GetLinkUri                 ComProc
	GetHasLinkText             ComProc
	GetLinkText                ComProc
	GetHasSourceUri            ComProc
	GetSourceUri               ComProc
	GetHasSelection            ComProc
	GetSelectionText           ComProc
}

type ICoreWebView2ContextMenuTarget struct {
	Vtbl *ICoreWebView2ContextMenuTargetVtbl
}

func (i *ICoreWebView2ContextMenuTarget) AddRef() uintptr {
	refCounter, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return refCounter
}

func (i *ICoreWebView2ContextMenuTarget) GetKind() (*COREWEBVIEW2_CONTEXT_MENU_TARGET_KIND, error) {

	var value COREWEBVIEW2_CONTEXT_MENU_TARGET_KIND

	hr, _, err := i.Vtbl.GetKind.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return nil, syscall.Errno(hr)
	}
	return &value, err
}

func (i *ICoreWebView2ContextMenuTarget) GetIsEditable() (*bool, error) {
	// Create int32 to hold bool result
	var _value int32

	hr, _, err := i.Vtbl.GetIsEditable.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_value)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return nil, syscall.Errno(hr)
	} // Get result and cleanup
	value := _value != 0
	return &value, err
}

func (i *ICoreWebView2ContextMenuTarget) GetIsRequestedForMainFrame() (*bool, error) {
	// Create int32 to hold bool result
	var _value int32

	hr, _, err := i.Vtbl.GetIsRequestedForMainFrame.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_value)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return nil, syscall.Errno(hr)
	} // Get result and cleanup
	value := _value != 0
	return &value, err
}

func (i *ICoreWebView2ContextMenuTarget) GetPageUri() (*string, error) {
	// Create *uint16 to hold result
	var _value *uint16

	hr, _, err := i.Vtbl.GetPageUri.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_value)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return nil, syscall.Errno(hr)
	} // Get result and cleanup
	value := UTF16PtrToString(_value)
	CoTaskMemFree(unsafe.Pointer(_value))
	return &value, err
}

func (i *ICoreWebView2ContextMenuTarget) GetFrameUri() (*string, error) {
	// Create *uint16 to hold result
	var _value *uint16

	hr, _, err := i.Vtbl.GetFrameUri.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_value)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return nil, syscall.Errno(hr)
	} // Get result and cleanup
	value := UTF16PtrToString(_value)
	CoTaskMemFree(unsafe.Pointer(_value))
	return &value, err
}

func (i *ICoreWebView2ContextMenuTarget) GetHasLinkUri() (*bool, error) {
	// Create int32 to hold bool result
	var _value int32

	hr, _, err := i.Vtbl.GetHasLinkUri.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_value)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return nil, syscall.Errno(hr)
	} // Get result and cleanup
	value := _value != 0
	return &value, err
}

func (i *ICoreWebView2ContextMenuTarget) GetLinkUri() (*string, error) {
	// Create *uint16 to hold result
	var _value *uint16

	hr, _, err := i.Vtbl.GetLinkUri.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_value)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return nil, syscall.Errno(hr)
	} // Get result and cleanup
	value := UTF16PtrToString(_value)
	CoTaskMemFree(unsafe.Pointer(_value))
	return &value, err
}

func (i *ICoreWebView2ContextMenuTarget) GetHasLinkText() (*bool, error) {
	// Create int32 to hold bool result
	var _value int32

	hr, _, err := i.Vtbl.GetHasLinkText.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_value)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return nil, syscall.Errno(hr)
	} // Get result and cleanup
	value := _value != 0
	return &value, err
}

func (i *ICoreWebView2ContextMenuTarget) GetLinkText() (*string, error) {
	// Create *uint16 to hold result
	var _value *uint16

	hr, _, err := i.Vtbl.GetLinkText.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_value)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return nil, syscall.Errno(hr)
	} // Get result and cleanup
	value := UTF16PtrToString(_value)
	CoTaskMemFree(unsafe.Pointer(_value))
	return &value, err
}

func (i *ICoreWebView2ContextMenuTarget) GetHasSourceUri() (*bool, error) {
	// Create int32 to hold bool result
	var _value int32

	hr, _, err := i.Vtbl.GetHasSourceUri.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_value)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return nil, syscall.Errno(hr)
	} // Get result and cleanup
	value := _value != 0
	return &value, err
}

func (i *ICoreWebView2ContextMenuTarget) GetSourceUri() (*string, error) {
	// Create *uint16 to hold result
	var _value *uint16

	hr, _, err := i.Vtbl.GetSourceUri.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_value)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return nil, syscall.Errno(hr)
	} // Get result and cleanup
	value := UTF16PtrToString(_value)
	CoTaskMemFree(unsafe.Pointer(_value))
	return &value, err
}

func (i *ICoreWebView2ContextMenuTarget) GetHasSelection() (*bool, error) {
	// Create int32 to hold bool result
	var _value int32

	hr, _, err := i.Vtbl.GetHasSelection.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_value)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return nil, syscall.Errno(hr)
	} // Get result and cleanup
	value := _value != 0
	return &value, err
}

func (i *ICoreWebView2ContextMenuTarget) GetSelectionText() (*string, error) {
	// Create *uint16 to hold result
	var _value *uint16

	hr, _, err := i.Vtbl.GetSelectionText.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_value)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return nil, syscall.Errno(hr)
	} // Get result and cleanup
	value := UTF16PtrToString(_value)
	CoTaskMemFree(unsafe.Pointer(_value))
	return &value, err
}
