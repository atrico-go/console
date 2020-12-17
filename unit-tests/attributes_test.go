package unit_tests

import (
	"testing"

	"github.com/atrico-go/testing/assert"
	"github.com/atrico-go/testing/is"

	"github.com/atrico-go/console/ansi"
	"github.com/atrico-go/console/ansi/color"
)

func Test_Attributes_Default(t *testing.T) {
	// Arrange

	// Act
	attribs := ansi.NewAttributes()

	// Assert
	assert.Assert(t).That(attribs.Foreground(), is.EqualTo(color.None), "Foreground")
	assert.Assert(t).That(attribs.Background(), is.EqualTo(color.None), "Background")
}

func Test_Attributes_Colours(t *testing.T) {
	// Arrange
	fore := randomColour()
	back := randomColour()

	// Act
	attribs := ansi.NewAttributeBuilder().
		WithForeground(fore).
		WithBackground(back).
		Build()

	// Assert
	assert.Assert(t).That(attribs.Foreground(), is.EqualTo(fore), "Foreground")
	assert.Assert(t).That(attribs.Background(), is.EqualTo(back), "Background")
}

func Test_Attributes_ModifyDefault(t *testing.T) {
	// Arrange
	fore := randomColour()
	back := randomColour()
	original := ansi.NewAttributeBuilder().
		WithForeground(fore).
		WithBackground(back).
		Build()

	// Act
	delta := ansi.NewAttributesDelta()
	newAttributes, changed := original.Modify(delta)

	// Assert
	assert.Assert(t).That(changed, is.False, "No change")
	assert.Assert(t).That(newAttributes.Foreground(), is.EqualTo(fore), "Foreground - no change")
	assert.Assert(t).That(newAttributes.Background(), is.EqualTo(back), "Background - no change")
}

func Test_Attributes_ModifyForeground(t *testing.T) {
	// Arrange
	fore := randomColour()
	back := randomColour()
	original := ansi.NewAttributeBuilder().
		WithForeground(fore).
		WithBackground(back).
		Build()

	// Act
	fore2 := fore
	for fore2 == fore {
		fore2 = randomColour()
	}
	delta := ansi.NewAttributeBuilder().
		AsDelta().
		WithForeground(fore2).
		Build()
	newAttributes, changed := original.Modify(delta)

	// Assert
	assert.Assert(t).That(changed, is.True, "Change")
	assert.Assert(t).That(newAttributes.Foreground(), is.EqualTo(fore2), "Foreground")
	assert.Assert(t).That(newAttributes.Background(), is.EqualTo(back), "Background - no change")
}

func Test_Attributes_ModifyBackground(t *testing.T) {
	// Arrange
	fore := randomColour()
	back := randomColour()
	original := ansi.NewAttributeBuilder().
		WithForeground(fore).
		WithBackground(back).
		Build()

	// Act
	back2 := back
	for back2 == back {
		back2 = randomColour()
	}
	delta := ansi.NewAttributeBuilder().
		AsDelta().
		WithBackground(back2).
		Build()
	newAttributes, changed := original.Modify(delta)

	// Assert
	assert.Assert(t).That(changed, is.True, "Change")
	assert.Assert(t).That(newAttributes.Foreground(), is.EqualTo(fore), "Foreground - no change")
	assert.Assert(t).That(newAttributes.Background(), is.EqualTo(back2), "Background")
}

func Test_Attributes_ModifyReset(t *testing.T) {
	// Arrange
	fore := randomColour()
	back := randomColour()
	original := ansi.NewAttributeBuilder().
		WithForeground(fore).
		WithBackground(back).
		Build()

	// Act
	delta := ansi.NewAttributeBuilder().
		AsDelta().
		WithBackground(color.Reset).
		Build()
	newAttributes, changed := original.Modify(delta)

	// Assert
	assert.Assert(t).That(changed, is.True, "Change")
	assert.Assert(t).That(newAttributes.Foreground(), is.EqualTo(fore), "Foreground - no change")
	assert.Assert(t).That(newAttributes.Background(), is.EqualTo(color.None), "Background")
}

func Test_Attributes_ModifyResetNoColour(t *testing.T) {
	// Arrange
	fore := randomColour()
	original := ansi.NewAttributeBuilder().
		WithForeground(fore).
		WithBackground(color.None).
		Build()

	// Act
	delta := ansi.NewAttributeBuilder().
		AsDelta().
		WithBackground(color.Reset).
		Build()
	_, changed := original.Modify(delta)

	// Assert
	assert.Assert(t).That(changed, is.False, "No change")
}

func Test_Attributes_ModifyResetIgnored(t *testing.T) {
	// Arrange
	fore := randomColour()
	original := ansi.NewAttributeBuilder().
		WithForeground(fore).
		Build()

	// Act
	delta := ansi.NewAttributeBuilder().
		AsDelta().
		WithBackground(color.Reset).
		Build()
	_, changed := original.Modify(delta)

	// Assert
	assert.Assert(t).That(changed, is.False, "No change")
}

func randomColour() color.Color {
	val := randomValues.IntBetween(30, 38)
	if randomValues.Bool() {
		// Bright
		val = val + 60
	}
	return color.Color(val)
}
