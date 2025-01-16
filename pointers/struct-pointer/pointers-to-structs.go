package main

import "fmt"

type Vertex struct{
	X int
	Y int
}

func main(){
	v := Vertex{1, 2}
	p := &v
	fmt.Println(v)

	p.X = 2e11
	fmt.Println(v)

	fmt.Println("-------------------")
	fmt.Println(p)
	fmt.Println(*p)
	
}