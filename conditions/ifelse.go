package main

import "fmt"

func main(){

	i := 5

	if i % 2 == 0{
		fmt.Println(i, "is even")
	}else{
		fmt.Println(i, "is odd")
	}

	k := 100

	if k == 100{
		fmt.Println(k, " is 100")
	}

	j := 50

	if j < 50{
		fmt.Println(j, " is less than 50")
	}else if j > 50{
		fmt.Println(j, " is greater than 50")
	}else{
		fmt.Println(j, " is equal 50")
	}
}