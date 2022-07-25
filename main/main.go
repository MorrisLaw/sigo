package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("sigo>>")
		token := scanner.Scan()
		if !token {
			return
		}
		code := scanner.Text()
		fmt.Println(code)
	}
}
