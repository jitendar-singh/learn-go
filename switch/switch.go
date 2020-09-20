package main

import "fmt"

func main(){

	i := 4

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println(i)
	}

	j := 1

	switch j{
	case 1,2:
		fmt.Println("one or two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println(j)
	}


	switch{
	case j == 5:
		fmt.Println("equal")
	case j > 5:
		fmt.Println("Greater")
	case j < 5:
		fmt.Println("less")
	}
}