package main

import (
	"fmt"
)

func main() {
	var a int8 = 127
	fmt.Println(a+1)

	// ----- Loops ----- //
	// Go only has a for loop!
	
	// Infinite loop
	/* 
	for {
		//	This will just keep looping over this code
	}
	*/ 
	
	// Verbose for loop
	// i := 0
	// for i < 10 {
	// 	fmt.Println("i:", i)
	// 	i = i + 1
	// }
	
	// Shorter variant
	// for i := 0; i < 10; i = i + 1 {
		// 	fmt.Println("i:", i)
		// }
		
		
	// ----- Conditionals ----- //
	test := true
	if test {
		fmt.Println("hi")
	} else if a > 127 {
		fmt.Println("bye")
	} else {
		fmt.Println("else")
	}
}