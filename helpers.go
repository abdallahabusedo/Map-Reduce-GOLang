package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func panicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func readFileToStringArray(fname string) ([]string, error) {
	file, err := os.Open(fname)
	panicIfError(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	words := make([]string, 0)

	for scanner.Scan() {
		line := strings.ToLower(scanner.Text())
		words = append(words, strings.Split(line, " ")...)
	}

	return words, scanner.Err()
}

type kv struct {
	Key   string
	Value int
}

func writeMapToFile(mp map[string]int, path string) error {

	var ss []kv
	for k, v := range mp {
		ss = append(ss, kv{k, v})
	}
	sort.Slice(ss, func(i, j int) bool {
		if ss[i].Value == ss[j].Value {
			return ss[i].Key < ss[j].Key
		}
		return ss[i].Value > ss[j].Value
	})

	file, err := os.Create(path)
	panicIfError(err)
	defer file.Close()
	writer := bufio.NewWriter(file)
	for _, kv := range ss {
		fmt.Fprintln(writer, kv.Key, ":", kv.Value,"")
	}
	return writer.Flush()
}
