//go:build !debug
package assert

func Check(v bool, msg string) {
}

func Checkf(v bool, msg string, args ...interface{}) {
}
