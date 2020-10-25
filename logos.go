package main

import (
	"fmt"
	"os"
	"strings"
)

// GetLogoImagePath returns a gif, jpeg, or png in the logos directory. fileName should NOT include the extension. Defaults to blank.png if not found
func GetLogoImagePath(fileName string) string {
	// List of possible extensions
	extension := []string{".png", ".jpg", ".jpeg", ".gif"}

	var logoPath string
	l := ".\\logos\\" + strings.ToLower(fileName)
	fileFound := false

	// Find the file
	for i := 0; i < len(extension); i++ {
		_, err := os.Stat(l + extension[i])
		if os.IsNotExist(err) {
			fmt.Println(l + extension[i] + " not found")
			fileFound = false
		} else {
			fileFound = true
			logoPath = l + extension[i]
			fmt.Println("Found " + logoPath)
			return logoPath
		}
	}

	// If not found, scrape Google Images for the logo and download it to the server.
	// If it's a vector, it will need to be rasterized
	if fileFound == false {
		fmt.Println(fileName + " not found, defaulting to blank")
		logoPath = ".\\logos\\blank.png"
		return logoPath
	}

	fmt.Println(logoPath)
	return logoPath

}
