package ansi

import (
	"regexp"
	"strconv"
	"strings"
)

type AttributeString struct {
	String     string
	Attributes Attributes
}

func ParseString(str string) []AttributeString {
	parts := make([]AttributeString, 0, 1)
	currentAttributes := NoAttributes
	currentPart := strings.Builder{}
	idx := 0
	strR := []rune(str)
	for idx < len(strR) {
		if strR[idx] == escape {
			idx++
			codes := parseAttributes(strR, &idx)
			change := newAttributeChange(codes)
			newAttrs := currentAttributes.Modify(change)
			if newAttrs != currentAttributes {
				if currentPart.Len() > 0 {
					parts = append(parts, AttributeString{currentPart.String(), currentAttributes})
					currentPart.Reset()
				}
				currentAttributes = newAttrs
			}
		} else {
			currentPart.WriteString(string(strR[idx]))
			idx++
		}
	}
	// Write last currentPart
	if currentPart.Len() > 0 {
		parts = append(parts, AttributeString{currentPart.String(), currentAttributes})
	}
	return parts
}

// ----------------------------------------------------------------------------------------------------------------------------
// internal
// ----------------------------------------------------------------------------------------------------------------------------

var regExp = regexp.MustCompile(`(\d+)(?:;(\d+))*m`)

func parseAttributes(str []rune, idx *int) (attributes []int) {
	attributes = make([]int, 0)
	strS := string(str[*idx:])
	matches := regExp.FindAllStringSubmatch(strS, 1)
	if matches != nil {
		*idx = *idx + len(matches[0][0])
		for i := 1; i < len(matches[0]); i++ {
			if val, err := strconv.Atoi(matches[0][i]); err == nil {
				attributes = append(attributes, val)
			}
		}
	}
	return attributes
}
