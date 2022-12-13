package main

import "fmt"

func pointer () {
	var p *int  // declaring a pointer
	i := 50     //initializing a variable
	p = &i      // pass the address to p
	fmt.Println(p) // print the value of a pointer
	fmt.Println(*p)  // print the value at the pinter address
	*p = 60    // changing the value of a address
	fmt.Println(i) // printing the pointer value
}