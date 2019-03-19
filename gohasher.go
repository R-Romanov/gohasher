package gohasher

import (
	"errors"
	"hash/fnv"
)

func HashString(string string, hashMethod string) (int, error) {
	result := 0

	if hashMethod == "fnv1" {
		result, err := hashStringToFnv1Int64(string)
		if err != nil {
			return int(result), errors.New("can not hash given string")
		} else {
			return int(result), nil
		}
	} else {
		return int(result), errors.New("unknown hash method")
	}
}

func hashStringToFnv1Int64(url string) (int64, error) {
	h := fnv.New64a()
	_, err := h.Write([]byte(url))

	if err != nil {
		return 0, err
	}

	return int64(h.Sum64()), nil
}
