package main

import (
	"fmt"
	"os"
)

func makeFuncProto() {
	fmt.Println("make func proto")
}

func snakeToCamel() {
	fmt.Println("snake_to_camel")
}

func camelToSnake() {
	fmt.Println("camel_to_snake")
}

func main() {
	if (len(os.Args) < 2) || (len(os.Args) > 3) {
		fmt.Println("main(): wrong args count")
		os.Exit(0)
	}

	if os.Args[1] == "-stc" {
		snakeToCamel()
		os.Exit(0)
	}

	if os.Args[1] == "-cts" {
		camelToSnake()
		os.Exit(0)
	}

	if os.Args[1] == "-fp" {
		makeFuncProto()
		os.Exit(0)
	}

	fmt.Println("unknow option", os.Args[1])
}
