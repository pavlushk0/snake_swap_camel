package main

import (
	"fmt"
	"os"
)

func makeFuncProto() {

}

func snakeToCamel(fname string) {
	file, err := os.Open(fname)
	if err != nil {
		fmt.Printf("snakeToCamel(): file %s no exist\n", fname)
		os.Exit(0)
	}
	defer file.Close()

}

func camelToSnake() {

}

func main() {
	if (len(os.Args) < 2) || (len(os.Args) > 3) {
		fmt.Println("main(): wrong args count")
		os.Exit(0)
	}

	switch os.Args[1] {
	case "-stc":
		snakeToCamel(os.Args[2])
		os.Exit(0)

	case "-cts":
		camelToSnake()
		os.Exit(0)

	case "-fp":
		makeFuncProto()
		os.Exit(0)
	}

	fmt.Println("unknow option", os.Args[1])
}
