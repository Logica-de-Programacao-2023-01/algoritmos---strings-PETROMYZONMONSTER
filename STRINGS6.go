package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Digite a string.")
	scanner.Scan()
	str := scanner.Text()
	fmt.Println(strings.Count(str, " ") + 1)
}
