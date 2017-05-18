package ansi

import (
	"strconv"
)

type cursor string

const (
	UP         cursor = "A"
	DOWN       cursor = "B"
	RIGHT      cursor = "C"
	LEFT       cursor = "D"
	NEXT_LINE  cursor = "E"
	PREV_LINE  cursor = "F"
	SELECT     cursor = "H"
	CLEAR      cursor = "J"
	LINE_CLEAR cursor = "K"
)

func (c cursor) itoa(i int) string {
	switch i {
	case 0:
		return ""
	default:
		return strconv.Itoa(i)
	}
}

func (c cursor) ctoa(i ...int) string {
	switch len(i) {
	case 1:
		return "\x1b[" + c.itoa(i[0]) + string(c)
	case 2:
		return "\x1b[" + c.itoa(i[0]) + ";" + c.itoa(i[1]) + string(c)
	default:
		return ""
	}
}

func Cursor(i int, c cursor) string {
	return c.ctoa(i)
}

func Clear() string {
	return CLEAR.ctoa(2) + SELECT.ctoa(1, 1)
}
