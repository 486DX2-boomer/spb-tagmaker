package main

import (
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

// SPBLogoWidth defines the SPB logo width in pts
const SPBLogoWidth float64 = 129

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

// DrawLongGunTag accepts a single tag and draws it to the coordinate specified
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
	pdf.ImageOptions(".\\logos\\shootpointblank.png", (zero.X + LongGunSPBLogoOffsetX), (zero.Y + LongGunSPBLogoOffsetY), SPBLogoWidth, 0, false, opt, 0, "")

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
