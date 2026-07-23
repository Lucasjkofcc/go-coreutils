package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso:", os.Args[0], "<segundos>")
		return
	}

	tempo, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	time.Sleep(time.Duration(tempo) * time.Second)
}
