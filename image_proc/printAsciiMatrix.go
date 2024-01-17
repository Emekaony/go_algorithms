package image_proc

import (
	"fmt"
	"strings"
)

// this whole code better work man!
func PrintAsciiMatrix(asciiMatrix [][]string) {
	for _, row := range asciiMatrix {
		line := []string{}
		for _, ch := range row {
			line = append(line, strings.Repeat(ch, 3)) // repeat it 3 times
		}
		fmt.Println(strings.Join(line, ""))
	}
}
