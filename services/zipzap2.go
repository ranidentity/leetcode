package services

import (
	"fmt"
	"strings"
)

func Convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	rows := make([]strings.Builder, numRows)
	cycle := 2*numRows - 2
	fmt.Printf("%d\n", cycle)
	for i, c := range s {
		pos := i % cycle
		if pos < numRows {
			rows[pos].WriteRune(c)
		} else {
			rows[cycle-pos].WriteRune(c)
		}
	}

	var result strings.Builder
	for _, row := range rows {
		fmt.Println(row.String())
		result.WriteString(row.String())
	}
	return result.String()

}
