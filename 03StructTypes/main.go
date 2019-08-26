// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare a struct type to maintain information about a user (name, email and age).
// Create a value of this type, initialize with values and display each field.
//
// Declare and initialize an anonymous struct type with the same three fields. Display the value.
package main

import "fmt"

// Add imports.

// Add user type and provide comment.
type user struct {
	name  string
	email string
	age   int8
}

func main() {

	// Declare variable of type user and init using a struct literal.
	userL := user{
		name:  "Literal Variable",
		email: "literal@foo.com",
		age:   31,
	}

	// Display the field values.
	fmt.Println("Name:", userL.name)
	fmt.Println("Email:", userL.email)
	fmt.Println("Age:", userL.age)

	// Declare a variable using an anonymous struct.

	userA := struct {
		name  string
		email string
		age   int8
	}{
		name:  "Anonymous Variable",
		email: "anonymous@foo.com",
		age:   31,
	}

	// Display the field values.
	fmt.Println("Name:", userA.name)
	fmt.Println("Email:", userA.email)
	fmt.Println("Age:", userA.age)
}
