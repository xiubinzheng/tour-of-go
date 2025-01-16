package main

import "fmt"

func main(){
	defer fmt.Println("you")

	fmt.Println("FUCK")
}