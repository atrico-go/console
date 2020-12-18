package unit_tests

import (
	"fmt"
	"strings"
	"testing"

	"github.com/atrico-go/testing/assert"
	"github.com/atrico-go/testing/is"

	"github.com/atrico-go/console/ansi"
	"github.com/atrico-go/console/ansi/color"
)

func Test_AttributesParse_NoAttributes(t *testing.T) {
	// Arrange
	str := randomValues.String()

	// Act
	parsed := ansi.ParseString(str)

	// Assert
	assert.Assert(t).That(len(parsed), is.EqualTo(1), "One entry")
	entry1 := parsed[0]
	assert.Assert(t).That(entry1.String, is.EqualTo(str), "Correct string")
	assert.Assert(t).That(entry1.Attributes.Foreground, is.EqualTo(color.None), "No foreground")
	assert.Assert(t).That(entry1.Attributes.Background, is.EqualTo(color.None), "No background")
}

func Test_AttributesParse_ForegroundColor(t *testing.T) {
	// Arrange
	fore := randomColour()
	attribs := ansi.Attributes{Foreground: fore, Background: color.None}.SetThis()
	raw := randomValues.String()
	str := attribs.ApplyTo(raw)
	fmt.Println(str)
	fmt.Println(ansi.ResetAll.GetCodeString())

	// Act
	parsed := ansi.ParseString(str)

	// Assert
	assert.Assert(t).That(len(parsed), is.EqualTo(1), "One entry")
	entry1 := parsed[0]
	assert.Assert(t).That(entry1.String, is.EqualTo(raw), "Correct string")
	assert.Assert(t).That(entry1.Attributes.Foreground, is.EqualTo(fore), "Correct foreground")
	assert.Assert(t).That(entry1.Attributes.Background, is.EqualTo(color.None), "No background")
}

func Test_AttributesParse_BackgroundColor(t *testing.T) {
	// Arrange
	back := randomColour()
	attribs := ansi.Attributes{Foreground: color.None, Background: back}.SetThis()
	raw := randomValues.String()
	str := attribs.ApplyTo(raw)
	fmt.Println(str)
	fmt.Println(ansi.ResetAll.GetCodeString())

	// Act
	parsed := ansi.ParseString(str)

	// Assert
	assert.Assert(t).That(len(parsed), is.EqualTo(1), "One entry")
	entry1 := parsed[0]
	assert.Assert(t).That(entry1.String, is.EqualTo(raw), "Correct string")
	assert.Assert(t).That(entry1.Attributes.Foreground, is.EqualTo(color.None), "No foreground")
	assert.Assert(t).That(entry1.Attributes.Background, is.EqualTo(back), "Correct background")
}

func Test_AttributesParse_BothColors(t *testing.T) {
	// Arrange
	fore := randomColour()
	back := randomColour()
	attribs := ansi.Attributes{Foreground: fore, Background: back}.SetThis()
	raw := randomValues.String()
	str := attribs.ApplyTo(raw)
	fmt.Println(str)
	fmt.Println(ansi.ResetAll)

	// Act
	parsed := ansi.ParseString(str)

	// Assert
	assert.Assert(t).That(len(parsed), is.EqualTo(1), "One entry")
	entry1 := parsed[0]
	assert.Assert(t).That(entry1.String, is.EqualTo(raw), "Correct string")
	assert.Assert(t).That(entry1.Attributes.Foreground, is.EqualTo(fore), "Correct foreground")
	assert.Assert(t).That(entry1.Attributes.Background, is.EqualTo(back), "Correct background")
}

