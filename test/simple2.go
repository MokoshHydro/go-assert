package main

import (
	assert "github.com/mokoshhydro/go-assert"
	"fmt"
)

func main() {
	assert.Assertf(true, "Format message: %d %s", 1, "test")
	fmt.Println("NOT REACHED")
}

