package main

import "fmt"
type ByteSize float64

const (
	_ = iota
	KB ByteSize = 1 << (iota * 10)
	MB
	GB
	TB
	PB
	EB
)

func main()  {
	fmt.Printf("%f , %b \n", KB, KB)	
	fmt.Printf("%f , %b \n", MB, MB)	
	fmt.Printf("%f , %b \n", GB, GB)	
	fmt.Printf("%f , %b \n", TB, TB)	
	fmt.Printf("%f , %b \n", EB, EB)	
}