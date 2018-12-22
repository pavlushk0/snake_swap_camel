package main

import (
	io "bufio"
	"fmt"
	"os"
	"strings"
)

func readFile(fname string) (out []string) {
	file, err := os.Open(fname)
	if err != nil {
		fmt.Printf("readFile(): file %s not exist\n", fname)
		os.Exit(0)
	}
	defer file.Close()

	scanner := io.NewScanner(file)
	for scanner.Scan() {
		out = append(out, scanner.Text())
	}

	return out
}

func makeFuncProto() {

}

func snakeToCamel(fname string) {
	fstr := readFile(fname)

	for i := 0; i < len(fstr); i++ {
		if fstr[i] == "/*" {
			i++
			for fstr[i] != "*/" {
				i++
			}
		}

		fid := strings.LastIndex(fstr[i], "func")
		if fid != -1 {
			bid := strings.Index(fstr[i], "(")

			funcWrds := strings.Split(fstr[i][fid+5:bid], "_")

			fmt.Println(funcWrds)
		}
	}
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
