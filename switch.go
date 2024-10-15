package main

import (
	"fmt"
)

func main() {
	
	var i int
	fmt.Print("Write a number from 1-3: ")
	fmt.Scanf("%d", &i)
	fmt.Printf("write %d as ", i)
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

}