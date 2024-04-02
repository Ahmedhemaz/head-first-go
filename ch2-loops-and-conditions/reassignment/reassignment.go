package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	// as long as at least one variable name is a short variable declaration is new, it's allowed 
	// The new variables names are treated as a declaration and the exisiting ones as an assignment
	a, err := strconv.ParseFloat("abc", 64)
	fmt.Println(err)
	b, err := strconv.ParseFloat("def", 64)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(a, b)

}