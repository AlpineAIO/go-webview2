//go:build windows

package edge

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

type iCoreWebView2_20Vtbl struct {
	_IUnknownVtbl
	GetFrameId ComProc
}

type ICoreWebView2_20 struct {
	vtbl *iCoreWebView2_20Vtbl
}

func (i *ICoreWebView2_20) GetFrameId() (uint32, error) {
	var result uint32

	_, _, err := i.vtbl.GetFrameId.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&result)))
	if err != windows.ERROR_SUCCESS {
		return 0, err
	}

	return result, nil
}

func (i *ICoreWebView2) GetICoreWebView2_20() *ICoreWebView2_20 {
	var result *ICoreWebView2_20

	iidICoreWebView2_20 := NewGUID("{b4bc1926-7305-11ee-b962-0242ac120002}")
	_, _, _ = i.vtbl.QueryInterface.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(iidICoreWebView2_20)),
		uintptr(unsafe.Pointer(&result)))

	return result
}

func (e *Chromium) GetICoreWebView2_20() *ICoreWebView2_20 {
	return e.webview.GetICoreWebView2_20()
}
