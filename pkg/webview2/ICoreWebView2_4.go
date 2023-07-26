//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

type ICoreWebView2_4Vtbl struct {
	IUnknownVtbl
	AddFrameCreated        ComProc
	RemoveFrameCreated     ComProc
	AddDownloadStarting    ComProc
	RemoveDownloadStarting ComProc
}

type ICoreWebView2_4 struct {
	Vtbl *ICoreWebView2_4Vtbl
}

func (i *ICoreWebView2_4) AddRef() uintptr {
	refCounter, _, _ := i.Vtbl.AddRef.Call(uintptr(unsafe.Pointer(i)))
	return refCounter
}

func (i *ICoreWebView2) GetICoreWebView2_4() *ICoreWebView2_4 {
	var result *ICoreWebView2_4

	iidICoreWebView2_4 := NewGUID("{20d02d59-6df2-42dc-bd06-f98a694b1302}")
	_, _, _ = i.Vtbl.QueryInterface.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(iidICoreWebView2_4)),
		uintptr(unsafe.Pointer(&result)))

	return result
}

func (i *ICoreWebView2_4) AddFrameCreated(eventHandler *ICoreWebView2FrameCreatedEventHandler) (*EventRegistrationToken, error) {

	var token *EventRegistrationToken

	hr, _, err := i.Vtbl.AddFrameCreated.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(&token)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return nil, syscall.Errno(hr)
	}
	return token, err
}

func (i *ICoreWebView2_4) RemoveFrameCreated(token EventRegistrationToken) error {

	hr, _, err := i.Vtbl.RemoveFrameCreated.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&token)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return syscall.Errno(hr)
	}
	return err
}

func (i *ICoreWebView2_4) AddDownloadStarting(eventHandler *ICoreWebView2DownloadStartingEventHandler) (*EventRegistrationToken, error) {

	var token *EventRegistrationToken

	hr, _, err := i.Vtbl.AddDownloadStarting.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(eventHandler)),
		uintptr(unsafe.Pointer(&token)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return nil, syscall.Errno(hr)
	}
	return token, err
}

func (i *ICoreWebView2_4) RemoveDownloadStarting(token EventRegistrationToken) error {

	hr, _, err := i.Vtbl.RemoveDownloadStarting.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&token)),
	)
	if windows.Handle(hr) != windows.S_OK {
		return syscall.Errno(hr)
	}
	return err
}
