package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	texto := "y"
	if len(os.Args) > 1 {
		texto = strings.Join(os.Args[1:], " ")
	}
	for {
		fmt.Println(texto)
	}
}
