package main

import "fmt"


func split(number int) (x, y int) {
	x = number + 13
	y = number + 15
	return
}
func main() {
	fmt.Println(split(1))
}