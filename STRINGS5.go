package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Digite a sua string.")
	scanner.Scan()
	string1 := scanner.Text()
	if _, err := strconv.ParseFloat(string1, 64); err == nil {
		fmt.Println("A string está em float.")
	} else {
		fmt.Println("A string não está em float")
	}
}
