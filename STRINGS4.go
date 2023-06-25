package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Digite a sua string.")
	scanner.Scan()
	string1 := scanner.Text()
	fmt.Println("Digite a segunda string.")
	scanner.Scan()
	string2 := scanner.Text()
	if string1 == string2 {
		fmt.Println("As strings são iguais.")
	} else {
		fmt.Println("As strings não são a mesma.")
	}
}
