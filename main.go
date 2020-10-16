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
	pdf.AddPage()

	var t Tag
	t.Caliber = "5.56mm"
	t.Manufacturer = "Daniel Defense"
	t.Model = "DDM4V7"
	t.New = true
	t.Price = "1699"
	t.TagSize = Big

	var tag1coord Coord
	tag1coord.X = LongGunPageTag1PositionX
	tag1coord.Y = LongGunPageTag1PositionY

	// DrawLongGunTag(t, tag1coord, pdf)
	t.Draw(tag1coord, pdf)

	t.Caliber = "308 Winchester"
	t.Manufacturer = "PTR Industries"
	t.Model = "PTR91"
	t.New = true
	t.Price = "1029"
	t.TagSize = Big

	var tag2coord Coord
	tag2coord.X = LongGunPageTag2PositionX
	tag2coord.Y = LongGunPageTag2PositionY

	// DrawLongGunTag(t, tag2coord, pdf)
	t.Draw(tag2coord, pdf)

	t.Caliber = "5.56x45mm"
	t.Manufacturer = "Wilson Combat"
	t.Model = "Protector"
	t.New = true
	t.Price = "1999"
	t.TagSize = Big

	var tag3coord Coord
	tag3coord.X = LongGunPageTag3PositionX
	tag3coord.Y = LongGunPageTag3PositionY

	// DrawLongGunTag(t, tag3coord, pdf)
	t.Draw(tag3coord, pdf)

	pdf.AddPage()

	t.Caliber = "5.56mm"
	t.Manufacturer = "Daniel Defense"
	t.Model = "DDM4V11"
	t.New = true
	t.Price = "1729"
	t.TagSize = Big

	tag1coord.X = LongGunPageTag1PositionX
	tag1coord.Y = LongGunPageTag1PositionY

	DrawLongGunTag(t, tag1coord, pdf)

	t.Caliber = "308 Winchester"
	t.Manufacturer = "PTR Industries"
	t.Model = "PTR A3S K"
	t.New = true
	t.Price = "1029"
	t.TagSize = Big

	tag2coord.X = LongGunPageTag2PositionX
	tag2coord.Y = LongGunPageTag2PositionY

	DrawLongGunTag(t, tag2coord, pdf)

	t.Caliber = "12 GA"
	t.Manufacturer = "Wilson Combat"
	t.Model = "Border Patrol"
	t.New = true
	t.Price = "1135"
	t.TagSize = Big

	tag3coord.X = LongGunPageTag3PositionX
	tag3coord.Y = LongGunPageTag3PositionY

	DrawLongGunTag(t, tag3coord, pdf)

	err := pdf.OutputFileAndClose("hello.pdf")
	fmt.Println(err)
	fmt.Println("PDF generated")
}
