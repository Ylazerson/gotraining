// B''H

// Sample program to show how the for range has both value and pointer semantics.
package main

import "fmt"

func main() {

	// -- -------------------------------------
	// Using the pointer semantic form of the for range.
	friends := [5]string{"Annie", "Betty", "Charley", "Doug", "Edward"}

	fmt.Println("--------------------------------------")
	fmt.Println("pointer semantic form of the for range")
	fmt.Println("--                                  --")
	fmt.Printf("friends[1] before the loop: [%s]\n", friends[1])

	for i := range friends {
		friends[1] = "Jack"

		if i == 1 {
			fmt.Printf("friends[1] within the loop: [%s]\n", friends[1])
		}
	}

	// -- -------------------------------------
	// Using the value semantic form of the for range.
	//     - `v` is a local copy and its looking at local copy of the array

	friends = [5]string{"Annie", "Betty", "Charley", "Doug", "Edward"}

	fmt.Println("--------------------------------------")
	fmt.Println("value semantic form of the for range")
	fmt.Println("--                                  --")
	fmt.Printf("friends[1] before the loop: [%s]\n", friends[1])

	for i, v := range friends {
		friends[1] = "Jack"

		if i == 1 {
			fmt.Printf("friends[1] within the loop: [%s]\n", friends[1])
			fmt.Printf("v          within the loop: [%s]\n", v)
		}
	}

	// -- -------------------------------------
	// Using the value semantic form of the for range but with pointer
	// semantic access. DON'T DO THIS.
	friends = [5]string{"Annie", "Betty", "Charley", "Doug", "Edward"}

	fmt.Println("--------------------------------------")
	fmt.Println("Mixed semantics - UGGHHHHH")
	fmt.Println("--                                  --")
	fmt.Printf("friends[1] before the loop: [%s]\n", friends[1])

	for i, v := range &friends { // notice the &
		friends[1] = "Jack"

		if i == 1 {
			fmt.Printf("friends[1] within the loop: [%s]\n", friends[1])
			fmt.Printf("v          within the loop: [%s]\n", v)
		}
	}
}
