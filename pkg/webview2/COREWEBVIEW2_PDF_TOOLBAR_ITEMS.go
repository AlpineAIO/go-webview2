//go:build windows

package webview2

type COREWEBVIEW2_PDF_TOOLBAR_ITEMS uint32

const (
	COREWEBVIEW2_PDF_TOOLBAR_ITEMS_NONE          = 0x0
	COREWEBVIEW2_PDF_TOOLBAR_ITEMS_SAVE          = 0x1
	COREWEBVIEW2_PDF_TOOLBAR_ITEMS_PRINT         = 0x2
	COREWEBVIEW2_PDF_TOOLBAR_ITEMS_SAVE_AS       = 0x4
	COREWEBVIEW2_PDF_TOOLBAR_ITEMS_ZOOM_IN       = 0x8
	COREWEBVIEW2_PDF_TOOLBAR_ITEMS_ZOOM_OUT      = 0x10
	COREWEBVIEW2_PDF_TOOLBAR_ITEMS_ROTATE        = 0x20
	COREWEBVIEW2_PDF_TOOLBAR_ITEMS_FIT_PAGE      = 0x40
	COREWEBVIEW2_PDF_TOOLBAR_ITEMS_PAGE_LAYOUT   = 0x80
	COREWEBVIEW2_PDF_TOOLBAR_ITEMS_BOOKMARKS     = 0x100
	COREWEBVIEW2_PDF_TOOLBAR_ITEMS_PAGE_SELECTOR = 0x200
	COREWEBVIEW2_PDF_TOOLBAR_ITEMS_SEARCH        = 0x400
	COREWEBVIEW2_PDF_TOOLBAR_ITEMS_FULL_SCREEN   = 0x800
	COREWEBVIEW2_PDF_TOOLBAR_ITEMS_MORE_SETTINGS = 0x1000
)
