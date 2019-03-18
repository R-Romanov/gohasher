package main

import (
	"fmt"
	"os"
)
import "hash/fnv"

func hashStringToInt64(url string) (int64, error) {
	h := fnv.New64a()
	_, err := h.Write([]byte(url))

	if err != nil {
		return 0, err
	}

	return int64(h.Sum64()), nil
}


func main() {
	var inputData string

	if len(os.Args[1:]) > 0 {
		inputData = os.Args[1:][0]
	} else {
		panic("you must specify at least one argument")
		return
	}

	test, err := hashStringToInt64(inputData)

	if err != nil {
		panic("can not hash input data ")
		return
	}

	fmt.Printf("%v", test)
}
