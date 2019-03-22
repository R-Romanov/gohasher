package hasher

import (
	"errors"
	"hash/fnv"
)

const (
	Algorithm64BitFnv1 = "64bit_fnv1"
	AlgorithmDefault   = Algorithm64BitFnv1
)

var (
	hashFunc = map[string]interface{}{
		Algorithm64BitFnv1: hashStringToFnv1Int64,
	}
)

func Hash(str string) (uint64, error) {

	result, err := hashFunc[AlgorithmDefault].(func(string) (uint64, error))(str)

	if err != nil {
		return 0, errors.New("can not hash given string")
	}

	return result, nil
}

func hashStringToFnv1Int64(url string) (uint64, error) {
	h := fnv.New64a()
	_, err := h.Write([]byte(url))

	if err != nil {
		return 0, err
	}

	return h.Sum64(), nil
}
