package utils

import (
	"math/rand"
	"strconv"
	"strings"
)

func GenarateCode(len int) string {
	var array = make([]string, len)

	for i := range array {
		n := strconv.Itoa(rand.Intn(9))
		array[i] = n
	}

	return strings.Join(array, "")
}
