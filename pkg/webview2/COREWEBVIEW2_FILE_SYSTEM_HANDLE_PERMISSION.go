//go:build windows

package webview2

type COREWEBVIEW2_FILE_SYSTEM_HANDLE_PERMISSION uint32

const (
	COREWEBVIEW2_FILE_SYSTEM_HANDLE_PERMISSION_READ_ONLY  = 0
	COREWEBVIEW2_FILE_SYSTEM_HANDLE_PERMISSION_READ_WRITE = 1
)