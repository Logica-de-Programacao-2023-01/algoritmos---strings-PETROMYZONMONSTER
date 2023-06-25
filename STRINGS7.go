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
	if strings.ContainsAny(str, "1234567890") {
		fmt.Println("Contém número.")
	} else {
		fmt.Println("Não contém número.")
	}
}
