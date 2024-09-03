//go:build windows

package edge

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

type iCoreWebView2_22Vtbl struct {
	_IUnknownVtbl
	AddWebResourceRequestedFilterWithRequestSourceKinds    ComProc
	RemoveWebResourceRequestedFilterWithRequestSourceKinds ComProc
}

type ICoreWebView2_22 struct {
	vtbl *iCoreWebView2_22Vtbl
}

func (i *ICoreWebView2_22) AddWebResourceRequestedFilterWithRequestSourceKinds(uri string, resourceContext COREWEBVIEW2_WEB_RESOURCE_CONTEXT, resourceKinds COREWEBVIEW2_WEB_RESOURCE_REQUEST_SOURCE_KINDS) error {
	_uri, _ := windows.UTF16PtrFromString(uri)

	_, _, err := i.vtbl.AddWebResourceRequestedFilterWithRequestSourceKinds.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_uri)),
		uintptr(resourceContext),
		uintptr(resourceKinds),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}

	return nil
}

func (i *ICoreWebView2_22) RemoveWebResourceRequestedFilterWithRequestSourceKinds(uri string, resourceContext COREWEBVIEW2_WEB_RESOURCE_CONTEXT, resourceKinds COREWEBVIEW2_WEB_RESOURCE_REQUEST_SOURCE_KINDS) error {
	_uri, _ := windows.UTF16PtrFromString(uri)

	_, _, err := i.vtbl.RemoveWebResourceRequestedFilterWithRequestSourceKinds.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_uri)),
		uintptr(resourceContext),
		uintptr(unsafe.Pointer(&resourceKinds)),
	)
	if err != windows.ERROR_SUCCESS {
		return err
	}

	return nil
}

func (i *ICoreWebView2_22) AddRef() uintptr {
	refCounter, _, _ := i.vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return refCounter
}

func (i *ICoreWebView2) GetICoreWebView2_22() *ICoreWebView2_22 {
	var result *ICoreWebView2_22

	iidICoreWebView2_22 := NewGUID("{DB75DFC7-A857-4632-A398-6969DDE26C0A}")
	_, _, _ = i.vtbl.QueryInterface.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(iidICoreWebView2_22)),
		uintptr(unsafe.Pointer(&result)))

	return result
}

func (e *Chromium) GetICoreWebView2_22() *ICoreWebView2_22 {
	return e.webview.GetICoreWebView2_22()
}
