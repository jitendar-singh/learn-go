package main

import "fmt"

func main(){
	fmt.Println("Hello everyone")

	for i :=0; i < 100; i++ {
		if i % 2 == 0{
			fmt.Printf("%d ",i)
		}
	}
	fmt.Println(hello())
}

func hello() string{
	return "hello world"
}