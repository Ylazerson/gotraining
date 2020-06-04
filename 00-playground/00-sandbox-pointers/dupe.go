// B''H

/*
go mod init sandbox/dupe
go run dupe.go
*/

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var nondupes = flag.Bool("nondupes", false, "display non-duplicates")

func main() {

	flag.Parse()

	counts := make(map[string]int)

	for _, filename := range flag.Args() {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}

	for line, n := range counts {

		if !*nondupes && n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}

		if *nondupes && n == 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

}
