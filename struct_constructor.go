package main

import (
	"fmt"
)

/* Define a strcut with name and power */
type Saiyan struct {
	Name  string
	Power int
}

/* Structures don't have constructors.
 *Instead, you create a function that returns an instance of the desired type */
func NewSaiyanPointer(name string, power int) *Saiyan {
	return &Saiyan{
		Name:  name,
		Power: power,
	}
}

/* Our Constructor does not have to return a pointer
   We can return the object */
func NewSaiyanObject(name string, power int) Saiyan {
	return Saiyan{
		Name:  name,
		Power: power,
	}
}

func main() {
	goku_pointer := NewSaiyanPointer("Goku", 9000)
	fmt.Println(goku_pointer)

	goku_object := NewSaiyanObject("Goku", 9000)
	fmt.Println(goku_object)

}
