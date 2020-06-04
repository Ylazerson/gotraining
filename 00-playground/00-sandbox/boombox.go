// B''H

/*
go mod init sandbox/boombox
go run boombox.go
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"sandbox/boombox/apple"
	"sandbox/boombox/asus"
	"sandbox/boombox/player"
)

func main() {
	// Create tapePlayer
	var tPlayer asus.TapePlayer
	tPlayer.Batteries = "Duracell"
	tPlayer.Radio = true
	tPlayer.Song = getSongName()

	// Test it out
	player.Process(
		asus.Play,
		tPlayer.Song,
	)

	// -- ---------------
	fmt.Println("-----")
	// -- ---------------

	// Create recorder
	var tRecorder apple.TapeRecorder
	tRecorder.Microphones = 2
	tRecorder.Song = getSongName()

	// Test it out
	player.Process(
		apple.Play,
		tRecorder.Song,
	)

}

// Get song name
func getSongName() string {

	fmt.Print("Enter song name: ")

	var reader *bufio.Reader = bufio.NewReader(os.Stdin)

	var input string
	var err error

	input, err = reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	return input
}
