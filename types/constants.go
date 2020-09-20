package main

import "fmt"

func main(){
	const pi float64 = 3.1415926
	fmt.Println(pi)

	// cant reassign to pi
	// pi = 4.5
	pi_backup := pi
	fmt.Println(pi_backup)
	pi_backup ++
	fmt.Println(pi_backup)
	


}