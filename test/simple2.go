package main

import (
	"github.com/mokoshhydro/go-assert"
	"fmt"
)

func main() {
	assert.Checkf(true, "Format message: %d %s", 1, "test")
	fmt.Println("NOT REACHED")
}

