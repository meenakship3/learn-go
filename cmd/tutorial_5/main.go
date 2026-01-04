package main

import (
	"fmt"
	"strings"
)

func main() {
	var strSlice = []string{"s", "e", "n", "t", "i", "m", "e", "n", "t"}
	var strBuilder strings.Builder
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i]) // each char added to the builder but string created later
	}
	var catStr = strBuilder.String()
	fmt.Printf("\n%v", catStr)

}
