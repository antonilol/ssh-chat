package color_tags

import (
	"strings"
)

var colors = map[string]string {
	"black": "30",
	"red": "31",
	"green": "32",
	"yellow": "33",
	"blue": "34",
	"magenta": "35",
	"cyan": "36",
	"white": "37",

	"bgblack": "40",
	"bgred": "41",
	"bggreen": "42",
	"bgyellow": "43",
	"bgblue": "44",
	"bgmagenta": "45",
	"bgcyan": "46",
	"bgwhite": "47",

	"b": "30",
	"r": "31",
	"g": "32",
	"y": "33",
	"l": "34",
	"m": "35",
	"c": "36",
	"w": "37",

	"bb": "40",
	"br": "41",
	"bg": "42",
	"by": "43",
	"bl": "44",
	"bm": "45",
	"bc": "46",
	"bw": "47",

	"reset": "0",
	"rs": "0",

	"bold": "1",
	"bo": "1",

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
			res += z;
		}
	}
	if first {
		return "<" + i + ">"
	}
	return res + "m"
}

func ParseColorTags(b string) string {
	p := ""
	//in_tag := false
	t := ""
	for _, r := range b {
		if r == '<' {
			//in_tag = true
			p += t
			t = ""
		} else if r == '>' {
			//in_tag = false
			if t == "" {
				//p += parseTag("rs")
			} else {
				p += parseTag(t)
				t = ""
			}
		} else {
			t += string(r)
		}
        }
	return p + t + parseTag("reset")
}
