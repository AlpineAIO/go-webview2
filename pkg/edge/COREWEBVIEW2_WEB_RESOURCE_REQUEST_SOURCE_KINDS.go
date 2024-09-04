//go:build windows

package edge

type COREWEBVIEW2_WEB_RESOURCE_REQUEST_SOURCE_KINDS uint32

const (
	COREWEBVIEW2_WEB_RESOURCE_REQUEST_SOURCE_KINDS_NONE           = 0
	COREWEBVIEW2_WEB_RESOURCE_REQUEST_SOURCE_KINDS_DOCUMENT       = 1
	COREWEBVIEW2_WEB_RESOURCE_REQUEST_SOURCE_KINDS_SHARED_WORKER  = 2
	COREWEBVIEW2_WEB_RESOURCE_REQUEST_SOURCE_KINDS_SERVICE_WORKER = 4
	COREWEBVIEW2_WEB_RESOURCE_REQUEST_SOURCE_KINDS_ALL            = 0xffffffff
)
