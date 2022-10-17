package main

import (
	"fmt"
	"sort"
	"strings"
)

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func foldString(s string) string {
	var (
		result   string
		counter  int
		lastChar int32
	)
	s = sortString(s)
	for _, chr := range s {
		if chr != lastChar && counter != 0 {
			result += fmt.Sprintf("%v%v", string(lastChar), counter)
			counter = 0
		}

		lastChar = chr
		counter++
	}
	result += fmt.Sprintf("%v%v", string(lastChar), counter)
	return result
}

func main() {
	var input string
	fmt.Scanf("%s\n", &input)
	fmt.Println(foldString(input))
}
