package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	target := rand.Intn(100) + 1
	guess(target, 10)
}

func guess(target int, numberOfTries int) {
	var guessedTarget int
	for i := 0; i < numberOfTries; i++ {
		fmt.Println("You have", numberOfTries-i, "Tries left")
		guessedTarget = guessTheTarget()
		if target == guessedTarget {
			fmt.Println("Good job! You guessed it!")
			return
		} else if target < guessedTarget {
			fmt.Println("Oops. Your guess was HIGH.")
		} else if target > guessedTarget {
			fmt.Println("Oops. Your guess was LOW.")
		}
	}
	fmt.Println("It was", target)
}

func guessTheTarget() int {
	fmt.Print("Guess the number: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	return convertStringToInt(input)
}

func convertStringToInt(input string) int {
	input = strings.TrimSpace(input)
	num, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)
	}
	return num
}
