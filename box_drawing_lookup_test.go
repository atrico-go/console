package console

import (
	"fmt"
	"testing"

	. "github.com/atrico-go/testing/assert"
	"github.com/atrico-go/testing/is"
)

func Test_BoxDrawing_Lookup(t *testing.T) {
	for _, tc := range lookupTestCases {
		t.Run(fmt.Sprintf("%v", tc.rune), func(t *testing.T) {
			// Act
			parts, ok := Lookup(tc.rune)
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
			_, ok := Lookup(tc)
			// Assert
			Assert(t).That(ok, is.False, "Char not found")
		})
	}
}

type testCaseLookup struct {
	rune
	up    BoxType
	down  BoxType
	left  BoxType
	right BoxType
}

var lookupTestCases = []testCaseLookup{
	{' ', BoxNone, BoxNone, BoxNone, BoxNone},
	{'─', BoxNone, BoxNone, BoxSingle, BoxSingle},
	{'╟', BoxDouble, BoxDouble, BoxNone, BoxSingle},
	{'┺', BoxHeavy, BoxNone, BoxSingle, BoxHeavy},
}

var lookupNotFoundTestCases = []rune{
	'a', '!', '=',
}
