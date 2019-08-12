package config

import "os"

const (
	fbToken         = "FB_APP_TOKEN"
	pageID          = "PAGE_ID"
	exportDirectory = "EXPORT_DIRECTORY"
)

// GetFBAppToken returns facebook app token from env var
func GetFBAppToken() string {
	return os.Getenv(fbToken)
}

// GetPageID returns page's id from env var
func GetPageID() string {
	return os.Getenv(pageID)
}

// GetExportDirectory return export directory from env var
func GetExportDirectory() string {
	return os.Getenv(exportDirectory)
}
