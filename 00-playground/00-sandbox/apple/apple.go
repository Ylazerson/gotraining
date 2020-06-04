// B''H

package apple

import "fmt"

// Create TapeRecorder struct and methods
type TapeRecorder struct {
	Microphones int64
	Song        string
}

func Play(song string) {
	fmt.Print("Now playing:", song)
	fmt.Println("On Apple Tape Recorder")
	fmt.Println("... song complete.")
}

func Record(song string) {
	fmt.Print("Now recording:", song)
}
