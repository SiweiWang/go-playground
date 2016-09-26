package main

import (
	"fmt"
	"os"
)

/* Basic language syntax */
func main() {
	/* Print statement */
	fmt.Println("hello world")

	/* String, which can be added together with +.*/
	fmt.Println("1+1=", 1+1)
	fmt.Println("7.0/3.0=", 7.0/3.0)

	/* Integers, floats  and boolean operators */
	fmt.Println(true && false)
	fmt.Println(true || false)

	/* Varaible declearation and definition */
	var power int
	power = 9000
	fmt.Printf("the power is %d\n", power)

	/* Declare and define at the same time */
	number := 42
	fmt.Printf("the number is %d\n", number)

	/* Note that a varaible can not get declared twice */
	// number :=43 // this will give complie error

	/* We can also declear and define mulitple variables at the same time */
	first_name, last_name, hobby := "Siwei", "Wang", "coding"
	fmt.Printf("My name is %s %s and I love %s!\n", first_name, last_name, hobby)

	/* Arguments  args[0] is the name of exec, passed in args start with 1*/
	fmt.Println(os.Args[0])
	if len(os.Args) == 2 {
		fmt.Println(os.Args[1])
	}
	fmt.Printf("1 + 2 = %d\n", add(1, 2))
	x, y := swap(1, 2)
	fmt.Printf("the swap result of 1 and 2 is %d  and %d\n", x, y)
}

func add(a int, b int) int {
	return a + b
}

func swap(a int, b int) (int, int) {
	return b, a
}
