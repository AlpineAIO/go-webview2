//go:build windows

package edge

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

type iCoreWebView2_20Vtbl struct {
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

type ICoreWebView2_20 struct {
	vtbl *iCoreWebView2_20Vtbl
}

func (i *ICoreWebView2_20) GetFrameId() (uint32, error) {
	var result uint32

	_, _, err := i.vtbl.FrameId.Call(
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
