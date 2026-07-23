package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso:", os.Args[0], "Arquivo")
		return
	}
	arquivos := os.Args[1:]
	for _, arquivo := range arquivos {
		_, err := os.Stat(arquivo)
		if err != nil {
			fmt.Println("Erro", arquivo, "Não encontrado")
			continue
		}
		erro := os.RemoveAll(arquivo)
		if erro != nil {
			fmt.Println("Erro Removendo", arquivo, ":", erro)
			return
		}
	}
}
