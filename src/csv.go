package app

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sort"
)

// writeCSV writes the following and follower list to a CSV file.
func (a *App) writeCSV() {
	// Sort the lists for user-friendliness.
	following := a.sortKeys(a.following)
	followers := a.sortKeys(a.followers)

	// Open the CSV file for writing.
	fh, err := os.Create("follower-lists.csv")
	if err != nil {
		log.Panicf("Can't create CSV output: %s", err)
	}

	// Create a writer.
	w := csv.NewWriter(fh)
	w.Write([]string{"Following", "Followers"})

	// Write the CSV records.
	for i := 0; i < max(len(following), len(followers)); i++ {
		record := []string{"", ""}
		if i < len(following) {
			record[0] = following[i]
		}
		if i < len(followers) {
			record[1] = followers[i]
		}

		w.Write(record)
	}

	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("CSV dumped to: follow-lists.csv")
}

// sortKeys returns the sorted keys from the follower/ing lists.
func (a *App) sortKeys(dict map[string]bool) []string {
	var keys []string
	for key, _ := range dict {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

// max returns the greater of two ints (math.Max does float64)
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
