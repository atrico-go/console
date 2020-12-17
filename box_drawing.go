package console

import (
	"fmt"
	"strings"
)

type BoxType int

const (
	// Type of line
	BoxNone   BoxType = 0
	BoxSingle BoxType = 1
	BoxDouble BoxType = 2
	BoxHeavy  BoxType = 3
)

func (bt BoxType) String() string {
	switch bt {
	case BoxNone:
		return "None"
	case BoxSingle:
		return "Single"
	case BoxDouble:
		return "Double"
	case BoxHeavy:
		return "Heavy"
	}
	panic("Unknown box type")
}

type BoxParts struct {
	Up    BoxType
	Down  BoxType
	Left  BoxType
	Right BoxType
}

func GetBoxChar(up bool, down bool, left bool, right bool, boxType BoxType) (char rune, ok bool) {
	upBt := ConditionalBoxType(up, boxType, BoxNone)
	downBt := ConditionalBoxType(down, boxType, BoxNone)
	leftBt := ConditionalBoxType(left, boxType, BoxNone)
	rightBt := ConditionalBoxType(right, boxType, BoxNone)
	return GetBoxCharMixed(BoxParts{upBt, downBt, leftBt, rightBt})
}

func GetBoxCharMixed(parts BoxParts) (char rune, ok bool) {
	char, ok = boxParts[parts]
	return char, ok
}

func MustGetBoxChar(up bool, down bool, left bool, right bool, boxType BoxType) rune {
	upBt := ConditionalBoxType(up, boxType, BoxNone)
	downBt := ConditionalBoxType(down, boxType, BoxNone)
	leftBt := ConditionalBoxType(left, boxType, BoxNone)
	rightBt := ConditionalBoxType(right, boxType, BoxNone)
	return MustGetBoxCharMixed(BoxParts{upBt, downBt, leftBt, rightBt})
}

func MustGetBoxCharMixed(parts BoxParts) rune {
	if char, ok := GetBoxCharMixed(parts); ok {
		return char
	}
	msg := strings.Builder{}
	msg.WriteString("Box drawing character not found: ")
	sep := ""
	if parts.Up != BoxNone {
		msg.WriteString(fmt.Sprintf("%sup=%s", sep, parts.Up))
		sep = ", "
	}
	if parts.Down != BoxNone {
		msg.WriteString(fmt.Sprintf("%sparts.Down=%s", sep, parts.Down))
		sep = ", "
	}
	if parts.Left != BoxNone {
		msg.WriteString(fmt.Sprintf("%sparts.Left=%s", sep, parts.Left))
		sep = ", "
	}
	if parts.Right != BoxNone {
		msg.WriteString(fmt.Sprintf("%sparts.Right=%s", sep, parts.Right))
	}
	panic(msg.String())
}

func GetHorizontal(bt BoxType) rune {
	return MustGetBoxChar(false, false, true, true, bt)
}

func GetVertical(bt BoxType) rune {
	return MustGetBoxChar(true, true, false, false, bt)
}

func ConditionalBoxType(c bool, t BoxType, f BoxType) BoxType {
	if c {
		return t
	} else {
		return f
	}
}

func (bt BoxType) HeavyIf(c bool) BoxType {
	return ConditionalBoxType(c, BoxHeavy, bt)
}

// ----------------------------------------------------------------------------------------------------------------------------
// Implementation
// ----------------------------------------------------------------------------------------------------------------------------

