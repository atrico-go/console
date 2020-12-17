package unit_tests

import (
	"fmt"
	"strings"
	"testing"

	"github.com/atrico-go/testing/assert"
	"github.com/atrico-go/testing/is"
	"github.com/atrico-go/testing/random"

	"github.com/atrico-go/console/ansi"
	"github.com/atrico-go/console/ansi/color"
)

var randomValues = random.NewValueGeneratorBuilder().
	WithDefaultStringLength(5).
	Build()

func Test_AttributesParse_NoAttributes(t *testing.T) {
	// Arrange
	str := randomValues.String()

	// Act
	parsed := ansi.ParseString(str)

	// Assert
	assert.Assert(t).That(len(parsed), is.EqualTo(1), "One entry")
	entry1 := parsed[0]
	assert.Assert(t).That(entry1.String, is.EqualTo(str), "Correct string")
	assert.Assert(t).That(entry1.Attributes.Foreground(), is.EqualTo(color.None), "No foreground")
	assert.Assert(t).That(entry1.Attributes.Background(), is.EqualTo(color.None), "No background")
}

func Test_AttributesParse_ForegroundColor(t *testing.T) {
	// Arrange
	fore := randomColour()
	attribs := ansi.NewAttributeBuilder().
		WithForeground(fore).
		Build()
	impl := func(t *testing.T, reset bool) {
		// Arrange
		raw := randomValues.String()
		str := attribs.ApplyTo(raw, reset)
		fmt.Println(str)
		fmt.Println(ansi.ResetAllStr)
		// Act
		parsed := ansi.ParseString(str)
		// Assert
		assert.Assert(t).That(len(parsed), is.EqualTo(1), "One entry")
		entry1 := parsed[0]
		assert.Assert(t).That(entry1.String, is.EqualTo(raw), "Correct string")
		assert.Assert(t).That(entry1.Attributes.Foreground(), is.EqualTo(fore), "Correct foreground")
		assert.Assert(t).That(entry1.Attributes.Background(), is.EqualTo(color.None), "No background")
	}
	t.Run("Reset", func(t *testing.T) { impl(t, true) })
	t.Run("No Reset", func(t *testing.T) { impl(t, false) })
}

func Test_AttributesParse_BackgroundColor(t *testing.T) {
	// Arrange
	back := randomColour()
	attribs := ansi.NewAttributeBuilder().
		WithBackground(back).
		Build()
	impl := func(t *testing.T, reset bool) {
		// Arrange
		raw := randomValues.String()
		str := attribs.ApplyTo(raw, reset)
		fmt.Println(str)
		fmt.Println(ansi.ResetAllStr)
		// Act
		parsed := ansi.ParseString(str)
		// Assert
		assert.Assert(t).That(len(parsed), is.EqualTo(1), "One entry")
		entry1 := parsed[0]
		assert.Assert(t).That(entry1.String, is.EqualTo(raw), "Correct string")
		assert.Assert(t).That(entry1.Attributes.Foreground(), is.EqualTo(color.None), "No foreground")
		assert.Assert(t).That(entry1.Attributes.Background(), is.EqualTo(back), "Correct background")
	}
	t.Run("Reset", func(t *testing.T) { impl(t, true) })
	t.Run("No Reset", func(t *testing.T) { impl(t, false) })
}

func Test_AttributesParse_BothColors(t *testing.T) {
	// Arrange
	fore := randomColour()
	back := randomColour()
	attribs := ansi.NewAttributeBuilder().
		WithForeground(fore).
		WithBackground(back).
		Build()
	impl := func(t *testing.T, reset bool) {
		// Arrange
		raw := randomValues.String()
		str := attribs.ApplyTo(raw, reset)
		fmt.Println(str)
		fmt.Println(ansi.ResetAllStr)
		// Act
		parsed := ansi.ParseString(str)
		// Assert
		assert.Assert(t).That(len(parsed), is.EqualTo(1), "One entry")
		entry1 := parsed[0]
		assert.Assert(t).That(entry1.String, is.EqualTo(raw), "Correct string")
		assert.Assert(t).That(entry1.Attributes.Foreground(), is.EqualTo(fore), "Correct foreground")
		assert.Assert(t).That(entry1.Attributes.Background(), is.EqualTo(back), "Correct background")
	}
	t.Run("Reset", func(t *testing.T) { impl(t, true) })
	t.Run("No Reset", func(t *testing.T) { impl(t, false) })
}

func Test_AttributesParse_MultipleColors(t *testing.T) {
	// Arrange
	fore1 := color.Red
	fore2 := color.Green
	back2 := color.Blue
	back3 := color.Yellow
	attribs1 := ansi.NewAttributeBuilder().
		WithForeground(fore1).
		Build()
	attribs2 := ansi.NewAttributeBuilder().
		WithForeground(fore2).
		WithBackground(back2).
		Build()
	attribs3 := ansi.NewAttributeBuilder().
		WithForeground(color.Reset).
		WithBackground(back3).
		Build()
	raw1 := randomValues.String()
	raw2 := randomValues.String()
	raw3 := randomValues.String()
	strB := strings.Builder{}
	strB.WriteString(attribs1.ApplyTo(raw1, false))
	strB.WriteString(attribs2.ApplyTo(raw2, false))
	strB.WriteString(attribs3.ApplyTo(raw3, true))
	str := strB.String()
	fmt.Println(str)

	// Act
	parsed := ansi.ParseString(str)
	// Assert
	assert.Assert(t).That(len(parsed), is.EqualTo(3), "3 entries")
	// 1
	assert.Assert(t).That(parsed[0].String, is.EqualTo(raw1), "1: Correct string")
	assert.Assert(t).That(parsed[0].Attributes.Foreground(), is.EqualTo(fore1), "1: Correct foreground")
	assert.Assert(t).That(parsed[0].Attributes.Background(), is.EqualTo(color.None), "1: No background")
	// 2
	assert.Assert(t).That(parsed[1].String, is.EqualTo(raw2), "2: Correct string")
	assert.Assert(t).That(parsed[1].Attributes.Foreground(), is.EqualTo(fore2), "2: Correct foreground")
	assert.Assert(t).That(parsed[1].Attributes.Background(), is.EqualTo(back2), "2: Correct background")
	// 3
	assert.Assert(t).That(parsed[2].String, is.EqualTo(raw3), "3: Correct string")
	assert.Assert(t).That(parsed[2].Attributes.Foreground(), is.EqualTo(color.None), "3: No foreground")
	assert.Assert(t).That(parsed[2].Attributes.Background(), is.EqualTo(back3), "3: Correct background")
}
