package main

import (
	"fmt"
	"os"
)

func main() {
	arquivo, err := os.ReadDir(".")
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}
	for _, arquivo := range arquivo {
		fmt.Print(arquivo.Name(), "\n")
	}
}
