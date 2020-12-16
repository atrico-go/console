package console

import (
	"fmt"
	"testing"

	. "github.com/atrico-go/testing/assert"
	"github.com/atrico-go/testing/is"
)

func Test_BoxDrawing_GetHorizontal(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc.BoxType), func(t *testing.T) {
			// Act
			char := GetHorizontal(tc.BoxType)
			// Assert
			Assert(t).That(char, is.EqualTo(tc.horizontal), "Correct char")
		})
	}
}

func Test_BoxDrawing_GetVertical(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc.BoxType), func(t *testing.T) {
			// Act
			char := GetVertical(tc.BoxType)
			// Assert
			Assert(t).That(char, is.EqualTo(tc.vertical), "Correct char")
		})
	}
}

func Test_BoxDrawing_TopLeft(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc.BoxType), func(t *testing.T) {
			// Act
			char, ok := GetBoxChar(false, true, false, true, tc.BoxType)
			// Assert
			Assert(t).That(ok, is.True, "Char found")
			Assert(t).That(char, is.EqualTo(tc.topLeft), "Correct char")
		})
	}
}

func Test_BoxDrawing_TopRight(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc.BoxType), func(t *testing.T) {
			// Act
			char, ok := GetBoxChar(false, true, true, false, tc.BoxType)
			// Assert
			Assert(t).That(ok, is.True, "Char found")
			Assert(t).That(char, is.EqualTo(tc.topRight), "Correct char")
		})
	}
}

func Test_BoxDrawing_BottomLeft(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc.BoxType), func(t *testing.T) {
			// Act
			char, ok := GetBoxChar(true, false, false, true, tc.BoxType)
			// Assert
			Assert(t).That(ok, is.True, "Char found")
			Assert(t).That(char, is.EqualTo(tc.bottomLeft), "Correct char")
		})
	}
}

func Test_BoxDrawing_BottomRight(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc.BoxType), func(t *testing.T) {
			// Act
			char, ok := GetBoxChar(true, false, true, false, tc.BoxType)
			// Assert
			Assert(t).That(ok, is.True, "Char found")
			Assert(t).That(char, is.EqualTo(tc.bottomRight), "Correct char")
		})
	}
}

func Test_BoxDrawing_ConditionalType(t *testing.T) {
	for _, tcT := range testCases {
		for _, tcF := range testCases {
			t.Run(fmt.Sprintf("%v / %v", tcT.BoxType, tcF.BoxType), func(t *testing.T) {
				// Act
				btT := ConditionalBoxType(true, tcT.BoxType, tcF.BoxType)
				btF := ConditionalBoxType(false, tcT.BoxType, tcF.BoxType)
				// Assert
				Assert(t).That(btT, is.EqualTo(tcT.BoxType), "True")
				Assert(t).That(btF, is.EqualTo(tcF.BoxType), "False")
			})
		}
	}
}

func Test_BoxDrawing_HeavyIf(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc.BoxType), func(t *testing.T) {
			// Act
			btF := tc.BoxType.HeavyIf(false)
			btT := tc.BoxType.HeavyIf(true)
			// Assert
			Assert(t).That(btF, is.EqualTo(tc.BoxType), "No Change")
			Assert(t).That(btT, is.EqualTo(BoxHeavy), "Correct char")
		})
	}
}

type testCase struct {
	BoxType
	horizontal  rune
	vertical    rune
	topLeft     rune
	topRight    rune
	bottomLeft  rune
	bottomRight rune
}

var testCases = []testCase{
	{BoxNone, ' ', ' ', ' ', ' ', ' ', ' '},
	{BoxSingle, '─', '│', '┌', '┐', '└', '┘'},
	{BoxDouble, '═', '║', '╔', '╗', '╚', '╝'},
	{BoxHeavy, '━', '┃', '┏', '┓', '┗', '┛'},
}
