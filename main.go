package main

import (
	"fmt"

	"github.com/jung-kurt/gofpdf"
)

func main() {

	// Step 1 : Set up a web server
	// Step 2 : Display all tags in memory
	// Step 3 : User clicks "add tag"
	// Step 4 : Collect information from form, add the tag to memory
	// Repeat from step 2 until PRINT TAG is clicked by user
	// Step 5 : Get all tags in memory, sort by handgun, long gun, new and used
	// Step 6 : Generate a pdf with all handgun and long gun tags (handgun and long gun tags are always separate pages)
	// If manufacturer logo isn't found in the /logos folder, query Google Images for a logo. Save the logo on the server for later
	// The web server runs indefinitely, waiting for new tags or for the user to delete all tags in memory and start over.

	// Creating a pdf

	// pdf.AddPage()

	var l List

	l = append(l, NewTag("Walther", "PPK", "380 Auto", false, "779", Small))
	l = append(l, NewTag("Colt", "M4 Carbine", "5.56mm", false, "1099", Big))
	l = append(l, NewTag("Springfield", "Saint", "5.56mm", true, "879", Big))

	l = append(l, NewTag("Glock", "G44", "22 LR", false, "389", Small))
	l = append(l, NewTag("Smith & Wesson", "617", "22 LR", true, "709", Small))
	l = append(l, NewTag("Kel Tec", "PMR 30", "22 WMR", true, "399", Small))

	l = append(l, NewTag("Ruger", "Wrangler", "22 LR", true, "199", Small))
	l = append(l, NewTag("Taurus", "Judge Tracker", "45 LC/410", true, "469", Small))
	l = append(l, NewTag("Charter Arms", "Lavender Lady", "38 Special", true, "414", Small))
	l = append(l, NewTag("Rock Island", "GI Standard CS", "45 ACP", true, "419", Small))
	l = append(l, NewTag("FN", "FNX-45 Tactical", "45 ACP", true, "1229", Small))
	l = append(l, NewTag("CZ", "97B", "45 ACP", true, "719", Small))
	l = append(l, NewTag("Smith & Wesson", "M&P 40 FDE", "40 S&W", true, "699", Small))
	l = append(l, NewTag("GSG", "GSG-16 Carbine", "22 LR", false, "349", Big))
	l = append(l, NewTag("HK", "HK 45c", "45 ACP", true, "779", Small))

	BuildDocument(l, NewDocument())

}

// BuildDocument accepts a list of tags and builds a full document, sorting the handgun and long gun tags and then drawing each tag to an appropriate page.
// BuildDocument must be supplied with a pointer to a gofpdf object which can be supplied with the NewDocument() function
func BuildDocument(inputList []Tag, pdf *gofpdf.Fpdf) {

	longGunPageCoord := GetLongGunPageCoord()
	handGunPageCoord := GetHandgunPageCoord()

	var handgunTags List
	var longGunTags List

	for i := 0; i < len(inputList); i++ {
		if inputList[i].TagSize == Big {
			longGunTags = append(longGunTags, inputList[i])
		} else if inputList[i].TagSize == Small {
			handgunTags = append(handgunTags, inputList[i])
		}
	}

	for i := 0; i < len(longGunTags); i++ {
		if i%3 == 0 {
			pdf.AddPage()
		}
		longGunTags[i].Draw(longGunPageCoord[(i%3)], pdf)
	}

	for i := 0; i < len(handgunTags); i++ {
		if i%10 == 0 {
			pdf.AddPage()
		}
		handgunTags[i].Draw(handGunPageCoord[(i%10)], pdf)
	}

	err := pdf.OutputFileAndClose("output.pdf")
	fmt.Println(err)
	fmt.Println("PDF generated")

}

// NewDocument initializes a pdf, sets the font and font size, and then returns the pdf. This is meant to be used in conjunction with BuildDocument()
func NewDocument() *gofpdf.Fpdf {
	pdf := gofpdf.New("P", "pt", "A4", "")
	pdf.AddUTF8Font(SPBTagFont, "", "framd.ttf")
	pdf.SetFont(SPBTagFont, "", 12) // Call set font with a default size to avoid compiler error
	pdf.SetTextColor(0, 0, 0)

	return pdf
}
