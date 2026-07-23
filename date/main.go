package main

import (
	"fmt"
	"time"
)

func main() {
	tempo := time.Now().Format("02/01/2006 15:04:05")
	fmt.Println(tempo)
}
