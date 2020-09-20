package main

import "fmt"

func seq(num int) func() bool{
	return func() bool{
		if num % 2 == 0{
			return true
		}else{
			return false
		}
	}
}

func main(){
	even_odd := seq(21)
	fmt.Println(even_odd())
	even_odd2 := seq(2)
	fmt.Println(even_odd2())

}