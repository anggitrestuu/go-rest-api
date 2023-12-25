package utils

import "strconv"

func StringToUint(s string) uint {
	i, _ := strconv.Atoi(s)
	return uint(i)
}

func StringToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