var boxParts = map[BoxParts]rune{
	// Space
	BoxParts{BoxNone, BoxNone, BoxNone, BoxNone}: ' ',
	// Half lines
	BoxParts{BoxSingle, BoxNone, BoxNone, BoxNone}: '╵',
	BoxParts{BoxNone, BoxSingle, BoxNone, BoxNone}: '╷',
	BoxParts{BoxNone, BoxNone, BoxSingle, BoxNone}: '╴',
	BoxParts{BoxNone, BoxNone, BoxNone, BoxSingle}: '╶',
	BoxParts{BoxHeavy, BoxNone, BoxNone, BoxNone}:  '╹',
	BoxParts{BoxNone, BoxHeavy, BoxNone, BoxNone}:  '╻',
	BoxParts{BoxNone, BoxNone, BoxHeavy, BoxNone}:  '╸',
	BoxParts{BoxNone, BoxNone, BoxNone, BoxHeavy}:  '╺',
	// Full lines
	BoxParts{BoxSingle, BoxSingle, BoxNone, BoxNone}: '│',
	BoxParts{BoxNone, BoxNone, BoxSingle, BoxSingle}: '─',
	BoxParts{BoxDouble, BoxDouble, BoxNone, BoxNone}: '║',
	BoxParts{BoxNone, BoxNone, BoxDouble, BoxDouble}: '═',
	BoxParts{BoxHeavy, BoxHeavy, BoxNone, BoxNone}:   '┃',
	BoxParts{BoxNone, BoxNone, BoxHeavy, BoxHeavy}:   '━',
	BoxParts{BoxHeavy, BoxSingle, BoxNone, BoxNone}:  '╿',
	BoxParts{BoxSingle, BoxHeavy, BoxNone, BoxNone}:  '╽',
	BoxParts{BoxNone, BoxNone, BoxHeavy, BoxSingle}:  '╾',
	BoxParts{BoxNone, BoxNone, BoxSingle, BoxHeavy}:  '╼',
	// Up-Left
	BoxParts{BoxSingle, BoxNone, BoxSingle, BoxNone}: '┘',
	BoxParts{BoxDouble, BoxNone, BoxDouble, BoxNone}: '╝',
	BoxParts{BoxDouble, BoxNone, BoxSingle, BoxNone}: '╜',
	BoxParts{BoxSingle, BoxNone, BoxDouble, BoxNone}: '╛',
	BoxParts{BoxHeavy, BoxNone, BoxHeavy, BoxNone}:   '┛',
	BoxParts{BoxHeavy, BoxNone, BoxSingle, BoxNone}:  '┚',
	BoxParts{BoxSingle, BoxNone, BoxHeavy, BoxNone}:  '┙',
	// Up-Right
	BoxParts{BoxSingle, BoxNone, BoxNone, BoxSingle}: '└',
	BoxParts{BoxDouble, BoxNone, BoxNone, BoxDouble}: '╚',
	BoxParts{BoxDouble, BoxNone, BoxNone, BoxSingle}: '╙',
	BoxParts{BoxSingle, BoxNone, BoxNone, BoxDouble}: '╘',
	BoxParts{BoxHeavy, BoxNone, BoxNone, BoxHeavy}:   '┗',
	BoxParts{BoxHeavy, BoxNone, BoxNone, BoxSingle}:  '┖',
	BoxParts{BoxSingle, BoxNone, BoxNone, BoxHeavy}:  '┕',
	// Down-Left
	BoxParts{BoxNone, BoxSingle, BoxSingle, BoxNone}: '┐',
	BoxParts{BoxNone, BoxDouble, BoxDouble, BoxNone}: '╗',
	BoxParts{BoxNone, BoxDouble, BoxSingle, BoxNone}: '╖',
	BoxParts{BoxNone, BoxSingle, BoxDouble, BoxNone}: '╕',
	BoxParts{BoxNone, BoxHeavy, BoxHeavy, BoxNone}:   '┓',
	BoxParts{BoxNone, BoxHeavy, BoxSingle, BoxNone}:  '┒',
	BoxParts{BoxNone, BoxSingle, BoxHeavy, BoxNone}:  '┑',
	// Down-Right
	BoxParts{BoxNone, BoxSingle, BoxNone, BoxSingle}: '┌',
	BoxParts{BoxNone, BoxDouble, BoxNone, BoxDouble}: '╔',
	BoxParts{BoxNone, BoxDouble, BoxNone, BoxSingle}: '╓',
	BoxParts{BoxNone, BoxSingle, BoxNone, BoxDouble}: '╒',
	BoxParts{BoxNone, BoxHeavy, BoxNone, BoxHeavy}:   '┏',
	BoxParts{BoxNone, BoxHeavy, BoxNone, BoxSingle}:  '┎',
	BoxParts{BoxNone, BoxSingle, BoxNone, BoxHeavy}:  '┍',
	// T-Up
	BoxParts{BoxSingle, BoxNone, BoxSingle, BoxSingle}: '┴',
	BoxParts{BoxDouble, BoxNone, BoxDouble, BoxDouble}: '╩',
	BoxParts{BoxSingle, BoxNone, BoxDouble, BoxDouble}: '╧',
	BoxParts{BoxDouble, BoxNone, BoxSingle, BoxSingle}: '╨',
	BoxParts{BoxHeavy, BoxNone, BoxHeavy, BoxHeavy}:    '┻',
	BoxParts{BoxSingle, BoxNone, BoxHeavy, BoxHeavy}:   '┷',
	BoxParts{BoxHeavy, BoxNone, BoxSingle, BoxSingle}:  '┸',
	BoxParts{BoxHeavy, BoxNone, BoxSingle, BoxHeavy}:   '┺',
	BoxParts{BoxSingle, BoxNone, BoxHeavy, BoxSingle}:  '┵',
	BoxParts{BoxHeavy, BoxNone, BoxHeavy, BoxSingle}:   '┹',
	BoxParts{BoxSingle, BoxNone, BoxSingle, BoxHeavy}:  '┶',
	// T-Down
	BoxParts{BoxNone, BoxSingle, BoxSingle, BoxSingle}: '┬',
	BoxParts{BoxNone, BoxDouble, BoxDouble, BoxDouble}: '╦',
	BoxParts{BoxNone, BoxSingle, BoxDouble, BoxDouble}: '╤',
	BoxParts{BoxNone, BoxDouble, BoxSingle, BoxSingle}: '╥',
	BoxParts{BoxNone, BoxHeavy, BoxHeavy, BoxHeavy}:    '┳',
	BoxParts{BoxNone, BoxSingle, BoxHeavy, BoxHeavy}:   '┯',
	BoxParts{BoxNone, BoxHeavy, BoxSingle, BoxSingle}:  '┰',
	BoxParts{BoxNone, BoxHeavy, BoxSingle, BoxHeavy}:   '┲',
	BoxParts{BoxNone, BoxSingle, BoxHeavy, BoxSingle}:  '┭',
	BoxParts{BoxNone, BoxHeavy, BoxHeavy, BoxSingle}:   '┱',
	BoxParts{BoxNone, BoxSingle, BoxSingle, BoxHeavy}:  '┮',
	// T-Left
	BoxParts{BoxSingle, BoxSingle, BoxSingle, BoxNone}: '┤',
	BoxParts{BoxDouble, BoxDouble, BoxSingle, BoxNone}: '╣',
	BoxParts{BoxDouble, BoxDouble, BoxSingle, BoxNone}: '╢',
	BoxParts{BoxSingle, BoxSingle, BoxDouble, BoxNone}: '╡',
	BoxParts{BoxHeavy, BoxHeavy, BoxHeavy, BoxNone}:    '┫',
	BoxParts{BoxSingle, BoxHeavy, BoxHeavy, BoxNone}:   '┨',
	BoxParts{BoxHeavy, BoxSingle, BoxSingle, BoxNone}:  '┥',
	BoxParts{BoxSingle, BoxHeavy, BoxHeavy, BoxNone}:   '┪',
	BoxParts{BoxHeavy, BoxSingle, BoxSingle, BoxNone}:  '┦',
	BoxParts{BoxHeavy, BoxSingle, BoxHeavy, BoxNone}:   '┩',
	BoxParts{BoxSingle, BoxHeavy, BoxSingle, BoxNone}:  '┧',
	// T-Right
	BoxParts{BoxSingle, BoxSingle, BoxNone, BoxSingle}: '├',
	BoxParts{BoxDouble, BoxDouble, BoxNone, BoxDouble}: '╠',
	BoxParts{BoxDouble, BoxDouble, BoxNone, BoxSingle}: '╟',
	BoxParts{BoxSingle, BoxSingle, BoxNone, BoxDouble}: '╞',
	BoxParts{BoxHeavy, BoxHeavy, BoxNone, BoxHeavy}:    '┣',
	BoxParts{BoxHeavy, BoxHeavy, BoxNone, BoxSingle}:   '┠',
	BoxParts{BoxSingle, BoxSingle, BoxNone, BoxHeavy}:  '┝',
	BoxParts{BoxSingle, BoxHeavy, BoxNone, BoxHeavy}:   '┢',
	BoxParts{BoxHeavy, BoxSingle, BoxNone, BoxSingle}:  '┞',
	BoxParts{BoxHeavy, BoxSingle, BoxNone, BoxHeavy}:   '┡',
	BoxParts{BoxSingle, BoxHeavy, BoxNone, BoxSingle}:  '┟',
	// Cross
	BoxParts{BoxSingle, BoxSingle, BoxSingle, BoxSingle}: '┼',
	BoxParts{BoxDouble, BoxDouble, BoxDouble, BoxDouble}: '╬',
	BoxParts{BoxSingle, BoxSingle, BoxDouble, BoxDouble}: '╪',
	BoxParts{BoxDouble, BoxDouble, BoxSingle, BoxSingle}: '╫',
	BoxParts{BoxHeavy, BoxHeavy, BoxHeavy, BoxHeavy}:     '╋',
	BoxParts{BoxSingle, BoxSingle, BoxHeavy, BoxHeavy}:   '┿',
	BoxParts{BoxHeavy, BoxHeavy, BoxSingle, BoxSingle}:   '╂',
	BoxParts{BoxSingle, BoxHeavy, BoxHeavy, BoxHeavy}:    '╈',
	BoxParts{BoxHeavy, BoxSingle, BoxSingle, BoxSingle}:  '╀',
	BoxParts{BoxHeavy, BoxSingle, BoxHeavy, BoxHeavy}:    '╇',
	BoxParts{BoxSingle, BoxHeavy, BoxSingle, BoxSingle}:  '╁',
	BoxParts{BoxHeavy, BoxHeavy, BoxSingle, BoxHeavy}:    '╊',
	BoxParts{BoxSingle, BoxSingle, BoxHeavy, BoxSingle}:  '┽',
	BoxParts{BoxHeavy, BoxHeavy, BoxHeavy, BoxSingle}:    '╉',
	BoxParts{BoxSingle, BoxSingle, BoxSingle, BoxHeavy}:  '┾',
	BoxParts{BoxSingle, BoxHeavy, BoxSingle, BoxHeavy}:   '╆',
	BoxParts{BoxHeavy, BoxSingle, BoxHeavy, BoxSingle}:   '╃',
	BoxParts{BoxSingle, BoxHeavy, BoxHeavy, BoxSingle}:   '╅',
	BoxParts{BoxHeavy, BoxSingle, BoxSingle, BoxHeavy}:   '╄',
}
