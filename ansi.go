package ansi

import (
	"strconv"
)

type ansi uint

const (
	BLACK         ansi = 0
	RED           ansi = 1
	GREEN         ansi = 2
	YELLOW        ansi = 3
	BLUE          ansi = 4
	MAGENTA       ansi = 5
	CYAN          ansi = 6
	WHITE         ansi = 7
	LIGHT_GRAY    ansi = 60
	LIGHT_RED     ansi = 61
	LIGHT_GREEN   ansi = 62
	LIGHT_YELLOW  ansi = 63
	LIGHT_BLUE    ansi = 64
	LIGHT_MAGENTA ansi = 65
	LIGHT_CYAN    ansi = 66
	LIGHT_WHITE   ansi = 67
)

func (c ansi) ctoa(i int) string {
	return "\x1b[" + strconv.Itoa(i) + "m"
}

func (c ansi) clear() string {
	return c.ctoa(0)
}

func (c ansi) fg() string {
	return c.ctoa(int(c) + 30)
}

func (c ansi) bg() string {
	return c.ctoa(int(c) + 40)
}

func Ansi(s string, colors ...ansi) string {
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
