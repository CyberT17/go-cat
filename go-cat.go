package main

import (
	"fmt"
	"os"
)

func main() {
	const fileName string = "./test.txt"
	dat, err := os.ReadFile(fileName)
	check(err)
	fmt.Print(string(dat))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
