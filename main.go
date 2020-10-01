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
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.AddUTF8Font("Franklin Gothic Medium", "", "framd.ttf")
	pdf.SetFont("Franklin Gothic Medium", "", 65)
	pdf.Cell(40, 10, "9999")
	pdf.SubWrite(-12, "99", 33, 0, 0, "")
	err := pdf.OutputFileAndClose("hello.pdf")
	fmt.Println(err)
}
