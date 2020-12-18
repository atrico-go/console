package color

import (
	"fmt"
)

type Color int

const (
	// Normal
	Black     Color = 30
	Red       Color = 31
	Green     Color = 32
	Yellow    Color = 33
	Blue      Color = 34
	Magenta   Color = 35
	Cyan      Color = 36
	LightGrey Color = 37
	// Bright
	DarkGrey      Color = 90
	BrightRed     Color = 91
	BrightGreen   Color = 92
	BrightYellow  Color = 93
	BrightBlue    Color = 94
	BrightMagenta Color = 95
	BrightCyan    Color = 96
	White         Color = 97
	// No color
	None		Color = -1
)

func (c Color) String() string {
	switch c {
	case Black:
		return "Black"
	case Red:
		return "Red"
	case Green:
		return "Green"
	case Yellow:
		return "Yellow"
	case Blue:
		return "Blue"
	case Magenta:
		return "Magenta"
	case Cyan:
		return "Cyan"
	case LightGrey:
		return "LightGrey"
	case DarkGrey:
		return "DarkGrey"
	case BrightRed:
		return "BrightRed"
	case BrightGreen:
		return "BrightGreen"
	case BrightYellow:
		return "BrightYellow"
	case BrightBlue:
		return "BrightBlue"
	case BrightMagenta:
		return "BrightMagenta"
	case BrightCyan:
		return "BrightCyan"
	case White:
		return "White"
	case None:
		return "None"
	}
	return fmt.Sprintf("Unknown Color (%d)", int(c))
}
