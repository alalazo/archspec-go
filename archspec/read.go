package main

import (
	"fmt"
	"bufio"
	"log"
	"github.com/markbates/pkger"
)

// Read reads a file
func Read(filename string) {
	file, err := pkger.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string
 
	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}
 
	for _, eachline := range txtlines {
		fmt.Println(eachline)
	}
}

func main() {
	pkger.Include("/json")
	Read("/json/cpu/microarchitectures.json")
}