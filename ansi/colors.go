package ansi

import (
	"github.com/atrico-go/console/ansi/color"
)

func toBackgroundColor(val int) (newCol color.Color, success bool) {
	newCol = color.Color(val)
	if success = isBackgroundColor(newCol); !success {
		if success = isForegroundColor(newCol); success {
			newCol = convertForegroundToBackgroundColor(newCol)
		}
	}
	return newCol, success
}

func toForegroundColor(val int) (newCol color.Color, success bool) {
	newCol = color.Color(val)
	if success = isForegroundColor(newCol); !success {
		if success = isBackgroundColor(newCol); success {
			newCol = convertBackgroundToForegroundColor(newCol)
		}
	}
	return newCol, success
}

// ----------------------------------------------------------------------------------------------------------------------------
// internal
// ----------------------------------------------------------------------------------------------------------------------------
var resetColorCode = 39

func isForegroundColor(col color.Color) bool {
	val := int(col)
	return (30 <= val && val <= 37) || (90 <= val && val <= 97)
}
func isBackgroundColor(col color.Color) bool {
	return isForegroundColor(convertBackgroundToForegroundColor(col))
}
func convertBackgroundToForegroundColor(col color.Color) color.Color {
	return color.Color(int(col) - 10)
}
func convertForegroundToBackgroundColor(col color.Color) color.Color {
	return color.Color(int(col) + 10)
}
