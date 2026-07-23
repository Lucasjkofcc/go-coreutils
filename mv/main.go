package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Uso:", os.Args[0], "Arquivo", "Destino")
		return
	}
	arquivo := os.Args[1]
	destino := os.Args[2]
	info, err := os.Stat(os.Args[2])

	if err == nil && info.IsDir() {
		destino = filepath.Join(os.Args[2], filepath.Base(os.Args[1]))
	}
	os.Rename(arquivo, destino)
}
