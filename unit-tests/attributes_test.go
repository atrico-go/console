package unit_tests

import (
	"fmt"
	"testing"

	"github.com/atrico-go/testing/assert"
	"github.com/atrico-go/testing/is"

	"github.com/atrico-go/console/ansi"
	"github.com/atrico-go/console/ansi/color"
)

func Test_Attributes_NoAttributes(t *testing.T) {
	// Arrange

	// Act
	attribs := ansi.NoAttributes
	fmt.Println(attribs)

	// Assert
	assert.Assert(t).That(attribs.Foreground, is.EqualTo(color.None), "Foreground")
	assert.Assert(t).That(attribs.Background, is.EqualTo(color.None), "Background")
}

func Test_Attributes_Set(t *testing.T) {
	// Arrange
	original := ansi.NoAttributes
	fore := randomColour()
	back := randomColour()
	target := ansi.Attributes{Foreground: fore, Background: back}
	fmt.Printf("%s => %s\n", original, target)

	// Act
	delta := target.SetThis()
	newAttributes := original.Modify(delta)
	fmt.Printf("= %s\n", newAttributes)

	// Assert
	assert.Assert(t).That(newAttributes.Foreground, is.EqualTo(fore), "Foreground - no change")
	assert.Assert(t).That(newAttributes.Background, is.EqualTo(back), "Background - no change")
}

func Test_Attributes_Modify(t *testing.T) {
	// Arrange
	fore1 := randomColour()
	back1 := randomColour()
	original := ansi.Attributes{Foreground: fore1, Background: back1}
	fore2 := fore1
	for fore2 == fore1 {
		fore2 = randomColour()
	}
	back2 := back1
	for back2 == back1 {
		back2 = randomColour()
	}
	target := ansi.Attributes{Foreground: fore2, Background: back2}
	fmt.Printf("%s => %s\n", original, target)

	// Act
	delta := original.CreateDeltaTo(target)
	newAttributes := original.Modify(delta)
	fmt.Printf("= %s\n", newAttributes)

	// Assert
	assert.Assert(t).That(newAttributes.Foreground, is.EqualTo(fore2), "Foreground")
	assert.Assert(t).That(newAttributes.Background, is.EqualTo(back2), "Background")
}

func Test_Attributes_Reset(t *testing.T) {
	// Arrange
	fore := randomColour()
	back := randomColour()
	original := ansi.Attributes{Foreground: fore, Background: back}
	fmt.Printf("%s => %s\n", original, ansi.NoAttributes)

	// Act
	delta := original.ResetThis()
	newAttributes := original.Modify(delta)
	fmt.Printf("= %s\n", newAttributes)

	// Assert
	assert.Assert(t).That(newAttributes.Foreground, is.EqualTo(color.None), "Foreground")
	assert.Assert(t).That(newAttributes.Background, is.EqualTo(color.None), "Background")
}

func randomColour() color.Color {
	val := randomValues.IntBetween(30, 38)
	if randomValues.Bool() {
		// Bright
		val = val + 60
	}
	return color.Color(val)
}
