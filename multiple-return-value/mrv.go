package main

import "fmt"

func add(a int, b int)(int, int, int){
	return a+b, a-b, a*b
}

func main(){
	sum, dif, mul := add(5,6)
	fmt.Println(sum)
	fmt.Println(dif)
	fmt.Println(mul)
}