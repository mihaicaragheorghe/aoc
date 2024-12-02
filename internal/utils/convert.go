package utils

import (
	"log"
	"strconv"
)

func S2i(s string) int {
	i, e := strconv.Atoi(s)
	if e != nil {
		log.Fatal(e)
	}
	return i
}

func B2i(b bool) int {
	if b {
		return 1
	}
	return 0
}
