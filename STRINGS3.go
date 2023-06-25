package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "string"
	fmt.Println(input)
	fmt.Println(strings.ReplaceAll(input, "i", "*"))
}
