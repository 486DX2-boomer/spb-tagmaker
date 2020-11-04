package main

// ListenPort defines the listen port of the http server
const ListenPort string = ":8080"

// Tag size. Big = 3x5". Small = 2x3"
const (
	Big   = false
	Small = true
)

// SPBTagFont is the font family used for all price tags
const SPBTagFont string = "Franklin Gothic Medium"

// Long gun 3x5 dimensions in points: 360 x 216 pt
// Hand gun 2x3 dimensions in points: 216 x 144 pt

// Long gun constants

// LongGunSPBLogoWidth defines the SPB logo width in pts
const LongGunSPBLogoWidth float64 = 129

// LongGunManufacturerLogoWidth defines the width of the manufacturer's logo in pts
const LongGunManufacturerLogoWidth float64 = 96

// LongGunManufacturerLogoHeight defines the height of the manufacturer's logo in pts
const LongGunManufacturerLogoHeight float64 = 60

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
	SeparatorBarColorG        int     = 87
	SeparatorBarColorB        int     = 42
)
