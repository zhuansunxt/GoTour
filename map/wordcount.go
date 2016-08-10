// Counting word appearance number in a given string
// The purpose of this practice is to get familiar with golang
// map data structure's usage

package wordcount

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	tokens := strings.Fields(s)
	wc := make(map[string]int)

	for _, w := range tokens {
		_, ok := wc[w]
		if !ok {
			wc[w] = 0
		} 
		wc[w] += 1
	}

	return wc
}

func main() {
	wc.Test(WordCount)
}