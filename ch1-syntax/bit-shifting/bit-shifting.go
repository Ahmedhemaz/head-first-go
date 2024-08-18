package main

import "fmt"


const (
	_ = iota
	c1
	c2
	c3
	c4
	c5
	c6
)

func main()  {
	fmt.Println(c1, c2, c3, c4, c5, c6)
	fmt.Printf("%d \t %b \n", 1<<c1, 1<<c1)
	fmt.Printf("%d \t %b \n", 1<<c2, 1<<c2)
	fmt.Printf("%d \t %b \n", 1<<c3, 1<<c3)
	fmt.Printf("%d \t %b \n", 1<<c4, 1<<c4)
	fmt.Printf("%d \t %b \n", 1<<c5, 1<<c5)
	fmt.Printf("%d \t %b \n", 1<<c6, 1<<c6)
}