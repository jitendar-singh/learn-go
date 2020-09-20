package main

import "fmt"

func main(){
	val := 20
	fmt.Println(val)

	ptr := &val
	// print out the address where the value 20 is stored
	fmt.Println(&val)
	fmt.Println(ptr)
	//print the value stored at the address(derefrence)
	fmt.Println(*ptr)
	*ptr = 10

	fmt.Println(*ptr)
	fmt.Println(val)

	val = 50
	fmt.Println(*ptr)
	fmt.Println(val)
}