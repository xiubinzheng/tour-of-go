package main

import "fmt"

func main(){
	names := [4]string{"jason", "john", "rachel", "angela"}
	fmt.Println(names)

	a := names[0:2]
	b := names[2:4]

	fmt.Println(a,b)
	b[0] = "devon"
	b[1] = "share"
	
	fmt.Println(b)
	
}