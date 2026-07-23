package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso:", os.Args[0], "Pasta")
		return
	}
	pastas := os.Args[1:]
	for _, pasta := range pastas {
		_, err := os.Stat(pasta)
		if err != nil {
			fmt.Println("Erro", pasta, "Não encontrado")
			continue
		}
		erro := os.Remove(pasta)
		if erro != nil {
			fmt.Println("Erro Removendo", pasta, ":", erro)
			return
		}
	}
}
