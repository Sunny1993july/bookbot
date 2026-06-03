package main

import "fmt"

func main (){

	const name string = "Sunny"
	const somenumber float32  = 31
	
	msg  := fmt.Sprintf("Hi %s, is killed %f times.", name, somenumber)
	fmt.Printf(msg)

}
