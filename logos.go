package main

// GetLogoFileTypes returns a slice of all acceptable image file types for use when checking the logos directory for manufacturer logo images. Currently unused
func GetLogoFileTypes() []string {
	var AcceptableLogoTypes = []string{"jpg", "JPG", "jpeg", "JPEG", "gif", "GIF", "webp", "WEBP"}
	return AcceptableLogoTypes
}
