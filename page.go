package main

// LongGunPage stores 3 tags per page. Used when sorting tags on pdf generation and calling DrawPage method on each
type LongGunPage struct {
	Tags [3]Tag
}

// HandgunPage stores 3 tags per page. Used when sorting tags on pdf generation and calling DrawPage method on each
type HandgunPage struct {
	Tags [10]Tag
}
