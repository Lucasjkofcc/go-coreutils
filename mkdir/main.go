package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso:", os.Args[0], "Diretório")
		return
	}
	for _, diretorio := range os.Args[1:] {
		os.Mkdir(diretorio, 0755)
	}
}
