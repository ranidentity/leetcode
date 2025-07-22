package services

import (
	"fmt"
	"strconv"
)

func IntToBinarySprintf(n int) string {
	return fmt.Sprintf("%b", n)
}

func BinaryToInt(s string) (int, error) {
	num, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		return 0, err
	}
	return int(num), nil
}
