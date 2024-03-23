package hash

import (
	"fmt"

	"github.com/cespare/xxhash"
)

// XXH64Sum returns the xxhash64 sum of the input
func XXH64Sum(in []byte) (uint64, error) {
	xxh := xxhash.New()
	if _, err := xxh.Write(in); err != nil {
		return 0, err
	}

	return xxh.Sum64(), nil
}

// XXH64SumString returns the string of the hex encoded hash sum of the input
func XXH64SumString(in []byte) (string, error) {
	s, err := XXH64Sum(in)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", s), nil
}

// XXH64SumBytes returns the bytes of the hex encoded hash sum of the input
func XXH64SumBytes(in []byte) ([]byte, error) {
	s, err := XXH64Sum(in)
	if err != nil {
		return nil, err
	}
	return []byte(fmt.Sprintf("%x", s)), nil
}
