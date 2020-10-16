package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/jung-kurt/gofpdf"
)

// Tag size. Big = 3x5". Small = 2x3"
const (
	Big   = false
	Small = true
)

// Long gun 3x5 dimensions in points: 360 x 216 pt

// Long gun SPB logo offset : 13, 9
// Long gun manufacturer logo offset: 214, 9
// Long gun $ sign offset : 13, 72
// Long gun price offset : 62, 72

// Hand gun 2x3 dimensions in points:

// Hand gun SPB logo offset :
// Hand gun manufacturer logo offset :
// Hand gun $ sign offset :
// Hand gun price offset :

// LongGunSPBLogoWidth defines the SPB logo width in pts
const LongGunSPBLogoWidth float64 = 129

// LongGunManufacturerLogoWidth defines the width of the manufacturer's logo in pts
const LongGunManufacturerLogoWidth float64 = 104

// Font sizes and font family
const (
	LongGunDollarFontSize  float64 = 60
	LongGunPriceFontSize   float64 = 113
	LongGunCentsFontSize   float64 = 60
	LongGunModelFontSize   float64 = 20
	LongGunCaliberFontSize float64 = 20
	SPBTagFont             string  = "Franklin Gothic Medium"
)

// The top left coordinates of each long gun tag per 8.5x11 page in pts
const (
	LongGunPageTag1PositionX float64 = 125
	LongGunPageTag1PositionY float64 = 72

	LongGunPageTag2PositionX float64 = 125
	LongGunPageTag2PositionY float64 = 287

	LongGunPageTag3PositionX float64 = 125
	LongGunPageTag3PositionY float64 = 504
)

// The coordinate offsets for the long gun tag in pts
const (
	LongGunSPBLogoOffsetX float64 = 15
	LongGunSPBLogoOffsetY float64 = 9

	LongGunManufacturerLogoOffsetX float64 = 215
	LongGunManufacturerLogoOffsetY float64 = 9

	LongGunDollarSignOffsetX float64 = 15
	LongGunDollarSignOffsetY float64 = 72

	LongGunPriceOffsetX float64 = 63
	LongGunPriceOffsetY float64 = 72

	LongGunSeparatorBarOffsetX float64 = 10
	LongGunSeparatorBarOffsetY float64 = 176

	LongGunModelOffsetX float64 = 11
	LongGunModelOffsetY float64 = 189

	LongGunCaliberOffsetX float64 = 231
	LongGunCaliberOffsetY float64 = 189
)

// Separator bar constants
const (
	SeparatorBarLineWidth     float64 = 2.5
	LongGunSeparatorBarLength float64 = 338
	// HandGunSeparatorBarLength
	SeparatorBarColorR int = 212
	SeparatorBarColorG int = 88
	SeparatorBarColorB int = 42
)

// Tag type defines all relevant information to be stored in the List and provided to the tag draw function
type Tag struct {
	Price        string // no 99 cents at the end
	Manufacturer string
	Model        string
	Caliber      string
	New          bool // whether the gun is new or used
	TagSize      Size // Big or Small
}

// Size type is used in the Tag type to store whether the tag is large format or small format.
type Size bool

// Coord is used to store x, y coordinate pairs
type Coord struct {
	X float64
	Y float64
}

// Constructor function for tag should go here

// DrawLongGunTag accepts a single tag and draws it to the coordinate specified. Deprecate in favor of Tag.Draw() method
func DrawLongGunTag(tag Tag, zero Coord, pdf *gofpdf.Fpdf) {

	// // draw 3x5 guides
	// pdf.MoveTo(125, 72)
	// pdf.LineTo((125 + 360), 72)
	// pdf.ClosePath()
	// pdf.SetLineWidth(0.5)
	// pdf.DrawPath("D")

	// pdf.SetXY(125, 72) // 0, 0 coordinate for tag 1

	// needed to set up image embed
	var opt gofpdf.ImageOptions
	// opt.ImageType = "png"

	// Add the SPB logo
	pdf.ImageOptions(".\\logos\\shootpointblank.png", (zero.X + LongGunSPBLogoOffsetX), (zero.Y + LongGunSPBLogoOffsetY), LongGunSPBLogoWidth, 0, false, opt, 0, "")

	// get logo path
	// This needs to be changed to accept .png and other file types
	LogoPath := (".\\logos\\" + tag.Manufacturer + ".jpg")

	pdf.ImageOptions(LogoPath, (zero.X + LongGunManufacturerLogoOffsetX), (zero.Y + LongGunManufacturerLogoOffsetY), LongGunManufacturerLogoWidth, 0, false, opt, 0, "")

	pdf.SetXY((zero.X + LongGunDollarSignOffsetX), (zero.Y + LongGunDollarSignOffsetY)) // Location of $ sign
	pdf.SetFont(SPBTagFont, "", LongGunDollarFontSize)
	pdf.Cell(45, 55, "$")

	pdf.SetXY((zero.X + LongGunPriceOffsetX), (zero.Y + LongGunPriceOffsetY))
	pdf.SetFont(SPBTagFont, "", LongGunPriceFontSize)
	pdf.Cell(270, 109, tag.Price)
	pdf.SetFont(SPBTagFont, "", LongGunCentsFontSize)
	pdf.Write(LongGunCentsFontSize, "99")

	pdf.SetXY((zero.X + LongGunModelOffsetX), (zero.Y + LongGunModelOffsetY))
	pdf.SetFont(SPBTagFont, "", LongGunModelFontSize)
	pdf.CellFormat(113, 18, tag.Model, "", 0, "L", false, 0, "")

	pdf.SetXY((zero.X + LongGunCaliberOffsetX), (zero.Y + LongGunCaliberOffsetY))
	pdf.SetFont(SPBTagFont, "", LongGunCaliberFontSize)
	pdf.CellFormat(113, 18, tag.Caliber, "", 0, "R", false, 0, "")

	// draw the red separator bar
	pdf.MoveTo((zero.X + LongGunSeparatorBarOffsetX), (zero.Y + LongGunSeparatorBarOffsetY))
	pdf.LineTo((zero.X + LongGunSeparatorBarLength), (zero.Y + LongGunSeparatorBarOffsetY))
	pdf.ClosePath()
	pdf.SetLineWidth(SeparatorBarLineWidth)
	pdf.SetDrawColor(SeparatorBarColorR, SeparatorBarColorB, SeparatorBarColorG)
	pdf.DrawPath("D")
}

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

