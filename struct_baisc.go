package main

import (
	"fmt"
)

/* Define a strcut with name and power */
type Saiyan struct {
	Name  string
	Power int
}

/*“In the code below, we say that the type *Saiyan is the receiver of the Super method. */
func (s *Saiyan) Super() {
	s.Power += 10000
}

func main() {
	/* Create a new struct and initilze it */
	goku := Saiyan{"Goku", 9000}
	fmt.Println(goku)

	/* go agruments are passed in by copy, so the value will not change*/
	/* So the power is sitll 9000 insteand of 19000 */
	Super(goku)
	fmt.Println(goku.Power) ///9000

	/* Here we store the address of the object Saiyan in the variable goku_pointer */
	/* Hence when passed, the address it copy and it still points to the same object */
	/* Now the object passed will get modified */
	goku_pointer := &Saiyan{"Goku", 9000}
	Super_Pointer(goku_pointer)
	fmt.Println(goku_pointer.Power) //19000

	/* Invoke function On Structure */
	/* Since we are using pointer here, the object is mutabled */
	goku.Super()
	/* In above code we say that Saiyan is the receiver of the function super */
	fmt.Println(goku.Power) //19000
}

func Super(s Saiyan) {
	s.Power += 10000
}

func Super_Pointer(s *Saiyan) {
	s.Power += 10000
}
