package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso:", os.Args[0], "Arquivo")
		return
	}

	arquivo, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}
	var linhas []string
	scanner := bufio.NewScanner(arquivo)
	for scanner.Scan() {
		linhas = append(linhas, scanner.Text())
	}
	for _, linha := range linhas[len(linhas)-10:] {
		fmt.Println(linha)
	}
}
