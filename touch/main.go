package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso:", os.Args[0], "<arquivo> <arquivo>")
		return
	}
	for _, arquivo := range os.Args[1:] {
		os.Create(arquivo)
	}
}
