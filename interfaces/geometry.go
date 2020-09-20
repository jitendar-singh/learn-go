package main

import (
	"fmt"
	"math"
)
type geometry interface{
	area()float64
	peri()float64
}

type rect struct{
	width, height float64
}
type circle struct{
	radius float64
}

func (r rect) area()float64{
	return r.width * r.height
}
func (c circle) area()float64{
	return math.Pi * c.radius * c.radius
}
func (r rect) peri()float64{
	return 2 * r.width + 2*r.height
}
func (c circle) peri()float64{
	return 2*math.Pi * c.radius
}

func measure(g geometry){
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.peri())
	
}
func main(){
	r := rect{2.5,3.4}
	fmt.Println("generating rectangle")
	measure(r)
	c := circle{5}
	fmt.Println("generating circle")
	measure(c)
}