package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Digite a sua string:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	fmt.Println("Sua string em letra maiúscula é:", strings.ToTitle(s))
}
