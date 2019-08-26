// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare three variables that are initialized to their zero value and three
// declared with a literal value. Declare variables of type string, int and
// bool. Display the values of those variables.
//
// Declare a new variable of type float32 and initialize the variable by
// converting the literal value of Pi (3.14).
package main

import (
	"fmt"
	"math"
)

// Add imports

// main is the entry point for the application.
func main() {

	// Declare variables that are set to their zero value.
	var strVar string
	var intVar int
	var boolVar bool

	// Display the value of those variables.
	fmt.Printf("var strVar string \t %T [%v] \n", strVar, strVar)
	fmt.Printf("var intVar int \t \t %T [%v] \n", intVar, intVar)
	fmt.Printf("var boolVar boolean \t %T [%v] \n", boolVar, boolVar)

	fmt.Println("**********************************")
	// Declare variables and initialize.
	// Using the short variable declaration operator.
	strVarI := "Marcela"
	intVarI := 31
	boolVarI := true

	// Display the value of those variables.
	fmt.Printf("var strVarI string \t %T [%v] \n", strVarI, strVarI)
	fmt.Printf("var intVarI int \t %T [%v] \n", intVarI, intVarI)
	fmt.Printf("var boolVarI boolean \t %T [%v] \n", boolVarI, boolVarI)

	fmt.Println("**********************************")

	// Perform a type conversion.
	piValue := math.Pi

	// Display the value of that variable.

	fmt.Printf("var piValue pi \t %T [%v] \n", piValue, piValue)
	fmt.Println("**********************************")

	// Perform a type conversion.
	pi := float32(math.Pi)

	// Display the value of that variable.
	fmt.Printf("var pi pi \t %T [%v] \n", pi, pi)
}
