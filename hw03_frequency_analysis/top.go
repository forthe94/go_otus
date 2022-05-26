package hw03frequencyanalysis

import (
	"fmt"
	"sort"
	"strings"
)

func Top10(input string) []string {
	// Place your code here.
	if input == "" {
		return nil
	}
	words := strings.Fields(input)

	fmt.Println(words)

	wordsCount := make(map[string]int)

	for _, field := range words {
		wordsCount[field]++
	}

	type kv struct {
		Key   string
		Value int
	}

	var ss []kv

	for k, v := range wordsCount {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})
	start := 0
	end := 0

	for end < len(ss) {
		end++
		if end == len(ss) || ss[start].Value != ss[end].Value {
			sort.Slice(ss[start:end], func(i, j int) bool {
				return ss[i+start].Key < ss[j+start].Key
			})
			start = end
		}

	}

	var ret []string
	count := 0
	for {
		if count == 10 || count == len(ss) {
			break
		}

		ret = append(ret, ss[count].Key)
		count++
	}
	return ret
}
