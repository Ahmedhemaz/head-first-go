package main

import (
	"fmt"
	"log"
)


func main()  {
	area , err := paintNeeded(4.2, 3.0)
	if err != nil {
		log.Fatal(err)
	}
	var myFloatPointer = &area
	fmt.Printf("Total needed %.2f \n" , area)
	fmt.Println(myFloatPointer)
	*myFloatPointer = 4.0
	fmt.Println(area)
}


func paintNeeded( width float64,height float64) (float64, error)  {
	if width < 0 {
		err := fmt.Errorf("width can't be less than 0")
		return 0,err
	}
	if height < 0 {
		err := fmt.Errorf("height can't be less than 0")
		return 0,err
	}
	area := width * height
	return area, nil;
	
}
