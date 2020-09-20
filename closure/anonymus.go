package main

import "fmt"
func say_hello(msg string)  {
	fmt.Println(msg)
}

func return_anoyn_func() func(string){
	// return a anoymms func
	return func (msg string){
		fmt.Println(msg)
	}
}
func main()  {
	say_hello("hello")

	//anaymous function declared
	func (msg string)  {
		fmt.Println(msg)
	}("hello anonymus")
	print_func := return_anoyn_func()
	print_func("hello returned from anoynmus")
}