package main

import "fmt"

func main()  {
	fmt.Printf("I'm %0.2f years old\n", 1.0/3.0)

	resultString := fmt.Sprintf("I'm a print statment from line %v", 8)
	fmt.Println(resultString)


	fmt.Printf("The %s cost %d cents each \n", "gumball", 23)
	fmt.Printf("That will be $%.2f please \n", .23 * 5)


	fmt.Printf("A float: %f \n", 3.1415)
	fmt.Printf("An integer: %d \n", 15)
	fmt.Printf("A string: %s \n", "hello")
	fmt.Printf("A boolean: %t \n", true)
	fmt.Printf("Values: %v %v %v %v %v \n", 1.2, 15, "string", "\n" ,true)
	fmt.Printf("Values: %#v %#v %#v %#v %#v \n", 1.2, 15, "string", "\n" ,true)
	fmt.Printf("Values: %T %T %T %T %T \n", 1.2, 15, "string", "\n" ,true)

	fmt.Println("==================")

	fmt.Printf("%12s | %s\n", "Product", "Cost in Cents")
	fmt.Printf("%12s | %d\n", "Stamps", 50)
	fmt.Printf("%12s | %d\n", "Paper Clips", 9)
	fmt.Printf("%12s | %d\n", "Tape", 1000)

}