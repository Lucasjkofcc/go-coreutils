package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso:", os.Args[0], "<caminho>")
		return
	}
	fmt.Println(filepath.Dir(os.Args[1]))
}
