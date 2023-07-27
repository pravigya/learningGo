package main

import "fmt"

func ifelse(){

	if num:=9; num<0 {
		fmt.Println(num," negative")

	} else if num<10 {
		fmt.Println(num, "single digit")
	} else {
		fmt.Println(num, "multi digit")
	}
}

// no ternary operator in go