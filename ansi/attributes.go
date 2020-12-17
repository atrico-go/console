package ansi

import (
	"fmt"
	"strings"

	"github.com/atrico-go/console/ansi/color"
)

var Escape = '\x9b'
var EscapeStr = string(Escape)
var ResetAllStr = createAnsiCode([]int{0})

type Attributes interface {
	Foreground() color.Color
	Background() color.Color
	// Get the ansi code
	GetCodeString() string
	// Apply to a string (with reset afterwards)
	ApplyTo(str string, resetAtEnd bool) string
	// Modify to a new set of attributes
	// Optimize the changes required
	Modify(attributes Attributes) (newAttributes Attributes, changed bool)
	// Make reset attributes for this set
	Reset() Attributes
}

type AttributeBuilder interface {
	AsDelta() AttributeBuilder
	WithForeground(color color.Color) AttributeBuilder
	WithBackground(color color.Color) AttributeBuilder
	Build() Attributes
}

// New attributes set
// Defaults to no attributes
func NewAttributes(codes ...int) Attributes {
	builder := NewAttributeBuilder()
	for _, code := range codes {
		// Foreground
		if color, err := color.ParseColor(code); err == nil {
			builder.WithForeground(color)
			continue
		}
		// Background
		if color, err := color.ParseColor(code - 10); err == nil {
			builder.WithBackground(color)
			continue
		}
	}
	return builder.Build()
}

// New attribute modification
// Defaults ot all attributes ignored
func NewAttributesDelta() Attributes {
	return NewAttributeBuilder().
		AsDelta().
		Build()
}

func NewAttributeBuilder() AttributeBuilder {
	return &attributes{color.None, color.None}
}

// ----------------------------------------------------------------------------------------------------------------------------
// Implementation
// ----------------------------------------------------------------------------------------------------------------------------
type attributes struct {
	foreground color.Color
	background color.Color
}

func (a attributes) String() string {
	return fmt.Sprintf("ATTR: fg = %s, bg = %s", a.foreground, a.background)
}

func (a attributes) Foreground() color.Color {
	return a.foreground
}

func (a attributes) Background() color.Color {
	return a.background
}

func (a attributes) GetCodeString() string {
	codes := getCodes(a)
	return createAnsiCode(codes)
}

func (a attributes) ApplyTo(str string, resetAtEnd bool) string {
	text := strings.Builder{}
	text.WriteString(a.GetCodeString())
	text.WriteString(str)
	if resetAtEnd {
		resetAttribs := a.Reset()
		text.WriteString(resetAttribs.GetCodeString())
	}
	return text.String()
}

func (a attributes) Modify(attributes Attributes) (newAttributes Attributes, changed bool) {
	newAttributes = NewAttributeBuilder().
		WithForeground(modifyColor(a.foreground, attributes.Foreground())).
		WithBackground(modifyColor(a.background, attributes.Background())).
		Build()
	if newAttributes != a {
		return newAttributes, true
	}
	return a, false
}

func modifyColor(oldColor color.Color, newColor color.Color) color.Color {
	// Ignore == no change
	if newColor == color.Ignore {
		return oldColor
	}
	// Existing ignore can only be changed to actual color (no reset or none)
	if oldColor == color.Ignore && (newColor == color.None || newColor == color.Reset) {
		return oldColor
	}
	// Reset returns color to none
	if newColor == color.Reset {
		return color.None
	}
	return newColor
}

func (a attributes) Reset() Attributes {
	reset := NewAttributeBuilder()
	if a.Foreground() != color.None && a.Foreground() != color.Reset {
		reset.WithForeground(color.Reset)
	}
	if a.Background() != color.None && a.Background() != color.Reset {
		reset.WithBackground(color.Reset)
	}
	return reset.Build()
}

func (a *attributes) AsDelta() AttributeBuilder {
	return a.
		WithForeground(color.Ignore).
		WithBackground(color.Ignore)
}


func (a *attributes) WithForeground(color color.Color) AttributeBuilder {
	a.foreground = color
	return a
}

func (a *attributes) WithBackground(color color.Color) AttributeBuilder {
	a.background = color
	return a
}

func (a *attributes) Build() Attributes {
	return *a
}

// ----------------------------------------------------------------------------------------------------------------------------
// internal
// ----------------------------------------------------------------------------------------------------------------------------

func getCodes(attr Attributes) []int {
	codes := make([]int, 0, 2)
	if int(attr.Foreground()) > 0 {
		codes = append(codes, int(attr.Foreground()))
	}
	if int(attr.Background()) > 0 {
		codes = append(codes, int(attr.Background()+10))
	}
	return codes
}

func createAnsiCode(codes []int) string {
	text := strings.Builder{}
	if len(codes) > 0 {
		sep := EscapeStr
		for _, code := range codes {
			text.WriteString(fmt.Sprintf("%s%d", sep, code))
			sep = ";"
		}
		text.WriteString("m")
	}
	return text.String()
}
