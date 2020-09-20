package main

import "fmt"
func main(){
	s := make([]int, 3)

	s[0] = 1
	s[1] = 2
	s[2] = 3
	fmt.Println(s)
	fmt.Println(len(s))

	// append function is unique in slice

	s = append(s,4,5,6,7,8,10)
	fmt.Println(s)

	// slice syntax

	fmt.Println(s[1:3])
	
	// x := s
	// fmt.Println(x)
	// fmt.Println(s)
	// x[0] = 500
	// fmt.Println(x)
	// fmt.Println(s)
// use copy to prevent from changing both slices

	x := make([]int, len(s))
	copy(x, s)
	x[0] = 500
	fmt.Println(x)
	fmt.Println(s)	

}