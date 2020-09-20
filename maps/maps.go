package main

import "fmt"

// Maps are similar to Python dictonary
func main(){
	m := make(map[string]int)

	m["a"] = 1
	m["b"] = 2

	fmt.Println(m)

	//value at particular key
	fmt.Println(m["a"])

	//length
	fmt.Println(len(m))

	//delete element
	delete(m,"a")
	fmt.Println(m)

	//define in a concise manner if key & values are known

	m2 := map[string]int{
		"Roll":1,
		"Marks":75,
	}
	fmt.Println(m2)

	//value is present
	// delete(m,"b")
	val, is_val_present := m["b"]
	fmt.Println(val," present is ",is_val_present)

	_, val_present := m["a"]
	fmt.Println("present is ",val_present)
}