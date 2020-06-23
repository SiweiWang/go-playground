package main

import "fmt"


//zeroval has an int parameter, so arguments will be passed to it by value. 
//zeroval will get a copy of ival distinct from the one in the calling function.

func zeroval(ival int){
	ival = 0
}


// zeroptr in contrast has an *int parameter, 
// meaning that it takes an int pointer. 
// The *iptr code in the function body then dereferences the pointer 
// from its memory address to the current value at that address. 
// Assigning a value to a dereferenced pointer changes the value 
// at the referenced address.
func zeroptr(iptr *int) {
	*iptr = 0
}

// zeroval doesn’t change the i in main, 
// but zeroptr does because it has a reference 
// to the memory address for that variable.
func main() {
	i := 1
	fmt.Println("initial:", 1)
	
	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr", i)

	fmt.Println("pointer:", &i)
}