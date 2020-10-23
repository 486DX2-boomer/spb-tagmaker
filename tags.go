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

// SPBTagFont is the font family used for all price tags
const SPBTagFont string = "Franklin Gothic Medium"

// Long gun 3x5 dimensions in points: 360 x 216 pt

// Long gun SPB logo offset : 13, 9
// Long gun manufacturer logo offset: 214, 9
// Long gun $ sign offset : 13, 72
// Long gun price offset : 62, 72

// Hand gun 2x3 dimensions in points: 216 x 144 pt

// Hand gun SPB logo offset :
// Hand gun manufacturer logo offset :
// Hand gun $ sign offset :
// Hand gun price offset :

// Long gun constants

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

// Handgun constants

// HandgunSPBLogoWidth defines the SPB logo width in pts
const HandgunSPBLogoWidth float64 = 92

// HandgunManufacturerLogoWidth defines the width of the manufacturer's logo in pts
const HandgunManufacturerLogoWidth float64 = 54

// HandgunManufacturerLogoHeight defines the height of the manufacturer's logo in pts
const HandgunManufacturerLogoHeight float64 = 32

// Font sizes and font family
const (
	HandgunDollarFontSize  float64 = 34
	HandgunPriceFontSize   float64 = 58
	HandgunCentsFontSize   float64 = 34
	HandgunModelFontSize   float64 = 12
	HandgunCaliberFontSize float64 = 12
)

// The top left coordinates of each handgun tag per 8.5x11 page in pts
const (
	HandgunPageTag1PositionX float64 = 90
	HandgunPageTag1PositionY float64 = 28

	HandgunPageTag2PositionX float64 = 305
	HandgunPageTag2PositionY float64 = 28

	HandgunPageTag3PositionX float64 = 90
	HandgunPageTag3PositionY float64 = 172

	HandgunPageTag4PositionX float64 = 305
	HandgunPageTag4PositionY float64 = 172

	HandgunPageTag5PositionX float64 = 90
	HandgunPageTag5PositionY float64 = 316

	HandgunPageTag6PositionX float64 = 305
	HandgunPageTag6PositionY float64 = 316

	HandgunPageTag7PositionX float64 = 90
	HandgunPageTag7PositionY float64 = 460

	HandgunPageTag8PositionX float64 = 305
	HandgunPageTag8PositionY float64 = 460

	HandgunPageTag9PositionX float64 = 90
	HandgunPageTag9PositionY float64 = 604

	HandgunPageTag10PositionX float64 = 305
	HandgunPageTag10PositionY float64 = 604
)

// The coordinate offsets for the handgun tag in pts
const (
	HandgunSPBLogoOffsetX float64 = 7
	HandgunSPBLogoOffsetY float64 = 20

	HandgunManufacturerLogoOffsetX float64 = 130
	HandgunManufacturerLogoOffsetY float64 = 20

	HandgunDollarSignOffsetX float64 = 7
	HandgunDollarSignOffsetY float64 = 64

	HandgunPriceOffsetX float64 = 34
	HandgunPriceOffsetY float64 = 64

	HandgunSeparatorBarOffsetX float64 = 7
	HandgunSeparatorBarOffsetY float64 = 120

	HandgunModelOffsetX float64 = 7
	HandgunModelOffsetY float64 = 130

	HandgunCaliberOffsetX float64 = 160
	HandgunCaliberOffsetY float64 = 130
)

