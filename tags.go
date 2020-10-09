package main

// Size type is used in the Tag type to store whether the tag is large format or small format.
type Size bool

// Tag size. Big = 3x5". Small = 2x3"
const (
	Big   = false
	Small = true
)

// All of these are in pts
// Long gun tag 1 0,0 coordinate : 125, 72
// Long gun tag 2 0,0 coordinate : 125, 287
// Long gun tag 3 0,0 coordinate : 125, 504

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

// Tag type defines all relevant information to be printed out. It should also include offsets for where each string should be positioned on the tag.
type Tag struct {
	Price        float64
	Manufacturer string
	Model        string
	Caliber      string
	TagSize      Size // Big or Small
}

// Constructor function for tag should go here
