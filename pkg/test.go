package pkg

import (
	"fmt"
	"path/filepath"
	"runtime"
	"testing"
)

func AssertInt(t *testing.T, res, exp int) {

	if res != exp {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf(
			"\033[31m%s:%d: expected '%d', got: '%d' \033[39m\n\n",
			filepath.Base(file),
			line,
			exp,
			res,
		)
		t.FailNow()
	}
}
