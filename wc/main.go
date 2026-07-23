package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	linhas := 0
	palavras := 0
	bytes := 0
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
		bytes += len(scanner.Bytes())
		bytes++
		linhas++
		palavras += len(strings.Fields(linha))
	}

	fmt.Println("Linhas:", linhas)
	fmt.Println("Palavras:", palavras)
	fmt.Println("Bytes:", bytes)
	fmt.Println("Arquivo:", filepath.Base(os.Args[1]))
}
