package ansi

import (
	"strconv"
)

type color uint

const (
	BLACK         color = 0
	RED           color = 1
	GREEN         color = 2
	YELLOW        color = 3
	BLUE          color = 4
	MAGENTA       color = 5
	CYAN          color = 6
	WHITE         color = 7
	LIGHT_GRAY    color = 60
	LIGHT_RED     color = 61
	LIGHT_GREEN   color = 62
	LIGHT_YELLOW  color = 63
	LIGHT_BLUE    color = 64
	LIGHT_MAGENTA color = 65
	LIGHT_CYAN    color = 66
	LIGHT_WHITE   color = 67
)

func (c color) ctoa(i int) string {
	return "\x1b[" + strconv.Itoa(i) + "m"
}

func (c color) clear() string {
	return c.ctoa(0)
}

func (c color) fg() string {
	return c.ctoa(int(c) + 30)
}

func (c color) bg() string {
	return c.ctoa(int(c) + 40)
}

func Color(s string, colors ...color) string {
	for i, c := range colors {
		if i == 0 {
			s = c.fg() + s + c.clear()
		} else {
			s = c.bg() + s + c.clear()
			break
		}
	}
	return s
}