// Separator bar constants
const (
	SeparatorBarLineWidth     float64 = 2.5
	LongGunSeparatorBarLength float64 = 338
	HandgunSeparatorBarLength float64 = 200
	SeparatorBarColorR        int     = 212
	SeparatorBarColorG        int     = 88
	SeparatorBarColorB        int     = 42
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

		// Add the SPB logo
		pdf.ImageOptions(".\\logos\\shootpointblank.png", (zero.X + LongGunSPBLogoOffsetX), (zero.Y + LongGunSPBLogoOffsetY), LongGunSPBLogoWidth, 0, false, opt, 0, "")

		// get logo path
		LogoPath := (GetLogoImagePath(t.Manufacturer))

		pdf.ImageOptions(LogoPath, (zero.X + LongGunManufacturerLogoOffsetX), (zero.Y + LongGunManufacturerLogoOffsetY), LongGunManufacturerLogoWidth, 0, false, opt, 0, "")

		pdf.SetXY((zero.X + LongGunDollarSignOffsetX), (zero.Y + LongGunDollarSignOffsetY)) // Location of $ sign
		pdf.SetFont(SPBTagFont, "", LongGunDollarFontSize)
		pdf.Cell(45, 55, "$")

		pdf.SetXY((zero.X + LongGunPriceOffsetX), (zero.Y + LongGunPriceOffsetY))
		// fontSizeOffset is used to scale down the LongGunPriceFontSize based on number of digits in the price
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
		// draw a handgun tag

		// needed to set up image embed
		var opt gofpdf.ImageOptions

		// Add the SPB logo
		pdf.ImageOptions(".\\logos\\shootpointblank.png", (zero.X + HandgunSPBLogoOffsetX), (zero.Y + HandgunSPBLogoOffsetY), HandgunSPBLogoWidth, 0, false, opt, 0, "")

		// get logo path
		LogoPath := (GetLogoImagePath(t.Manufacturer))

		pdf.ImageOptions(LogoPath, (zero.X + HandgunManufacturerLogoOffsetX), (zero.Y + HandgunManufacturerLogoOffsetY), HandgunManufacturerLogoWidth, HandgunManufacturerLogoHeight, false, opt, 0, "")

		pdf.SetXY((zero.X + HandgunDollarSignOffsetX), (zero.Y + HandgunDollarSignOffsetY)) // Location of $ sign
		pdf.SetFont(SPBTagFont, "", HandgunDollarFontSize)
		pdf.Cell(18, 28, "$")

		pdf.SetXY((zero.X + HandgunPriceOffsetX), (zero.Y + HandgunPriceOffsetY))
		// fontSizeOffset is used to scale down the HandgunPriceFontSize based on number of digits in the price
		var fontSizeOffset float64
		if len(t.Price) <= 3 {
			fontSizeOffset = 0
		} else if len(t.Price) == 4 {
			fontSizeOffset = 12
		} else if len(t.Price) == 5 {
			fontSizeOffset = 24
		} else if len(t.Price) == 6 {
			fontSizeOffset = 30
		} else if len(t.Price) >= 7 {
			fontSizeOffset = 28
		}
		pdf.SetFont(SPBTagFont, "", (HandgunPriceFontSize - fontSizeOffset))
		// NOTE: Remember to move the cell sizes to constants
		pdf.CellFormat(110, 45, t.Price, "", 0, "R", false, 0, "")
		pdf.SetFont(SPBTagFont, "", HandgunCentsFontSize)
		// Subtracting a few points from the font size superscripts the 99 cent slightly
		pdf.Write((HandgunCentsFontSize - 16), "99")

		pdf.SetXY((zero.X + HandgunModelOffsetX), (zero.Y + HandgunModelOffsetY))
		// Scale down the HandgunModelFontSize based on number of characters in the model name
		// This hand tuned offset will gracefully handle around ~44 characters before it starts overflowing, but may collide with the caliber cell if the caliber is lengthy
		if len(t.Model) <= 12 {
			fontSizeOffset = 0
		} else if len(t.Model) > 12 && len(t.Model) < 18 {
			fontSizeOffset = 2
		} else if len(t.Model) > 18 {
			fontSizeOffset = 5
		}
		pdf.SetFont(SPBTagFont, "", (HandgunModelFontSize - fontSizeOffset))
		pdf.CellFormat(60, 8, t.Model, "", 0, "L", false, 0, "")

		pdf.SetXY((zero.X + HandgunCaliberOffsetX), (zero.Y + HandgunCaliberOffsetY))
		// Scale down the HandgunCaliberFontSize based on number of characters in the caliber designation
		pdf.SetFont(SPBTagFont, "", HandgunCaliberFontSize)
		pdf.CellFormat(40, 8, t.Caliber, "", 0, "R", false, 0, "")

		// draw the red separator bar
		pdf.MoveTo((zero.X + HandgunSeparatorBarOffsetX), (zero.Y + HandgunSeparatorBarOffsetY))
		pdf.LineTo((zero.X + HandgunSeparatorBarLength), (zero.Y + HandgunSeparatorBarOffsetY))
		pdf.ClosePath()
		pdf.SetLineWidth(SeparatorBarLineWidth)
		pdf.SetDrawColor(SeparatorBarColorR, SeparatorBarColorB, SeparatorBarColorG)
		pdf.DrawPath("D")
	}

}
