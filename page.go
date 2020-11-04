package main

// Coord is used to store x, y coordinate pairs
type Coord struct {
	X float64
	Y float64
}

// GetLongGunPageCoord returns an array of coordinates helpful when placing long gun tags on long gun pages
func GetLongGunPageCoord() [3]Coord {
	var c [3]Coord
	c[0].X = LongGunPageTag1PositionX
	c[0].Y = LongGunPageTag1PositionY

	c[1].X = LongGunPageTag2PositionX
	c[1].Y = LongGunPageTag2PositionY

	c[2].X = LongGunPageTag3PositionX
	c[2].Y = LongGunPageTag3PositionY

	return c
}

// GetHandgunPageCoord returns an array of coordinates helpful when placing handgun tags on handgun pages
func GetHandgunPageCoord() [10]Coord {
	var c [10]Coord

	c[0].X = HandgunPageTag1PositionX
	c[0].Y = HandgunPageTag1PositionY

	c[1].X = HandgunPageTag2PositionX
	c[1].Y = HandgunPageTag2PositionY

	c[2].X = HandgunPageTag3PositionX
	c[2].Y = HandgunPageTag3PositionY

	c[3].X = HandgunPageTag4PositionX
	c[3].Y = HandgunPageTag4PositionY

	c[4].X = HandgunPageTag5PositionX
	c[4].Y = HandgunPageTag5PositionY

	c[5].X = HandgunPageTag6PositionX
	c[5].Y = HandgunPageTag6PositionY

	c[6].X = HandgunPageTag7PositionX
	c[6].Y = HandgunPageTag7PositionY

	c[7].X = HandgunPageTag8PositionX
	c[7].Y = HandgunPageTag8PositionY

	c[8].X = HandgunPageTag9PositionX
	c[8].Y = HandgunPageTag9PositionY

	c[9].X = HandgunPageTag10PositionX
	c[9].Y = HandgunPageTag10PositionY

	return c
}
