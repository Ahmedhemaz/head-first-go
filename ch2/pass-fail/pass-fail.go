package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	printStatus(convertStringToFloat64(readGradFromUser()))
}

func readGradFromUser() string {
	fmt.Print("Enter a grade:")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	return input
}

func convertStringToFloat64(input string) float64 {
	input = strings.TrimSpace(input)
	num, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal(err)
	}
	return num
}

func printStatus(grade float64) {
	var staus string
	if grade >= 60 {
		staus = "Passing!"
	} else {
		staus = "Failing!"
	}
	fmt.Println("A grade of", grade, "is", staus)
}
