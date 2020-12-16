package unit_tests

import (
	"fmt"
	"testing"

	. "github.com/atrico-go/testing/assert"
	"github.com/atrico-go/testing/is"

	"github.com/atrico-go/console/box_drawing"
)

func Test_BoxDrawing_Lookup(t *testing.T) {
	for _, tc := range lookupTestCases {
		t.Run(fmt.Sprintf("%v", tc.rune), func(t *testing.T) {
			// Act
			parts, ok := box_drawing.Lookup(tc.rune)
			// Assert
			Assert(t).That(ok, is.True, "Char found")
			Assert(t).That(parts.Up, is.EqualTo(tc.up), "Correct Up")
			Assert(t).That(parts.Down, is.EqualTo(tc.down), "Correct Down")
			Assert(t).That(parts.Left, is.EqualTo(tc.left), "Correct Left")
			Assert(t).That(parts.Right, is.EqualTo(tc.right), "Correct Right")
		})
	}
}

func Test_BoxDrawing_LookupNotFound(t *testing.T) {
	for _, tc := range lookupNotFoundTestCases {
		t.Run(fmt.Sprintf("%v", tc), func(t *testing.T) {
			// Act
			_, ok := box_drawing.Lookup(tc)
			// Assert
			Assert(t).That(ok, is.False, "Char not found")
		})
	}
}

type testCaseLookup struct {
	rune
	up    box_drawing.BoxType
	down  box_drawing.BoxType
	left  box_drawing.BoxType
	right box_drawing.BoxType
}

var lookupTestCases = []testCaseLookup{
	{' ', box_drawing.BoxNone, box_drawing.BoxNone, box_drawing.BoxNone, box_drawing.BoxNone},
	{'─', box_drawing.BoxNone, box_drawing.BoxNone, box_drawing.BoxSingle, box_drawing.BoxSingle},
	{'╟', box_drawing.BoxDouble, box_drawing.BoxDouble, box_drawing.BoxNone, box_drawing.BoxSingle},
	{'┺', box_drawing.BoxHeavy, box_drawing.BoxNone, box_drawing.BoxSingle, box_drawing.BoxHeavy},
}

var lookupNotFoundTestCases = []rune{
	'a', '!', '=',
}
