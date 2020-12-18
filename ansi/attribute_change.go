package ansi

import (
	"fmt"
	"strings"

	"github.com/atrico-go/console/ansi/color"
)

type AttributeChange interface {
	// Apply to a string
	ApplyTo(str string) string
	// Get the ansi code for this set
	GetCodeString() string
	// Get the underlying codes
	GetCodes() []int
}

var ResetAll = attributes([]int{0})

// ----------------------------------------------------------------------------------------------------------------------------
// Implementation
// ----------------------------------------------------------------------------------------------------------------------------
type attributes []int

func (a attributes) ApplyTo(str string) string {
	text := strings.Builder{}
	text.WriteString(a.GetCodeString())
	text.WriteString(str)
	return text.String()
}

func (a attributes) GetCodeString() string {
	return createAnsiCode(a)
}

func (a attributes) GetCodes() []int {
	return a
}

// ----------------------------------------------------------------------------------------------------------------------------
// internal
// ----------------------------------------------------------------------------------------------------------------------------
var escape = '\x9b'
var escapeStr = string(escape)

func newAttributeChange(codes []int) AttributeChange {
	return attributes(codes)
}

func createAnsiCode(codes []int) string {
	text := strings.Builder{}
	if len(codes) > 0 {
		sep := escapeStr
		for _, code := range codes {
			text.WriteString(fmt.Sprintf("%s%d", sep, code))
			sep = ";"
		}
		text.WriteString("m")
	}
	return text.String()
}

func getDeltaCodes(oldAttribs, newAttribs Attributes) []int {
	codes := make([]int, 0, 2)
	if code, required := colorModificationCode(oldAttribs.Foreground, newAttribs.Foreground); required {
		codes = append(codes, code)
	}
	// Handle background as foreground
	if code, required := colorModificationCode(oldAttribs.Background, newAttribs.Background); required {
		codes = append(codes, int(convertForegroundToBackgroundColor(color.Color(code))))
	}
	return codes
}

func colorModificationCode(oldColor, newColor color.Color) (code int, required bool) {
	// No change
	if oldColor == newColor {
		return 0, false
	}
	// Reset color (already removed none->none above)
	if newColor == color.None {
		return resetColorCode, true
	}
	return int(newColor), true
}

func modifyAttributes(attributes Attributes, delta AttributeChange) Attributes {
	newAttribs := attributes
	for _,code := range delta.GetCodes() {
		// Reset foreground/background
		if code == resetColorCode {
			newAttribs.Foreground = color.None
			continue
		}
		if code == int(convertForegroundToBackgroundColor(color.Color(resetColorCode))) {
			newAttribs.Background = color.None
			continue
		}
		col := color.Color(code)
		// Foreground
		if isForegroundColor(col) {
			newAttribs.Foreground = col
			continue
		}
		// Background
		if isBackgroundColor(col) {
			newAttribs.Background = convertBackgroundToForegroundColor(col)
			continue
		}
	}
	return newAttribs
}

//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
// func (a attributes) GetDeltaCodeString(delta Attributes) string {
// 	codes := getDeltaCodes(a, newAttributes(delta.GetCodes()))
// 	return createAnsiCode(codes)
// }
//
//
//
//
// func (a Attributes) getCodes() []int {
// 	codes := make([]int, 0, 2)
// 	if a.Foreground != color.None {
// 		codes = append(codes, int(a.Foreground))
// 	}
// 	if a.Background != color.None {
// 		codes = append(codes, convertForegroundToBackgroundColor(int(a.Background)))
// 	}
// 	return codes
// }
//
//
//
// func (a *attributes) activateResets() {
// 	if a.foreground == int(color.Reset) {
// 		a.foreground = noCode
// 	}
//}