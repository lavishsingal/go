// count the words in the string
// mapping string->count
package main

import (
	"strings"
	"code.google.com/p/go-tour/wc"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	parsestr := strings.Fields(s)

	for _ ,x := range  parsestr{
			count := m[x] 
			m[x]  = count+1
	}
	return map[string]int(m)
}

func main() {
	wc.Test(WordCount)
}