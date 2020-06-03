// B''H

package player

import "fmt"

type play func(song string)

func Process(f play, song string) {
	fmt.Println("... interface here just doin my job...")
	f(song)
}
