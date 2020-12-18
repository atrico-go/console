package ansi

import (
	"fmt"

	"github.com/atrico-go/console/ansi/color"
)

type Attributes struct {
	Foreground color.Color
	Background color.Color
}

var NoAttributes = Attributes{color.None, color.None}

func (a Attributes) String() string {
	return fmt.Sprintf("[%s,%s]", a.Foreground, a.Background)
}

// Get a change to create this state from no attributes
func (a Attributes) SetThis() AttributeChange {
	return a.CreateDeltaFrom(NoAttributes)
}

// Get a change to reset this state
func (a Attributes) ResetThis() AttributeChange {
	return a.CreateDeltaTo(NoAttributes)
}

// Get a change to change from this set to the new one
func (a Attributes) CreateDeltaTo(newAttributes Attributes) AttributeChange {
	return newAttributeChange(getDeltaCodes(a, newAttributes))
}

// Get a change to change to this set from the old one
func (a Attributes) CreateDeltaFrom(oldAttributes Attributes) AttributeChange {
	return newAttributeChange(getDeltaCodes(oldAttributes, a))
}

// Modify these attributes with the change
func (a Attributes) Modify(delta AttributeChange) Attributes {
	return modifyAttributes(a, delta)
}
