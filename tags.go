package main

// Size type is used in the Tag type to store whether the tag is large format or small format.
type Size bool

// Tag size. Big = 3x5". Small = 2x3"
const (
	Big   = false
	Small = true
)

// Font sizes and font family
const (
	LongGunDollarFontSize  float64 = 60
	LongGunPriceFontSize   float64 = 113
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
