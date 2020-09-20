package main

import "fmt"

type employee struct{
	first_name string
	last_name string
	id int
}
func main(){
	fmt.Println(employee{first_name:"Jitendar",last_name:"Singh"})
	fmt.Println(employee{"Sunny","",1})
	emp := employee{
		first_name:"Monalisa",
		last_name:"Sarangi",
		id: 3646,
	}
	fmt.Println(emp)

	emp_ptr := &emp
	fmt.Println(emp_ptr.first_name)
	emp_ptr.first_name="Marge"
	fmt.Println(emp_ptr.first_name)

}