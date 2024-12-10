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

func SliceAtoi(sa []string) []int {
	si := make([]int, 0, len(sa))
	for _, a := range sa {
		i, err := strconv.Atoi(a)
		if err != nil {
			panic(err)
		}
		si = append(si, i)
	}
	return si
}

func SliceAtoi64(sa []string) []int64 {
	si := make([]int64, 0, len(sa))
	for _, a := range sa {
		i, err := strconv.ParseInt(a, 10, 64)
		if err != nil {
			panic(err)
		}
		si = append(si, i)
	}
	return si
}

func Atoi64(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return i
}
