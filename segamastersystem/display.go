package sms

const (
	DISPLAY_WIDTH      = 256
	DISPLAY_HEIGHT     = 192
	DISPLAY_WIDTH_LOG2 = 8
	DISPLAY_SIZE       = DISPLAY_WIDTH * DISPLAY_HEIGHT
	BORDER_LEFT_RIGHT  = 64
	BORDER_TOP_BOTTOM  = 48
	SCREEN_WIDTH       = DISPLAY_WIDTH + BORDER_LEFT_RIGHT*2
	SCREEN_HEIGHT      = DISPLAY_HEIGHT + BORDER_TOP_BOTTOM*2
)

type DisplayData [DISPLAY_WIDTH * DISPLAY_HEIGHT]byte

type PaletteValue struct {
	index   byte
	r, g, b byte
}

// Interface for rendering backend
type DisplayLoop interface {
	Display() chan<- *DisplayData
	WritePalette() chan<- PaletteValue
	UpdateBorder() chan<- byte
}
