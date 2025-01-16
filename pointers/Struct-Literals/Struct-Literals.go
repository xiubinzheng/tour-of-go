package main

import "fmt"

type Vertex struct {
	X, Y int
}

func main() {
	v1 := Vertex{1, 2}
	v2 := Vertex{X: 1, Y: 69}
	v3 := Vertex{}
	p := &Vertex{1, 2}

	fmt.Println(v1, v2, v3, p)

	var (
		x1 = Vertex{1, 2}
		x2 = Vertex{X: 1, Y: 69}
		x3 = Vertex{}
		p2 = &Vertex{1, 2}
	)

	fmt.Println(x1, x2, x3, p2)

}
