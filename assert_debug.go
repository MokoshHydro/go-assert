//go:build debug
package assert
import (
	"fmt"
	"runtime/debug"
	"os"
)

var AssertExitCode = 255

func Assert(v bool, msg string) {
	if(v) {
		fmt.Fprint(os.Stderr, "ASSERT FIRED: ")
		fmt.Fprintln(os.Stderr, msg)
		debug.PrintStack()
		os.Exit(AssertExitCode)
	}
}

func Assertf(v bool, msg string, args ...interface{}) {
	if(v) {
		fmt.Fprint(os.Stderr, "ASSERT FIRED: ")
		s := fmt.Sprintf(msg, args...)
		fmt.Fprintln(os.Stderr, s)
		debug.PrintStack()
		os.Exit(AssertExitCode)
	}
}
