package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Uso:", os.Args[0], "arquivo", "diretorio")
		return
	}

	origemPath := os.Args[1]
	destinoPath := os.Args[2]

	info, err := os.Stat(destinoPath)
	if err == nil && info.IsDir() {
		destinoPath = filepath.Join(destinoPath, filepath.Base(origemPath))
	}

	origem, err := os.Open(origemPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer origem.Close()

	destino, err := os.Create(destinoPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer destino.Close()

	io.Copy(destino, origem)
}
