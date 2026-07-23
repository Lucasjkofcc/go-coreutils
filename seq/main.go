package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso:", os.Args[0], "Número")
		return
	}

	if len(os.Args) == 2 {
		final, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("Erro:", err)
			return
		}

		for i := 1; i <= final; i++ {
			fmt.Println(i)
		}
	}

	if len(os.Args) == 3 {
		comeco, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("Erro:", err)
			return
		}

		fim, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Erro:", err)
			return
		}

		for i := comeco; i <= fim; i++ {
			fmt.Println(i)
		}
	}

	if len(os.Args) == 4 {
		comeco, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("Erro:", err)
			return
		}

		meio, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Erro:", err)
			return
		}

		fim, err := strconv.Atoi(os.Args[3])
		if err != nil {
			fmt.Println("Erro:", err)
			return
		}

		for i := comeco; i <= fim; i += meio {
			fmt.Println(i)
		}
	}
}
