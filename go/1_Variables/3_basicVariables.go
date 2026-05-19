package main

import "fmt"

func main(){
	var username string = "eddie_cabot"
        fmt.Println("username:", username)

	var isAdmin bool = true
	fmt.Println("isAdmin:", isAdmin)

        var permissions byte = 0x1F  // this can be both byte and int
        fmt.Println("permissions:", permissions) 

	var costPerSMS float64 = 0.05
	fmt.Println("costPerSMS:", costPerSMS)
	



}
