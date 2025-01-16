package main

import "fmt"

func main(){
	fmt.Prinltn("counting")
	
	for i := 0; i < 10; i ++{
		defer fmt.Prinltn(i)
	}

	fmt.Println("done")
}