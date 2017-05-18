package ansi_test

import (
	"github.com/tk-shirasaka/ansi"
	"testing"
)

func TestColors(t *testing.T) {
	var want string
	var get string

	want = "\x1b[30m \x1b[0m"
	get = ansi.Color(" ", ansi.BLACK)

	if want != get {
		t.Errorf("Want %s, get %s", want, get)
	}

	want = "\x1b[31m \x1b[0m"
	get = ansi.Color(" ", ansi.RED)

	if want != get {
		t.Errorf("Want %s, get %s", want, get)
	}

	want = "\x1b[32m \x1b[0m"
	get = ansi.Color(" ", ansi.GREEN)

	if want != get {
		t.Errorf("Want %s, get %s", want, get)
	}

	want = "\x1b[33m \x1b[0m"
	get = ansi.Color(" ", ansi.YELLOW)

	if want != get {
		t.Errorf("Want %s, get %s", want, get)
	}

	want = "\x1b[34m \x1b[0m"
	get = ansi.Color(" ", ansi.BLUE)

	if want != get {
		t.Errorf("Want %s, get %s", want, get)
	}

	want = "\x1b[35m \x1b[0m"
	get = ansi.Color(" ", ansi.MAGENTA)

	if want != get {
		t.Errorf("Want %s, get %s", want, get)
	}

	want = "\x1b[36m \x1b[0m"
	get = ansi.Color(" ", ansi.CYAN)

	if want != get {
		t.Errorf("Want %s, get %s", want, get)
	}

	want = "\x1b[37m \x1b[0m"
	get = ansi.Color(" ", ansi.WHITE)

	if want != get {
		t.Errorf("Want %s, get %s", want, get)
	}

	want = "\x1b[90m \x1b[0m"
	get = ansi.Color(" ", ansi.LIGHT_GRAY)

	if want != get {
		t.Errorf("Want %s, get %s", want, get)
	}

	want = "\x1b[91m \x1b[0m"
	get = ansi.Color(" ", ansi.LIGHT_RED)

	if want != get {
		t.Errorf("Want %s, get %s", want, get)
	}

	want = "\x1b[92m \x1b[0m"
	get = ansi.Color(" ", ansi.LIGHT_GREEN)

	if want != get {
		t.Errorf("Want %s, get %s", want, get)
	}

	want = "\x1b[93m \x1b[0m"
	get = ansi.Color(" ", ansi.LIGHT_YELLOW)

	if want != get {
		t.Errorf("Want %s, get %s", want, get)
	}

	want = "\x1b[94m \x1b[0m"
	get = ansi.Color(" ", ansi.LIGHT_BLUE)

	if want != get {
		t.Errorf("Want %s, get %s", want, get)
	}

	want = "\x1b[95m \x1b[0m"
	get = ansi.Color(" ", ansi.LIGHT_MAGENTA)

	if want != get {
		t.Errorf("Want %s, get %s", want, get)
	}

	want = "\x1b[96m \x1b[0m"
	get = ansi.Color(" ", ansi.LIGHT_CYAN)

	if want != get {
		t.Errorf("Want %s, get %s", want, get)
	}

	want = "\x1b[97m \x1b[0m"
	get = ansi.Color(" ", ansi.LIGHT_WHITE)

	if want != get {
		t.Errorf("Want %s, get %s", want, get)
	}
}
