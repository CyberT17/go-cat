package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {

	if len(os.Args) < 2 {
		return
	}

	fileNameArg := os.Args[1]
	file, err := os.Open(fileNameArg)
	check(err)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
