package main

import (
	"fmt"
	"github.com/alalazo/archspec-go/archspec"
	"log"
)

func main() {
	textLines, err := archspec.ReadJSONAsStringVec()
	if err != nil {
		log.Fatal(err)
	}

	for _, eachline := range textLines {
		fmt.Println(eachline)
	}
}
