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
	pdf := gofpdf.New("P", "pt", "A4", "")
	pdf.AddUTF8Font(SPBTagFont, "", "framd.ttf")
	pdf.SetFont(SPBTagFont, "", 12) // Call set font with a default size to avoid compiler error

	// pdf.AddPage()

	var l List

	l = append(l, NewTag("Daniel Defense", "DDM4V7", "5.56mm", true, "1699", Big))
	l = append(l, NewTag("PTR Industries", "PTR91", "308 Winchester", true, "1029", Big))
	l = append(l, NewTag("Wilson Combat", "Protector", "5.56mm", true, "1999", Big))

	l = append(l, NewTag("Daniel Defense", "DDM4V11", "5.56mm", true, "1729", Big))
	l = append(l, NewTag("PTR Industries", "PTR A3S K", "308 Winchester", true, "1029", Big))
	l = append(l, NewTag("Wilson Combat", "Border Patrol", "12 GA", true, "1135", Big))

	l = append(l, NewTag("Glock", "G19C Gen 3", "9x19mm", true, "549", Small))
	l = append(l, NewTag("Canik", "TP9", "9x19mm", true, "349", Small))
	l = append(l, NewTag("Bersa", "Thunder 380", "380 ACP", true, "329", Small))
	l = append(l, NewTag("CZ", "P-10C", "9x19mm", true, "529", Small))
	l = append(l, NewTag("Chiappa", "Rhino 60DS", "357 Magnum", true, "1299", Small))
	l = append(l, NewTag("Smith & Wesson", "M&P 9 Shield EZ PC", "9x19mm", true, "559", Small))
	l = append(l, NewTag("Springfield", "Loaded Operator", "45 ACP", true, "1299", Small))
	l = append(l, NewTag("Walther", "PPQ", "9x19mm", true, "619", Small))
	l = append(l, NewTag("Magnum Research", "Desert Eagle", "50 Action Express", true, "1999", Small))
	l = append(l, NewTag("FN", "509C", "9x19mm", true, "899", Small))

	l = append(l, NewTag("FN", "Five SeveN", "5.7x28mm", true, "999", Small))
	l = append(l, NewTag("Ruger", "10/22", "22 LR", true, "329", Big))

	BuildDocument(l, pdf)

	err := pdf.OutputFileAndClose("hello.pdf")
	fmt.Println(err)
	fmt.Println("PDF generated")
}

// BuildDocument accepts a list of tags and builds a full document, sorting the handgun and long gun tags and then drawing each tag to an appropriate page.
// BuildDocument must be supplied with a pointer to a gofpdf object which should be created on program initialization
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

}
