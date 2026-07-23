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
	conteudo, err := os.ReadFile(os.Args[1])

	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	fmt.Println(string(conteudo))
}
