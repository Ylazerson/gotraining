// B''H

// Sample program to show how one needs to be careful when appending
// to a slice when you have a reference to an element.
package main

import "fmt"

type user struct {
	likes int
}

func main() {

	// Declare a slice of 3 users.
	users := make([]user, 3)

	inspectSlice(users)
	// -- -----------------------------------------

	// -- -----------------------------------------
	// Share the user at index 1.
	shareUser := &users[1]

	shareUser.likes++

	fmt.Println("*************************")
	fmt.Println("shareUser.likes", shareUser.likes, &shareUser.likes)

	inspectSlice(users)
	// -- -----------------------------------------

	// -- -----------------------------------------
	// Add a new user.
	users = append(users, user{})

	inspectSlice(users)
	// -- -----------------------------------------

	// -- -----------------------------------------
	// Add another like for the user that was shared.
	shareUser.likes++

	fmt.Println("*************************")
	fmt.Println("shareUser.likes", shareUser.likes, &shareUser.likes)

	// Notice the last like has not been recorded.
	inspectSlice(users)
	// -- -----------------------------------------

}

// inspectSlice exposes the slice header for review.
func inspectSlice(users []user) {

	fmt.Println("*************************")

	fmt.Printf("Length[%d] Capacity[%d]\n", len(users), cap(users))

	for i := range users {

		fmt.Println(
			i,
			users[i].likes,
			&users[i].likes,
		)
	}

}
