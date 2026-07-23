package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	linhas := 0
	if len(os.Args) < 2 {
		fmt.Println("Uso:", os.Args[0], "Arquivo")
		return
	}

	arquivo, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}
	defer arquivo.Close()

	scanner := bufio.NewScanner(arquivo)

	for scanner.Scan() {
		linha := scanner.Text()
		fmt.Println(linha)
		linhas++
		if linhas == 10 {
			break
		}
	}
}
