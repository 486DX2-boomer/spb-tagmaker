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

	// gofpdf.MakeFont("framd.ttf", "Franklin Gothic Medium.json", ".", nil, true)
	pdf := gofpdf.New("P", "pt", "A4", "")

	pdf.AddPage()
	pdf.AddUTF8Font(SPBTagFont, "", "framd.ttf")
	pdf.SetFont(SPBTagFont, "", 12) // Call set font with a default size to avoid compiler error

	// draw 3x5 guides
	pdf.MoveTo(125, 72)
	pdf.LineTo((125 + 360), 72)
	pdf.ClosePath()
	pdf.SetLineWidth(0.5)
	pdf.DrawPath("D")

	pdf.SetXY(125, 72) // 0, 0 coordinate for tag 1

	// needed to set up image embed
	var opt gofpdf.ImageOptions
	// opt.ImageType = "png"
	pdf.ImageOptions(".\\logos\\shootpointblank.png", 140, 81, SPBLogoWidth, 0, false, opt, 0, "")
	pdf.ImageOptions(".\\logos\\daniel defense.jpg", 340, 81, LongGunManufacturerLogoWidth, 0, false, opt, 0, "")

	pdf.SetXY(140, 144) // Location of $ sign
	pdf.SetFont(SPBTagFont, "", LongGunDollarFontSize)
	pdf.Cell(45, 55, "$")

	pdf.SetXY(188, 144)
	pdf.SetFont(SPBTagFont, "", LongGunPriceFontSize)
	pdf.Cell(270, 109, "1699")
	pdf.SetFont(SPBTagFont, "", LongGunCentsFontSize)
	pdf.Write(LongGunCentsFontSize, "99")

	pdf.SetXY(136, 256)
	pdf.SetFont(SPBTagFont, "", LongGunModelFontSize)
	pdf.Cell(113, 18, "DDM4V7")

	pdf.SetXY(356, 261)
	pdf.SetFont(SPBTagFont, "", LongGunCaliberFontSize)
	pdf.CellFormat(113, 18, "5.56x45", "", 0, "R", false, 0, "")

	// draw the red separator bar
	pdf.MoveTo(135, 248)
	pdf.LineTo((135 + 338), 248)
	pdf.ClosePath()
	pdf.SetLineWidth(2.5)
	pdf.SetDrawColor(212, 88, 42)
	pdf.DrawPath("D")

	err := pdf.OutputFileAndClose("hello.pdf")
	fmt.Println(err)
}
