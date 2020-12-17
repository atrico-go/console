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
	// Apply to a string (with reset afterwards)
	ApplyTo(str string, resetAtEnd bool) string
	// Merge a new set of attributes
	Merge(attributes Attributes) (newAttributes Attributes, changed bool)
	// Make reset attributes for this set
	Reset() Attributes
}

type AttributeBuilder interface {
	WithForeground(color color.Color) AttributeBuilder
	WithBackground(color color.Color) AttributeBuilder
	Build() Attributes
}

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

func (a attributes) Merge(attributes Attributes) (newAttributes Attributes, changed bool) {
	newAttributesB := NewAttributeBuilder()
	// Foreground
	fore := a.foreground
	if attributes.Foreground() != color.None {
		fore = attributes.Foreground()
		if attributes.Foreground() == color.Reset {
			fore = color.None
		}
	}
	newAttributesB.WithForeground(fore)
	// Background
	back := a.background
	if attributes.Background() != color.None {
		back = attributes.Background()
		if attributes.Background() == color.Reset {
			back = color.None
		}
	}
	newAttributesB.WithBackground(back)
	newAttributes = newAttributesB.Build()
	if newAttributes != a {
		return newAttributes, true
	}
	return a, false
}

func (a attributes) ApplyTo(str string, resetAtEnd bool) string {
	text := strings.Builder{}
	codes := getCodes(a)
	text.WriteString(createAnsiCode(codes))
	text.WriteString(str)
	if resetAtEnd {
		resetAttribs := a.Reset()
		codes = getCodes(resetAttribs)
		text.WriteString(createAnsiCode(codes))
	}
	return text.String()
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
	if attr.Foreground() != color.None {
		codes = append(codes, int(attr.Foreground()))
	}
	if attr.Background() != color.None {
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
