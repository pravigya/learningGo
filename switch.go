package main

import (
	"fmt"
	"time"
)

func switchcase(){

	t:= time.Now()
	switch{
	case t.Hour()<12:
		fmt.Println("before Noon")
	default:
		fmt.Println("After noon")
	}
	
	fmt.Println()
}