func Test_AttributesParse_MultipleColorsDeltaTo(t *testing.T) {
	// Arrange
	fore1 := color.Red
	fore2 := color.Green
	back2 := color.Blue
	back3 := color.Yellow
	attribs1 := ansi.Attributes{Foreground: fore1, Background: color.None}
	attribs2 := ansi.Attributes{Foreground: fore2, Background: back2}
	attribs3 := ansi.Attributes{Foreground: color.None, Background: back3}
	delta1 := attribs1.SetThis()
	delta2 := attribs1.CreateDeltaTo(attribs2)
	delta3 := attribs2.CreateDeltaTo(attribs3)
	raw1 := randomValues.String()
	raw2 := randomValues.String()
	raw3 := randomValues.String()
	strB := strings.Builder{}
	strB.WriteString(delta1.ApplyTo(raw1))
	strB.WriteString(delta2.ApplyTo(raw2))
	strB.WriteString(delta3.ApplyTo(raw3))
	str := strB.String()
	fmt.Println(str)
	fmt.Println(ansi.ResetAll)

	// Act
	parsed := ansi.ParseString(str)

	// Assert
	assert.Assert(t).That(len(parsed), is.EqualTo(3), "3 entries")
	// 1
	assert.Assert(t).That(parsed[0].String, is.EqualTo(raw1), "1: Correct string")
	assert.Assert(t).That(parsed[0].Attributes.Foreground, is.EqualTo(fore1), "1: Correct foreground")
	assert.Assert(t).That(parsed[0].Attributes.Background, is.EqualTo(color.None), "1: No background")
	// 2
	assert.Assert(t).That(parsed[1].String, is.EqualTo(raw2), "2: Correct string")
	assert.Assert(t).That(parsed[1].Attributes.Foreground, is.EqualTo(fore2), "2: Correct foreground")
	assert.Assert(t).That(parsed[1].Attributes.Background, is.EqualTo(back2), "2: Correct background")
	// 3
	assert.Assert(t).That(parsed[2].String, is.EqualTo(raw3), "3: Correct string")
	assert.Assert(t).That(parsed[2].Attributes.Foreground, is.EqualTo(color.None), "3: No foreground")
	assert.Assert(t).That(parsed[2].Attributes.Background, is.EqualTo(back3), "3: Correct background")
}

func Test_AttributesParse_MultipleColorsDeltaFrom(t *testing.T) {
	// Arrange
	fore1 := color.Red
	fore2 := color.Green
	back2 := color.Blue
	back3 := color.Yellow
	attribs1 := ansi.Attributes{Foreground: fore1, Background: color.None}
	attribs2 := ansi.Attributes{Foreground: fore2, Background: back2}
	attribs3 := ansi.Attributes{Foreground: color.None, Background: back3}
	delta1 := attribs1.SetThis()
	delta2 := attribs2.CreateDeltaFrom(attribs1)
	delta3 := attribs3.CreateDeltaFrom(attribs2)
	raw1 := randomValues.String()
	raw2 := randomValues.String()
	raw3 := randomValues.String()
	strB := strings.Builder{}
	strB.WriteString(delta1.ApplyTo(raw1))
	strB.WriteString(delta2.ApplyTo(raw2))
	strB.WriteString(delta3.ApplyTo(raw3))
	str := strB.String()
	fmt.Println(str)
	fmt.Println(ansi.ResetAll)

	// Act
	parsed := ansi.ParseString(str)

	// Assert
	assert.Assert(t).That(len(parsed), is.EqualTo(3), "3 entries")
	// 1
	assert.Assert(t).That(parsed[0].String, is.EqualTo(raw1), "1: Correct string")
	assert.Assert(t).That(parsed[0].Attributes.Foreground, is.EqualTo(fore1), "1: Correct foreground")
	assert.Assert(t).That(parsed[0].Attributes.Background, is.EqualTo(color.None), "1: No background")
	// 2
	assert.Assert(t).That(parsed[1].String, is.EqualTo(raw2), "2: Correct string")
	assert.Assert(t).That(parsed[1].Attributes.Foreground, is.EqualTo(fore2), "2: Correct foreground")
	assert.Assert(t).That(parsed[1].Attributes.Background, is.EqualTo(back2), "2: Correct background")
	// 3
	assert.Assert(t).That(parsed[2].String, is.EqualTo(raw3), "3: Correct string")
	assert.Assert(t).That(parsed[2].Attributes.Foreground, is.EqualTo(color.None), "3: No foreground")
	assert.Assert(t).That(parsed[2].Attributes.Background, is.EqualTo(back3), "3: Correct background")
}
