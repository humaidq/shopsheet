package models

import (
	"mime/multipart"
)

// NewForm represents the data binding structure of a form to create
// a new website.
type NewForm struct {
	SiteName        string                `form:"site_name" binding:"Required"`
	SiteDescription string                `form:"site_description" binding:"Required"`
	Email           string                `form:"email" binding:"Required;Email"`
	LogoURL         string                `form:"logo_url" binding:"Required;Url"`
	BannerURL       string                `form:"banner_url" binding:"Required;Url"`
	Spreadsheet     *multipart.FileHeader `form:"spreadsheet" binding:"Required"`
}
