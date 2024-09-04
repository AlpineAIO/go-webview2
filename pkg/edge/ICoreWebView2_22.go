//go:build windows

package edge

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

type iCoreWebView2_22Vtbl struct {
	_IUnknownVtbl
	GetSettings                                            ComProc
	GetSource                                              ComProc
	Navigate                                               ComProc
	NavigateToString                                       ComProc
	AddNavigationStarting                                  ComProc
	RemoveNavigationStarting                               ComProc
	AddContentLoading                                      ComProc
	RemoveContentLoading                                   ComProc
	AddSourceChanged                                       ComProc
	RemoveSourceChanged                                    ComProc
	AddHistoryChanged                                      ComProc
	RemoveHistoryChanged                                   ComProc
	AddNavigationCompleted                                 ComProc
	RemoveNavigationCompleted                              ComProc
	AddFrameNavigationStarting                             ComProc
	RemoveFrameNavigationStarting                          ComProc
	AddFrameNavigationCompleted                            ComProc
	RemoveFrameNavigationCompleted                         ComProc
	AddScriptDialogOpening                                 ComProc
	RemoveScriptDialogOpening                              ComProc
	AddPermissionRequested                                 ComProc
	RemovePermissionRequested                              ComProc
	AddProcessFailed                                       ComProc
	RemoveProcessFailed                                    ComProc
	AddScriptToExecuteOnDocumentCreated                    ComProc
	RemoveScriptToExecuteOnDocumentCreated                 ComProc
	ExecuteScript                                          ComProc
	CapturePreview                                         ComProc
	Reload                                                 ComProc
	PostWebMessageAsJSON                                   ComProc
	PostWebMessageAsString                                 ComProc
	AddWebMessageReceived                                  ComProc
	RemoveWebMessageReceived                               ComProc
	CallDevToolsProtocolMethod                             ComProc
	GetBrowserProcessID                                    ComProc
	GetCanGoBack                                           ComProc
	GetCanGoForward                                        ComProc
	GoBack                                                 ComProc
	GoForward                                              ComProc
	GetDevToolsProtocolEventReceiver                       ComProc
	Stop                                                   ComProc
	AddNewWindowRequested                                  ComProc
	RemoveNewWindowRequested                               ComProc
	AddDocumentTitleChanged                                ComProc
	RemoveDocumentTitleChanged                             ComProc
	GetDocumentTitle                                       ComProc
	AddHostObjectToScript                                  ComProc
	RemoveHostObjectFromScript                             ComProc
	OpenDevToolsWindow                                     ComProc
	AddContainsFullScreenElementChanged                    ComProc
	RemoveContainsFullScreenElementChanged                 ComProc
	GetContainsFullScreenElement                           ComProc
	AddWebResourceRequested                                ComProc
	RemoveWebResourceRequested                             ComProc
	AddWebResourceRequestedFilter                          ComProc
	RemoveWebResourceRequestedFilter                       ComProc
	AddWindowCloseRequested                                ComProc
	RemoveWindowCloseRequested                             ComProc
	AddWebResourceResponseReceived                         ComProc
	RemoveWebResourceResponseReceived                      ComProc
	NavigateWithWebResourceRequest                         ComProc
	AddDOMContentLoaded                                    ComProc
	RemoveDOMContentLoaded                                 ComProc
	GetCookieManager                                       ComProc
	GetEnvironment                                         ComProc
	TrySuspend                                             ComProc
	Resume                                                 ComProc
	IsSuspended                                            ComProc
	SetVirtualHostNameToFolderMapping                      ComProc
	ClearVirtualHostNameToFolderMapping                    ComProc
	AddFrameCreated                                        ComProc
	RemoveFrameCreated                                     ComProc
	AddDownloadStarting                                    ComProc
	RemoveDownloadStarting                                 ComProc
	AddClientCertificateRequested                          ComProc
	RemoveClientCertificateRequested                       ComProc
	OpenTaskManagerWindow                                  ComProc
	PrintToPDF                                             ComProc
	AddIsMutedChanged                                      ComProc
	RemoveIsMutedChanged                                   ComProc
	GetIsMuted                                             ComProc
	SetIsMuted                                             ComProc
	AddIsDocumentPlayingAudioChanged                       ComProc
	RemoveIsDocumentPlayingAudioChanged                    ComProc
	IsDocumentPlayingAudio                                 ComProc
	AddIsDefaultDownloadDialogOpenChanged                  ComProc
	RemoveIsDefaultDownloadDialogOpenChanged               ComProc
	IsDefaultDownloadDialogOpen                            ComProc
	OpenDefaultDownloadDialog                              ComProc
	CloseDefaultDownloadDialog                             ComProc
	GetDefaultDownloadDialogCornerAlignment                ComProc
	SetDefaultDownloadDialogCornerAlignment                ComProc
	GetDefaultDownloadDialogMargin                         ComProc
	SetDefaultDownloadDialogMargin                         ComProc
	AddBasicAuthenticationRequested                        ComProc
	RemoveBasicAuthenticationRequested                     ComProc
	CallDevToolsProtocolMethodForSession                   ComProc
	AddContextMenuRequested                                ComProc
	RemoveContextMenuRequested                             ComProc
	AddStatusBarTextChanged                                ComProc
	RemoveStatusBarTextChanged                             ComProc
	StatusBarText                                          ComProc
	Profile                                                ComProc
	AddServerCertificateErrorDetected                      ComProc
	RemoveServerCertificateErrorDetected                   ComProc
	ClearServerCertificateErrorActions                     ComProc
	AddFaviconChanged                                      ComProc
	RemoveFaviconChanged                                   ComProc
	FaviconUri                                             ComProc
	GetFavicon                                             ComProc
	Print                                                  ComProc
	ShowPrintUI                                            ComProc
	PrintToPdfStream                                       ComProc
	PostSharedBufferToScript                               ComProc
	AddLaunchingExternalUriScheme                          ComProc
	RemoveLaunchingExternalUriScheme                       ComProc
	GetMemoryUsageTargetLevel                              ComProc
	SetMemoryUsageTargetLevel                              ComProc
	FrameId                                                ComProc
	ExecuteScriptWithResult                                ComProc
	AddWebResourceRequestedFilterWithRequestSourceKinds    ComProc
	RemoveWebResourceRequestedFilterWithRequestSourceKinds ComProc
	PostWebMessageAsJsonWithAdditionalObjects              ComProc
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
		uintptr(resourceKinds),
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