// Draw the tag to the position coordinate provided
func (t Tag) Draw(zero Coord, pdf *gofpdf.Fpdf) {
	// draw a long gun tag
	if t.TagSize == Big {
		// // draw 3x5 guides
		// pdf.MoveTo(125, 72)
		// pdf.LineTo((125 + 360), 72)
		// pdf.ClosePath()
		// pdf.SetLineWidth(0.5)
		// pdf.DrawPath("D")

		// pdf.SetXY(125, 72) // 0, 0 coordinate for tag 1

		// needed to set up image embed
		var opt gofpdf.ImageOptions
		// opt.ImageType = "png"

		// Add the SPB logo
		pdf.ImageOptions(".\\logos\\shootpointblank.png", (zero.X + LongGunSPBLogoOffsetX), (zero.Y + LongGunSPBLogoOffsetY), LongGunSPBLogoWidth, 0, false, opt, 0, "")

		// get logo path
		// This needs to be changed to accept .png and other file types
		LogoPath := (GetLogoImagePath(t.Manufacturer))

		pdf.ImageOptions(LogoPath, (zero.X + LongGunManufacturerLogoOffsetX), (zero.Y + LongGunManufacturerLogoOffsetY), LongGunManufacturerLogoWidth, 0, false, opt, 0, "")

		pdf.SetXY((zero.X + LongGunDollarSignOffsetX), (zero.Y + LongGunDollarSignOffsetY)) // Location of $ sign
		pdf.SetFont(SPBTagFont, "", LongGunDollarFontSize)
		pdf.Cell(45, 55, "$")

		pdf.SetXY((zero.X + LongGunPriceOffsetX), (zero.Y + LongGunPriceOffsetY))
		// Scale down the LongGunPriceFontSize based on number of digits in the price
		var fontSizeOffset float64
		if len(t.Price) <= 3 {
			fontSizeOffset = 0
		} else if len(t.Price) == 4 {
			fontSizeOffset = 28
		} else if len(t.Price) == 5 {
			fontSizeOffset = 44
		} else if len(t.Price) == 6 {
			fontSizeOffset = 58
		} else if len(t.Price) >= 7 {
			fontSizeOffset = 72
		}
		pdf.SetFont(SPBTagFont, "", (LongGunPriceFontSize - fontSizeOffset))
		// Original cell width for the price was 270 points, this needs to be adjusted down to 200 so that the 99 cents fits in the tag
		// Need to move the cell dimensions to constants
		// pdf.Cell(270, 109, t.Price)
		pdf.Cell(200, 109, t.Price)
		pdf.SetFont(SPBTagFont, "", LongGunCentsFontSize)
		pdf.Write(LongGunCentsFontSize, "99")

		pdf.SetXY((zero.X + LongGunModelOffsetX), (zero.Y + LongGunModelOffsetY))
		// Scale down the LongGunModelFontSize based on number of characters in the model name
		// This hand tuned offset will gracefully handle around ~44 characters before it starts overflowing, but may collide with the caliber cell if the caliber is lengthy
		if len(t.Model) <= 12 {
			fontSizeOffset = 0
		} else if len(t.Model) > 12 && len(t.Model) < 18 {
			fontSizeOffset = 4
		} else if len(t.Model) > 18 {
			fontSizeOffset = 8
		}
		pdf.SetFont(SPBTagFont, "", (LongGunModelFontSize - fontSizeOffset))
		pdf.CellFormat(113, 18, t.Model, "", 0, "L", false, 0, "")

		pdf.SetXY((zero.X + LongGunCaliberOffsetX), (zero.Y + LongGunCaliberOffsetY))
		// Scale down the LongGunCaliberFontSize based on number of characters in the caliber designation
		pdf.SetFont(SPBTagFont, "", LongGunCaliberFontSize)
		pdf.CellFormat(113, 18, t.Caliber, "", 0, "R", false, 0, "")

		// draw the red separator bar
		pdf.MoveTo((zero.X + LongGunSeparatorBarOffsetX), (zero.Y + LongGunSeparatorBarOffsetY))
		pdf.LineTo((zero.X + LongGunSeparatorBarLength), (zero.Y + LongGunSeparatorBarOffsetY))
		pdf.ClosePath()
		pdf.SetLineWidth(SeparatorBarLineWidth)
		pdf.SetDrawColor(SeparatorBarColorR, SeparatorBarColorB, SeparatorBarColorG)
		pdf.DrawPath("D")
	} else if t.TagSize == Small {
		return
	}

}
