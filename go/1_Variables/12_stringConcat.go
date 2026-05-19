package main

import "fmt"


func main() {
	var username string = "presidentSkroob"
//	var password1 int = 12345
	var password string = "12345" 

	fmt.Println("Authorization: Basic", username+":"+password)
}
