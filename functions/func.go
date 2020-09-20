package main

import "fmt"

func add(a int, b int) int{
	return a+b
}

func diff(a, b, c int)int{
	return a-b-c
}
func main(){
	sum := add(5,6)
	fmt.Println(sum)
	dif := diff(4,5,6)
	fmt.Println(dif)
}