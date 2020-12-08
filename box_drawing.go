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

func GetBoxChar(up bool, down bool, left bool, right bool, boxType BoxType) (char string, ok bool) {
	upBt := ConditionalBoxType(up, boxType, BoxNone)
	downBt := ConditionalBoxType(down, boxType, BoxNone)
	leftBt := ConditionalBoxType(left, boxType, BoxNone)
	rightBt := ConditionalBoxType(right, boxType, BoxNone)
	return GetBoxCharMixed(upBt, downBt, leftBt, rightBt)
}

func GetBoxCharMixed(up BoxType, down BoxType, left BoxType, right BoxType) (char string, ok bool) {
	char, ok = boxParts[createDir(up, down, left, right)]
	return char, ok
}

func MustGetBoxChar(up bool, down bool, left bool, right bool, boxType BoxType) string {
	upBt := ConditionalBoxType(up, boxType, BoxNone)
	downBt := ConditionalBoxType(down, boxType, BoxNone)
	leftBt := ConditionalBoxType(left, boxType, BoxNone)
	rightBt := ConditionalBoxType(right, boxType, BoxNone)
	return MustGetBoxCharMixed(upBt, downBt, leftBt, rightBt)
}

func MustGetBoxCharMixed(up BoxType, down BoxType, left BoxType, right BoxType) string {
	if char, ok := GetBoxCharMixed(up, down, left, right); ok {
		return char
	}
	msg := strings.Builder{}
	msg.WriteString("Box drawing character not found: ")
	sep := ""
	if up != BoxNone {
		msg.WriteString(fmt.Sprintf("%sup=%s", sep, up))
		sep = ", "
	}
	if down != BoxNone {
		msg.WriteString(fmt.Sprintf("%sdown=%s", sep, down))
		sep = ", "
	}
	if left != BoxNone {
		msg.WriteString(fmt.Sprintf("%sleft=%s", sep, left))
		sep = ", "
	}
	if right != BoxNone {
		msg.WriteString(fmt.Sprintf("%sright=%s", sep, right))
	}
	panic(msg.String())
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

type direction uint32

func (d direction) getUp() BoxType {
	return d.getDir(0xff000000, 24)
}
func (d *direction) setUp(boxType BoxType) *direction {
	return d.setDir(0x00ffffff, 24, boxType)
}
func (d direction) getDown() BoxType {
	return d.getDir(0x00ff0000, 16)
}
func (d *direction) setDown(boxType BoxType) *direction {
	return d.setDir(0xff00ffff, 16, boxType)
}
func (d direction) getLeft() BoxType {
	return d.getDir(0x0000ff00, 8)
}
func (d *direction) setLeft(boxType BoxType) *direction {
	return d.setDir(0xffff00ff, 8, boxType)
}
func (d direction) getRight() BoxType {
	return d.getDir(0x000000ff, 0)
}
func (d *direction) setRight(boxType BoxType) *direction {
	return d.setDir(0xffffff00, 0, boxType)
}
func (d direction) getDir(mask uint32, shift int) BoxType {
	return BoxType((uint32(d) & mask) >> shift)
}
func (d *direction) setDir(mask uint32, shift int, boxType BoxType) *direction {
	*d = *d & direction(mask)
	*d = *d | direction(boxType<<shift)
	return d
}
func createDir(up BoxType, down BoxType, left BoxType, right BoxType) direction {
	d := direction(0)
	return *d.setUp(up).setDown(down).setLeft(left).setRight(right)
}

var boxParts = map[direction]string{
	// Space
	createDir(BoxNone, BoxNone, BoxNone, BoxNone): " ",
	// Half lines
	createDir(BoxSingle, BoxNone, BoxNone, BoxNone): "╵",
	createDir(BoxNone, BoxSingle, BoxNone, BoxNone): "╷",
	createDir(BoxNone, BoxNone, BoxSingle, BoxNone): "╴",
	createDir(BoxNone, BoxNone, BoxNone, BoxSingle): "╶",
	createDir(BoxHeavy, BoxNone, BoxNone, BoxNone):  "╹",
	createDir(BoxNone, BoxHeavy, BoxNone, BoxNone):  "╻",
	createDir(BoxNone, BoxNone, BoxHeavy, BoxNone):  "╸",
	createDir(BoxNone, BoxNone, BoxNone, BoxHeavy):  "╺",
	// Full lines
	createDir(BoxSingle, BoxSingle, BoxNone, BoxNone): "│",
	createDir(BoxNone, BoxNone, BoxSingle, BoxSingle): "─",
	createDir(BoxDouble, BoxDouble, BoxNone, BoxNone): "║",
	createDir(BoxNone, BoxNone, BoxDouble, BoxDouble): "═",
	createDir(BoxHeavy, BoxHeavy, BoxNone, BoxNone):   "┃",
	createDir(BoxNone, BoxNone, BoxHeavy, BoxHeavy):   "━",
	createDir(BoxHeavy, BoxSingle, BoxNone, BoxNone):  "╿ ",
	createDir(BoxSingle, BoxHeavy, BoxNone, BoxNone):  "╽",
	createDir(BoxNone, BoxNone, BoxHeavy, BoxSingle):  "╾",
	createDir(BoxNone, BoxNone, BoxSingle, BoxHeavy):  "╼",
	// Up-Left
	createDir(BoxSingle, BoxNone, BoxSingle, BoxNone): "┘",
	createDir(BoxDouble, BoxNone, BoxDouble, BoxNone): "╝",
	createDir(BoxDouble, BoxNone, BoxSingle, BoxNone): "╜",
	createDir(BoxSingle, BoxNone, BoxDouble, BoxNone): "╛",
	createDir(BoxHeavy, BoxNone, BoxHeavy, BoxNone):   "┛",
	createDir(BoxHeavy, BoxNone, BoxSingle, BoxNone):  "┚",
	createDir(BoxSingle, BoxNone, BoxHeavy, BoxNone):  "┙",
	// Up-Right
	createDir(BoxSingle, BoxNone, BoxNone, BoxSingle): "└",
	createDir(BoxDouble, BoxNone, BoxNone, BoxDouble): "╚",
	createDir(BoxDouble, BoxNone, BoxNone, BoxSingle): "╙",
	createDir(BoxSingle, BoxNone, BoxNone, BoxDouble): "╘",
	createDir(BoxHeavy, BoxNone, BoxNone, BoxHeavy):   "┗",
	createDir(BoxHeavy, BoxNone, BoxNone, BoxSingle):  "┖",
	createDir(BoxSingle, BoxNone, BoxNone, BoxHeavy):  "┕",
	// Down-Left
	createDir(BoxNone, BoxSingle, BoxSingle, BoxNone): "┐",
	createDir(BoxNone, BoxDouble, BoxDouble, BoxNone): "╗",
	createDir(BoxNone, BoxDouble, BoxSingle, BoxNone): "╖",
	createDir(BoxNone, BoxSingle, BoxDouble, BoxNone): "╕",
	createDir(BoxNone, BoxHeavy, BoxHeavy, BoxNone):   "┓",
	createDir(BoxNone, BoxHeavy, BoxSingle, BoxNone):  "┒",
	createDir(BoxNone, BoxSingle, BoxHeavy, BoxNone):  "┑",
	// Down-Right
	createDir(BoxNone, BoxSingle, BoxNone, BoxSingle): "┌",
	createDir(BoxNone, BoxDouble, BoxNone, BoxDouble): "╔",
	createDir(BoxNone, BoxDouble, BoxNone, BoxSingle): "╓",
	createDir(BoxNone, BoxSingle, BoxNone, BoxDouble): "╒",
	createDir(BoxNone, BoxHeavy, BoxNone, BoxHeavy):   "┏",
	createDir(BoxNone, BoxHeavy, BoxNone, BoxSingle):  "┎",
	createDir(BoxNone, BoxSingle, BoxNone, BoxHeavy):  "┍",
	// T-Up
	createDir(BoxSingle, BoxNone, BoxSingle, BoxSingle): "┴",
	createDir(BoxDouble, BoxNone, BoxDouble, BoxDouble): "╩",
	createDir(BoxSingle, BoxNone, BoxDouble, BoxDouble): "╧",
	createDir(BoxDouble, BoxNone, BoxSingle, BoxSingle): "╨",
	createDir(BoxHeavy, BoxNone, BoxHeavy, BoxHeavy):    "┻",
	createDir(BoxSingle, BoxNone, BoxHeavy, BoxHeavy):   "┷",
	createDir(BoxHeavy, BoxNone, BoxSingle, BoxSingle):  "┸",
	createDir(BoxHeavy, BoxNone, BoxSingle, BoxHeavy):   "┺",
	createDir(BoxSingle, BoxNone, BoxHeavy, BoxSingle):  "┵",
	createDir(BoxHeavy, BoxNone, BoxHeavy, BoxSingle):   "┹",
	createDir(BoxSingle, BoxNone, BoxSingle, BoxHeavy):  "┶",
	// T-Down
	createDir(BoxNone, BoxSingle, BoxSingle, BoxSingle): "┬",
	createDir(BoxNone, BoxDouble, BoxDouble, BoxDouble): "╦",
	createDir(BoxNone, BoxSingle, BoxDouble, BoxDouble): "╤",
	createDir(BoxNone, BoxDouble, BoxSingle, BoxSingle): "╥",
	createDir(BoxNone, BoxHeavy, BoxHeavy, BoxHeavy):    "┳",
	createDir(BoxNone, BoxSingle, BoxHeavy, BoxHeavy):   "┯",
	createDir(BoxNone, BoxHeavy, BoxSingle, BoxSingle):  "┰",
	createDir(BoxNone, BoxHeavy, BoxSingle, BoxHeavy):   "┲",
	createDir(BoxNone, BoxSingle, BoxHeavy, BoxSingle):  "┭",
	createDir(BoxNone, BoxHeavy, BoxHeavy, BoxSingle):   "┱",
	createDir(BoxNone, BoxSingle, BoxSingle, BoxHeavy):  "┮",
	// T-Left
	createDir(BoxSingle, BoxSingle, BoxSingle, BoxNone): "┤",
	createDir(BoxDouble, BoxDouble, BoxSingle, BoxNone): "╣",
	createDir(BoxDouble, BoxDouble, BoxSingle, BoxNone): "╢",
	createDir(BoxSingle, BoxSingle, BoxDouble, BoxNone): "╡",
	createDir(BoxHeavy, BoxHeavy, BoxHeavy, BoxNone):    "┫",
	createDir(BoxSingle, BoxHeavy, BoxHeavy, BoxNone):   "┨",
	createDir(BoxHeavy, BoxSingle, BoxSingle, BoxNone):  "┥",
	createDir(BoxSingle, BoxHeavy, BoxHeavy, BoxNone):   "┪",
	createDir(BoxHeavy, BoxSingle, BoxSingle, BoxNone):  "┦",
	createDir(BoxHeavy, BoxSingle, BoxHeavy, BoxNone):   "┩",
	createDir(BoxSingle, BoxHeavy, BoxSingle, BoxNone):  "┧",
	// T-Right
	createDir(BoxSingle, BoxSingle, BoxNone, BoxSingle): "├",
	createDir(BoxDouble, BoxDouble, BoxNone, BoxDouble): "╠",
	createDir(BoxDouble, BoxDouble, BoxNone, BoxSingle): "╟",
	createDir(BoxSingle, BoxSingle, BoxNone, BoxDouble): "╞",
	createDir(BoxHeavy, BoxHeavy, BoxNone, BoxHeavy):    "┣",
	createDir(BoxHeavy, BoxHeavy, BoxNone, BoxSingle):   "┠",
	createDir(BoxSingle, BoxSingle, BoxNone, BoxHeavy):  "┝",
	createDir(BoxSingle, BoxHeavy, BoxNone, BoxHeavy):   "┢",
	createDir(BoxHeavy, BoxSingle, BoxNone, BoxSingle):  "┞",
	createDir(BoxHeavy, BoxSingle, BoxNone, BoxHeavy):   "┡",
	createDir(BoxSingle, BoxHeavy, BoxNone, BoxSingle):  "┟",
	// Cross
	createDir(BoxSingle, BoxSingle, BoxSingle, BoxSingle): "┼",
	createDir(BoxDouble, BoxDouble, BoxDouble, BoxDouble): "╬",
	createDir(BoxSingle, BoxSingle, BoxDouble, BoxDouble): "╪",
	createDir(BoxDouble, BoxDouble, BoxSingle, BoxSingle): "╫",
	createDir(BoxHeavy, BoxHeavy, BoxHeavy, BoxHeavy):     "╋",
	createDir(BoxSingle, BoxSingle, BoxHeavy, BoxHeavy):   "┿",
	createDir(BoxHeavy, BoxHeavy, BoxSingle, BoxSingle):   "╂",
	createDir(BoxSingle, BoxHeavy, BoxHeavy, BoxHeavy):    "╈",
	createDir(BoxHeavy, BoxSingle, BoxSingle, BoxSingle):  "╀",
	createDir(BoxHeavy, BoxSingle, BoxHeavy, BoxHeavy):    "╇",
	createDir(BoxSingle, BoxHeavy, BoxSingle, BoxSingle):  "╁",
	createDir(BoxHeavy, BoxHeavy, BoxSingle, BoxHeavy):    "╊",
	createDir(BoxSingle, BoxSingle, BoxHeavy, BoxSingle):  "┽",
	createDir(BoxHeavy, BoxHeavy, BoxHeavy, BoxSingle):    "╉",
	createDir(BoxSingle, BoxSingle, BoxSingle, BoxHeavy):  "┾",
	createDir(BoxSingle, BoxHeavy, BoxSingle, BoxHeavy):   "╆",
	createDir(BoxHeavy, BoxSingle, BoxHeavy, BoxSingle):   "╃",
	createDir(BoxSingle, BoxHeavy, BoxHeavy, BoxSingle):   "╅",
	createDir(BoxHeavy, BoxSingle, BoxSingle, BoxHeavy):   "╄",
}
