package main

import "fmt"

func main()  {
	var (
		x int
		l bool
		p float64
	)
	y := 1;
	z,c,v := 1,2,3;

	fmt.Printf("%v, %v, %v \n", x, l, p)
	fmt.Printf("%v, %T \n", y, y)
	fmt.Printf("%v , %v, %v \n", z,c,v)
	
}