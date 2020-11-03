package main

import (
	"strings"

	"github.com/jung-kurt/gofpdf"
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

// NewTag constructs and returns a Tag struct
func NewTag(manufacturer string, model string, caliber string, new bool, price string, size Size) Tag {
	var t Tag

	t.Manufacturer = manufacturer
	t.Model = model
	t.Caliber = caliber
	t.New = new
	t.Price = price
	t.TagSize = size

	return t
}

// Draw the tag to the position coordinate provided
func (t Tag) Draw(zero Coord, pdf *gofpdf.Fpdf) {

	// Strip the .99 cents from the price field if present (only dollar amount is needed)
	if strings.Contains(t.Price, ".") {
		t.Price = strings.Split(t.Price, ".")[0]
	}
	// Strip the $ sign from the price field if present
	if strings.Contains(t.Price, "$") {
		t.Price = strings.Split(t.Price, "$")[1]
	}

	if t.TagSize == Big && t.New == true {
		// draw a long gun tag

		// needed to set up image embed
		var opt gofpdf.ImageOptions

		// Add the SPB logo
		pdf.ImageOptions("./logos/shootpointblank.png", (zero.X + LongGunSPBLogoOffsetX), (zero.Y + LongGunSPBLogoOffsetY), LongGunSPBLogoWidth, 0, false, opt, 0, "")

		// get logo path
		LogoPath := (GetLogoImagePath(t.Manufacturer))

		pdf.ImageOptions(LogoPath, (zero.X + LongGunManufacturerLogoOffsetX), (zero.Y + LongGunManufacturerLogoOffsetY), LongGunManufacturerLogoWidth, LongGunManufacturerLogoHeight, false, opt, 0, "")

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
		pdf.SetDrawColor(SeparatorBarColorR, SeparatorBarColorG, SeparatorBarColorB)
		pdf.DrawPath("D")

	} else if t.TagSize == Small && t.New == true {
		// draw a handgun tag

		// needed to set up image embed
		var opt gofpdf.ImageOptions

		// Add the SPB logo
		pdf.ImageOptions("./logos/shootpointblank.png", (zero.X + HandgunSPBLogoOffsetX), (zero.Y + HandgunSPBLogoOffsetY), HandgunSPBLogoWidth, 0, false, opt, 0, "")

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
		pdf.SetDrawColor(SeparatorBarColorR, SeparatorBarColorG, SeparatorBarColorB)
		pdf.DrawPath("D")
	} else if t.TagSize == Big && t.New == false {
		// draw a used long gun tag

		// needed to set up image embed
		var opt gofpdf.ImageOptions

		// Draw Preowned Firearm in top left
		pdf.SetFillColor(254, 0, 0)
		pdf.SetDrawColor(254, 0, 0)
		pdf.MoveTo(zero.X, zero.Y)
		pdf.LineTo((zero.X + 128), zero.Y)
		pdf.CurveTo(zero.X+150, zero.Y+80, zero.X, zero.Y+70)
		pdf.ClosePath()
		pdf.DrawPath("F")
		pdf.SetTextColor(255, 255, 255)
		pdf.SetXY(zero.X+8, zero.Y+20)
		pdf.SetFont(SPBTagFont, "", 20)
		pdf.Write(0, "Preowned")
		pdf.SetXY(zero.X+8, zero.Y+44)
		pdf.Write(0, "Firearm")
		pdf.SetTextColor(0, 0, 0)

		// get manufacturer logo path
		LogoPath := (GetLogoImagePath(t.Manufacturer))
		pdf.ImageOptions(LogoPath, (zero.X + LongGunManufacturerLogoOffsetX), (zero.Y + LongGunManufacturerLogoOffsetY), LongGunManufacturerLogoWidth, LongGunManufacturerLogoHeight, false, opt, 0, "")

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

		// On used guns, the caliber is written in white on a red semi circle
		pdf.SetFillColor(254, 0, 0)
		pdf.SetDrawColor(254, 0, 0)
		pdf.MoveTo(zero.X+360, zero.Y+216)
		pdf.LineTo(zero.X+360, zero.Y+164)
		pdf.CurveTo(zero.X+238, zero.Y+140, zero.X+220, zero.Y+216)
		pdf.ClosePath()
		pdf.DrawPath("F")

		pdf.SetTextColor(255, 255, 255)
		pdf.SetXY((zero.X + LongGunCaliberOffsetX), (zero.Y + LongGunCaliberOffsetY))
		// Scale down the LongGunCaliberFontSize based on number of characters in the caliber designation
		if len(t.Caliber) <= 6 {
			fontSizeOffset = 0
		} else if len(t.Caliber) > 6 && len(t.Caliber) < 12 {
			fontSizeOffset = 4
		} else if len(t.Caliber) > 12 {
			fontSizeOffset = 8
		}
		pdf.SetFont(SPBTagFont, "", LongGunCaliberFontSize-fontSizeOffset)
		pdf.CellFormat(113, 18, t.Caliber, "", 0, "R", false, 0, "")
		pdf.SetTextColor(0, 0, 0)

	} else if t.TagSize == Small && t.New == false {
		// draw a used handgun tag

		// needed to set up image embed
		var opt gofpdf.ImageOptions

		// Draw Preowned Firearm in top left
		pdf.SetFillColor(254, 0, 0)
		pdf.SetDrawColor(254, 0, 0)
		pdf.MoveTo(zero.X, zero.Y)
		pdf.LineTo((zero.X + 90), zero.Y)
		pdf.CurveTo(zero.X+100, zero.Y+64, (zero.X), zero.Y+50)
		pdf.ClosePath()
		pdf.DrawPath("F")
		pdf.SetTextColor(255, 255, 255)
		pdf.SetXY(zero.X+4, zero.Y+14)
		pdf.SetFont(SPBTagFont, "", 14)
		pdf.Write(0, "Preowned")
		pdf.SetXY(zero.X+4, zero.Y+32)
		pdf.Write(0, "Firearm")
		pdf.SetTextColor(0, 0, 0)

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

		// On used guns, the caliber is written in white on a red semi circle

		pdf.SetFillColor(254, 0, 0)
		pdf.SetDrawColor(254, 0, 0)
		pdf.MoveTo(zero.X+216, zero.Y+144)
		pdf.LineTo(zero.X+216, zero.Y+110)
		pdf.CurveTo(zero.X+140, zero.Y+90, zero.X+112, zero.Y+144)
		pdf.ClosePath()
		pdf.DrawPath("F")

		pdf.SetTextColor(255, 255, 255)
		pdf.SetXY((zero.X + HandgunCaliberOffsetX), (zero.Y + HandgunCaliberOffsetY))
		// Scale down the HandgunCaliberFontSize based on number of characters in the caliber designation
		if len(t.Caliber) <= 6 {
			fontSizeOffset = 0
		} else if len(t.Caliber) > 9 && len(t.Caliber) < 14 {
			fontSizeOffset = 2
		} else if len(t.Caliber) > 14 {
			fontSizeOffset = 4
		}
		pdf.SetFont(SPBTagFont, "", HandgunCaliberFontSize-fontSizeOffset)
		pdf.CellFormat(40, 8, t.Caliber, "", 0, "R", false, 0, "")
		pdf.SetTextColor(0, 0, 0)

	}

}
