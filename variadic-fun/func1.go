package main

import "fmt"

func mult(nums...int)int{
	total := 1
	for _,num := range nums{
		total *= num
	}
	return total
}

func main(){
	// func that takes arbitary number of inputs is called a variadic function Println
	fmt.Println(mult(1,2,3,4,5,6,7))

}