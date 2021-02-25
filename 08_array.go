// an array is a numbered sequence of elements of a specific length.

package main

import "fmt"

func main() {
	// create an array a that will hold exactly 5 ints.
	// The type of elements and length are both part of the array’s type.
	var a [5]int
	fmt.Println("emp:",a)

	// can set a value at an index using the array[index] = value syntax, 
	// and get a value with array[index].
	a[4] = 100
	fmt.Println("set:" ,a)
	
	// The builtin len returns the length of an array.
	fmt.Println("len:", len(a))

	// Use this syntax to declare and initialize an array in one line.
	b := [5]int{1,2,3,4,5}
	fmt.Println("dcl:",b)

	// compose types to build multi-dimensional data structures.
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3 ; j++ {
			twoD[i][j] = i+j
		}
	}
	fmt.Println("2d: ", twoD)
}