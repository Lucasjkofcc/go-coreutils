package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	formato := os.Args[1]
	args := os.Args[2:]

	resultado := ""
	argIndex := 0

	for i := 0; i < len(formato); i++ {
		if formato[i] == '%' && i+1 < len(formato) {
			switch formato[i+1] {
			case 's':
				if argIndex < len(args) {
					resultado += args[argIndex]
					argIndex++
				}
				i++

			case 'd':
				if argIndex < len(args) {
					numero, err := strconv.Atoi(args[argIndex])
					if err == nil {
						resultado += strconv.Itoa(numero)
					}
					argIndex++
				}
				i++

			case '%':
				resultado += "%"
				i++

			default:
				resultado += "%"
			}
		} else if formato[i] == '\\' && i+1 < len(formato) {
			switch formato[i+1] {
			case 'n':
				resultado += "\n"
				i++

			case 't':
				resultado += "\t"
				i++

			default:
				resultado += string(formato[i])
			}
		} else {
			resultado += string(formato[i])
		}
	}

	fmt.Print(resultado)

	_ = strings.Builder{}
}
