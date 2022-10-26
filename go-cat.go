package main

import (
	"fmt"
	"os"
)

func main() {
	
	fileNameArg := os.Args[1]

	dat, err := os.ReadFile(fileNameArg)
	check(err)
	fmt.Print(string(dat))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
