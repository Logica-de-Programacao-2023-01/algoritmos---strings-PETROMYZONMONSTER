package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Digite a string.")
	scanner.Scan()
	str := scanner.Text()
	for i := len(str) - 1; i >= 0; i-- {
		rts := ""
		rts += string(str[i])
		fmt.Println(rts)
	}
}
