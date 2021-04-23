package main

import (
	"fmt"
)

func mapper(words []string, start int, end int, word_freq_map *SafeFrequencyMap, finish_chan chan int) {
	for i := start; i < end; i++ {
		word_freq_map.IncrementFrequency(words[i])
	}
	finish_chan <- 0
}

func reducer(word_freq_map *SafeFrequencyMap) {
	writeMapToFile(word_freq_map.mp, OUTPUT_FILE_NAME)
}

func main() {
	// Read the file then File -> String array
	words, err := readFileToStringArray(INPUT_FILE_NAME)
	panicIfError(err)

	// Init the map
	word_freq_map := SafeFrequencyMap{mp: make(map[string]int)}

	// Create the waiting channel
	w8ng_chan := make(chan int, 5)

	// Create the go routines with the appropriate parameters
	portion := len(words) / ROUTINES_COUNT
	start, end := 0, 0

	for i := 1; i < ROUTINES_COUNT; i++ {
		fmt.Printf("Routine %v : %v -> %v\n", i, start, end)
		start, end = end, end+portion
		go mapper(words, start, end, &word_freq_map, w8ng_chan)
	}

	// Mapper for the rest
	fmt.Printf("Routine %v : %v -> %v\n", ROUTINES_COUNT, end, len(words))
	go mapper(words, end, len(words), &word_freq_map, w8ng_chan)

	// Wait for the routines to finish
	for i := 0; i < ROUTINES_COUNT; i++ {
		<-w8ng_chan
	}

	// Call the reducer
	reducer(&word_freq_map)
}
