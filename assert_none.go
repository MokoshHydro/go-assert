//go:build !debug
package assert

func Assert(v bool, msg string) {
}

func Assertf(v bool, msg string, args ...interface{}) {
}
