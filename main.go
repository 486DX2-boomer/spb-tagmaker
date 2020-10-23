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

	t.Draw(tag1coord, pdf)

	t.Caliber = "308 Winchester"
	t.Manufacturer = "PTR Industries"
	t.Model = "PTR A3S K"
	t.New = true
	t.Price = "1029"
	t.TagSize = Big

	tag2coord.X = LongGunPageTag2PositionX
	tag2coord.Y = LongGunPageTag2PositionY

	t.Draw(tag2coord, pdf)

	t.Caliber = "12 GA"
	t.Manufacturer = "Wilson Combat"
	t.Model = "Border Patrol"
	t.New = true
	t.Price = "1135"
	t.TagSize = Big

	tag3coord.X = LongGunPageTag3PositionX
	tag3coord.Y = LongGunPageTag3PositionY

	t.Draw(tag3coord, pdf)

	// Test handgun page draw
	pdf.AddPage()

	t.Caliber = "9x19mm"
	t.Manufacturer = "Glock"
	t.Model = "G19C Gen 3"
	t.New = true
	t.Price = "549"
	t.TagSize = Small

	tag1coord.X = HandgunPageTag1PositionX
	tag1coord.Y = HandgunPageTag1PositionY

	t.Draw(tag1coord, pdf)

	t.Caliber = "9x19mm"
	t.Manufacturer = "Canik"
	t.Model = "TP9"
	t.New = true
	t.Price = "349"
	t.TagSize = Small

	tag2coord.X = HandgunPageTag2PositionX
	tag2coord.Y = HandgunPageTag2PositionY

	t.Draw(tag2coord, pdf)

	tag3coord.X = HandgunPageTag3PositionX
	tag3coord.Y = HandgunPageTag3PositionY

	var tag4coord Coord
	var tag5coord Coord
	var tag6coord Coord
	var tag7coord Coord
	var tag8coord Coord
	var tag9coord Coord
	var tag10coord Coord

	tag4coord.X = HandgunPageTag4PositionX
	tag4coord.Y = HandgunPageTag4PositionY

	tag5coord.X = HandgunPageTag5PositionX
	tag5coord.Y = HandgunPageTag5PositionY

	tag6coord.X = HandgunPageTag6PositionX
	tag6coord.Y = HandgunPageTag6PositionY

	tag7coord.X = HandgunPageTag7PositionX
	tag7coord.Y = HandgunPageTag7PositionY

	tag8coord.X = HandgunPageTag8PositionX
	tag8coord.Y = HandgunPageTag8PositionY

	tag9coord.X = HandgunPageTag9PositionX
	tag9coord.Y = HandgunPageTag9PositionY

	tag10coord.X = HandgunPageTag10PositionX
	tag10coord.Y = HandgunPageTag10PositionY

	t.Caliber = "380 ACP"
	t.Manufacturer = "Bersa"
	t.Model = "Thunder 380"
	t.New = true
	t.Price = "329"
	t.TagSize = Small
	t.Draw(tag3coord, pdf)

	t.Caliber = "9x19mm"
	t.Manufacturer = "CZ"
	t.Model = "P-10C"
	t.New = true
	t.Price = "529"
	t.TagSize = Small
	t.Draw(tag4coord, pdf)

	t.Caliber = "357 Magnum"
	t.Manufacturer = "Chiappa"
	t.Model = "Rhino 60DS"
	t.New = true
	t.Price = "1299"
	t.TagSize = Small
	t.Draw(tag5coord, pdf)

	t.Caliber = "9x19mm"
	t.Manufacturer = "FN"
	t.Model = "509c"
	t.New = true
	t.Price = "899"
	t.TagSize = Small
	t.Draw(tag6coord, pdf)

	t.Caliber = "9mm"
	t.Manufacturer = "Smith & Wesson"
	t.Model = "M&P 9 Shield EZ PC"
	t.New = true
	t.Price = "559"
	t.TagSize = Small
	t.Draw(tag7coord, pdf)

	t.Caliber = "45 ACP"
	t.Manufacturer = "Springfield"
	t.Model = "Loaded Operator"
	t.New = true
	t.Price = "1299"
	t.TagSize = Small
	t.Draw(tag8coord, pdf)

	t.Caliber = "9x19mm"
	t.Manufacturer = "Walther"
	t.Model = "PPQ"
	t.New = true
	t.Price = "619"
	t.TagSize = Small
	t.Draw(tag9coord, pdf)

	t.Caliber = "50 Action Express"
	t.Manufacturer = "Magnum Research"
	t.Model = "Desert Eagle"
	t.New = true
	t.Price = "1999"
	t.TagSize = Small
	t.Draw(tag10coord, pdf)

	err := pdf.OutputFileAndClose("hello.pdf")
	fmt.Println(err)
	fmt.Println("PDF generated")
}
