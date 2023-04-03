//Escreva um programa que receba uma string e remova todos os espaços em branco. Informe ao usuário o resultado.

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
	s := scanner.Text
	fmt.Print("Sua string sem espaços é:", strings.ReplaceAll(s(), " ", ""))
}
