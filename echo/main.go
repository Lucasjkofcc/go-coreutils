package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	texto := strings.Join(os.Args[1:], " ")
	fmt.Println(texto)
}
