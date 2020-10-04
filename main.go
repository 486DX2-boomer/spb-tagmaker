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
	pdf.AddUTF8Font(SPBTagFont, "", "framd.ttf")
	pdf.SetFont(SPBTagFont, "", 12) // Call set font with a default size to avoid compiler error

	fmt.Println(pdf.GetXY())

	pdf.CellFormat(24, 36, "", "", 0, "L", true, 0, "") // empty spacer
	fmt.Println(pdf.GetXY())

	pdf.SetFont(SPBTagFont, "", LongGunDollarFontSize)
	pdf.CellFormat(16, 19, "$", "", 0, "L", false, 0, "")
	fmt.Println(pdf.GetXY())

	pdf.SetFont(SPBTagFont, "", LongGunPriceFontSize)
	pdf.CellFormat(99, 36, "1699", "", 0, "L", false, 0, "")
	fmt.Println(pdf.GetXY())

	pdf.SetFont(SPBTagFont, "", LongGunDollarFontSize)
	pdf.CellFormat(34, 16, "99", "", 0, "L", false, 0, "")
	fmt.Println(pdf.GetXY())

	pdf.CellFormat(24, 36, "", "", 1, "L", true, 0, "") // empty spacer. Adding 1 to ln int breaks to the next line
	fmt.Println(pdf.GetXY())

	pdf.CellFormat(24, 36, "", "", 0, "L", true, 0, "") // empty spacer
	fmt.Println(pdf.GetXY())

	pdf.SetFont(SPBTagFont, "", LongGunModelFontSize)
	pdf.CellFormat(4, 19, "DDM4V7", "", 0, "L", false, 0, "")
	fmt.Println(pdf.GetXY())

	pdf.SetFont(SPBTagFont, "", LongGunCaliberFontSize)
	pdf.CellFormat(0, 16, "5.56mm", "", 0, "R", false, 0, "")
	fmt.Println(pdf.GetXY())

	pdf.SetXY(0, 0)
	fmt.Println(pdf.GetXY()) // I need to convert the above code to use SetXY() instead of advancing the position with cells

	err := pdf.OutputFileAndClose("hello.pdf")
	fmt.Println(err)
}
