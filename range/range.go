package main

import "fmt"

func main(){

	str_arr := []string{"a","b","c","d","e"}
	for i,c := range str_arr{
		fmt.Println("str[",i,"] = ",c)
	}

	m := map[string]string{
		"name":"jitendar",
		"gender":"male",
	}

	for k,v := range m{
		fmt.Println(k,":",v)
	}
}