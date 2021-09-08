package color_tags

import (
	"strings"
	"regexp"
)

var colors = map[string]string {
	// foreground colors
	"black": "30",
	"a": "30",

	"red": "31",
	"r": "31",

	"green": "32",
	"g": "32",

	"yellow": "33",
	"y": "33",

	"blue": "34",
	"l": "34",

	"magenta": "35",
	"m": "35",

	"cyan": "36",
	"c": "36",

	"white": "37",
	"w": "37",

	// background colors
	"bgblack": "40",
	"ba": "40",

	"bgred": "41",
	"br": "41",

	"bggreen": "42",
	"bg": "42",

	"bgyellow": "43",
	"by": "43",

	"bgblue": "44",
	"bl": "44",

	"bgmagenta": "45",
	"bm": "45",

	"bgcyan": "46",
	"bc": "46",

	"bgwhite": "47",
	"bw": "47",

	// special
	"reset": "0",
	"rs": "0",
	"/": "0",

	"bold": "1",
	"b": "1",

	"dim": "2",
	"d": "2",

	"italic": "3",
	"i": "3",

	"underline": "4",
	"u": "4",

	"blink": "5",
	"bi": "5",

	"reversed": "7",
	"rv": "7",

	"rgb": "38;2",
	"rgbfg": "38;2",
	"rgbbg": "48;2",
}

func parseTag(i string) string {
	str := string(i[:])
	res := "\033["
	first := true
	splitted := strings.Split(str, ",")
	for _, v := range splitted {
		if z, ex := colors[v]; ex {
			if !first {
				res += ";"
			} else {
				first = false
			}
			res += z
		} else if m, _ := regexp.MatchString("[0-9]+", v); m {
			if !first {
				res += ";"
			} else {
				first = false
			}
			res += v
		} else {
			first = true
			break
		}
	}
	if first {
		return "<" + i + ">"
	}
	return res + "m"
}

func ParseColorTags(msg string) string {
	p := ""
	in_tag := false
	t := ""
	for i, r := range msg {
		if r == '<' && !in_tag {
			in_tag = true
			p += t
			t = ""
		} else if r == '>' {
			if in_tag {
				in_tag = false
				p += parseTag(t)
				t = ""
			} else if i == 0 {
				return msg
			}
		} else {
			t += string(r)
		}
	}
	if in_tag {
		p += "<"
	}
	return p + t + parseTag("reset")
}
