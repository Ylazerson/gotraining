// B''H

package asus

import "fmt"

// Create TapePlayer struct and methods
type TapePlayer struct {
	Batteries string
	Radio     bool
	Song      string
}

func Play(song string) {
	fmt.Print("Now playing:", song)
	fmt.Println("On Asus Tape Player")
	fmt.Println("... song complete.")
}
