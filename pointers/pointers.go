package main

import "fmt"

func addOne(num *int) {
	*num = *num + 1
}

// Uses of pointers
// - Can be useful for mutating values without having to copy the value around
// - Can be useful for referencing large structs in functions to keep memory usage to a minimum instead of making copies and reassigning them to the original variable


type position struct {
	x float32
	y float32
}

type badGuy struct {
	name 		string
	health 	int
	pos 		position
}

func whereIsBadGuy(b *badGuy) {
	x := b.pos.x
	y := b.pos.y

	fmt.Println("(", x,",",y,")")
}


func main() {

	x := 5

	fmt.Println(x)
	
	// * means pointer to value that follows
	// & is the symbol that returns virtual memory address of thing that follows (var or func etc.)
	// xPtr := &x // style 1
	var xPtr *int = &x // style 2
	fmt.Println(xPtr)

	addOne(xPtr)
	fmt.Println(x)

	p := position{4, 2} // x = 4, y = 2
	b := badGuy{"Jabba The Hut", 100, p}
	whereIsBadGuy(&b)
}