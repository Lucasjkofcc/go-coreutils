package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso:", os.Args[0], "Programa")
		return
	}
	programa := os.Args[1]
	path := os.Getenv("PATH")
	dirs := strings.Split(path, ":")
	for _, dir := range dirs {
		caminho := filepath.Join(dir, programa)
		_, err := os.Stat(caminho)
		if err == nil {
			fmt.Println(caminho)
			return
		}
	}
	fmt.Println(programa, "não encontrado")

}
