// B''H //

// This will be the quickest because no need to
// make the `append` function call.
package main

import "fmt"

func main() {

	// Need to set length now so all indexes are
	// available for use up front.
	data := make([]string, 1e5)

	// Capture the capacity of the slice.
	lastCap := cap(data)

	// Append ~100k strings to the slice.
	for record := 0; record < 1e5; record++ {

		// Use the built-in function append to add to the slice.
		value := fmt.Sprintf("Rec: %d", record)

		// No `append` function call
		data[record] = value

		// When the capacity of the slice changes, display the changes.
		if lastCap != cap(data) {

			// Calculate the percent of change.
			capChg := float64(cap(data)-lastCap) / float64(lastCap) * 100

			// Save the new values for capacity.
			lastCap = cap(data)

			// Display the results.
			fmt.Printf("Addr[%p]\tIndex[%d]\t\tCap[%d - %2.f%%]\n",
				&data[0],
				record,
				cap(data),
				capChg)
		}
	}
}
