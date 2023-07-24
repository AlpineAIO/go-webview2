//go:build windows

package webview2

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

type _ICoreWebView2CertificateVtbl struct {
	_IUnknownVtbl
	GetSubject                          ComProc
	GetIssuer                           ComProc
	GetValidFrom                        ComProc
	GetValidTo                          ComProc
	GetDerEncodedSerialNumber           ComProc
	GetDisplayName                      ComProc
	ToPemEncoding                       ComProc
	GetPemEncodedIssuerCertificateChain ComProc
}

type ICoreWebView2Certificate struct {
	vtbl *_ICoreWebView2CertificateVtbl
}

func (i *ICoreWebView2Certificate) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2Certificate) GetSubject() (string, error) {
	var err error
	// Create *uint16 to hold result
	var _value *uint16

	_, _, err = i.vtbl.GetSubject.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_value)),
	)
	if err != windows.ERROR_SUCCESS {
		return "", err
	} // Get result and cleanup
	value := windows.UTF16PtrToString(_value)
	windows.CoTaskMemFree(unsafe.Pointer(_value))
	return value, nil
}

func (i *ICoreWebView2Certificate) GetIssuer() (string, error) {
	var err error
	// Create *uint16 to hold result
	var _value *uint16

	_, _, err = i.vtbl.GetIssuer.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_value)),
	)
	if err != windows.ERROR_SUCCESS {
		return "", err
	} // Get result and cleanup
	value := windows.UTF16PtrToString(_value)
	windows.CoTaskMemFree(unsafe.Pointer(_value))
	return value, nil
}

func (i *ICoreWebView2Certificate) GetValidFrom() (float64, error) {
	var err error

	var value float64

	_, _, err = i.vtbl.GetValidFrom.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return 0.0, err
	}
	return value, nil
}

func (i *ICoreWebView2Certificate) GetValidTo() (float64, error) {
	var err error

	var value float64

	_, _, err = i.vtbl.GetValidTo.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return 0.0, err
	}
	return value, nil
}

func (i *ICoreWebView2Certificate) GetDerEncodedSerialNumber() (string, error) {
	var err error
	// Create *uint16 to hold result
	var _value *uint16

	_, _, err = i.vtbl.GetDerEncodedSerialNumber.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_value)),
	)
	if err != windows.ERROR_SUCCESS {
		return "", err
	} // Get result and cleanup
	value := windows.UTF16PtrToString(_value)
	windows.CoTaskMemFree(unsafe.Pointer(_value))
	return value, nil
}

func (i *ICoreWebView2Certificate) GetDisplayName() (string, error) {
	var err error
	// Create *uint16 to hold result
	var _value *uint16

	_, _, err = i.vtbl.GetDisplayName.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_value)),
	)
	if err != windows.ERROR_SUCCESS {
		return "", err
	} // Get result and cleanup
	value := windows.UTF16PtrToString(_value)
	windows.CoTaskMemFree(unsafe.Pointer(_value))
	return value, nil
}

func (i *ICoreWebView2Certificate) ToPemEncoding() (string, error) {
	var err error
	// Create *uint16 to hold result
	var _pemEncodedData *uint16

	_, _, err = i.vtbl.ToPemEncoding.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(_pemEncodedData)),
	)
	if err != windows.ERROR_SUCCESS {
		return "", err
	} // Get result and cleanup
	pemEncodedData := windows.UTF16PtrToString(_pemEncodedData)
	windows.CoTaskMemFree(unsafe.Pointer(_pemEncodedData))
	return pemEncodedData, nil
}

func (i *ICoreWebView2Certificate) GetPemEncodedIssuerCertificateChain() (*ICoreWebView2StringCollection, error) {
	var err error

	var value *ICoreWebView2StringCollection

	_, _, err = i.vtbl.GetPemEncodedIssuerCertificateChain.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&value)),
	)
	if err != windows.ERROR_SUCCESS {
		return nil, err
	}
	return value, nil
}