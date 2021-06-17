package main

import (
	"math/rand"
	"reflect"
	"time"
)

func ShuffleArrayStrings(a []string) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
}

func Contains(slice []string, e string) bool {
	for _, a := range slice {
		if a == e {
			return true
		}
	}
	return false
}

func Remove(slice []string, toRemove string) []string {
	for i, v := range slice {
		if v == toRemove {
			slice = append(slice[:i], slice[i+1:]...)
			break
		}
	}

	return slice
}

func Prepend(data []string, item string) []string {
	data = append([]string{item}, data...)
	return data
}

func PrependSlice(data [][]string, item []string) [][]string {
	data = append([][]string{item}, data...)
	return data
}

func GetIntStringKeys(objmap map[string]int) []string {
	keys := make([]string, 0, len(objmap))
	for k := range objmap {
		keys = append(keys, k)
	}

	return keys
}

func Shift(a []string) (string, []string) {
	x, a := a[0], a[1:]

	return x, a
}

func Filter(slice []string, filter []string) []string {
	result := []string{}
	for _, k := range slice {
		if !Contains(filter, k) {
			result = append(result, k)
		}
	}

	return result
}

func Factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}

func Combinations(n int, r int) uint64 {
	un := uint64(n)
	ur := uint64(r)

	if un < ur {
		return 0
	}

	if un == ur {
		return 0
	}

	return Factorial(un) / (Factorial(ur) * Factorial(un-ur))
}

func reverseAny(s interface{}) {
	n := reflect.ValueOf(s).Len()
	swap := reflect.Swapper(s)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}
