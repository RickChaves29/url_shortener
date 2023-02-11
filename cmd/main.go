package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("genarateCode: %v\n", genarateCode(3))
	fmt.Printf("genarateCode: %v\n", genarateCode(6))
	fmt.Printf("genarateCode: %v\n", genarateCode(9))
}

func genarateCode(len int) string {
	var array = make([]string, len)

	for i := range array {
		n := strconv.Itoa(rand.Intn(9))
		array[i] = n
	}

	return strings.Join(array, "")
}
