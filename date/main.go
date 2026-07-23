package main

import (
	"fmt"
	"time"
)

func main() {
	tempo := time.Now().Format("15:04:05")
	fmt.Println(tempo)
}
