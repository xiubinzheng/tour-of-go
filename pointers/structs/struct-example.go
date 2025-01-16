package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {

	v := Vertex{69, 420}
	v.X = 66666
	fmt.Println(v.X)
}
