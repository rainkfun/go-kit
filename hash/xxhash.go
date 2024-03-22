package hash

import (
	"fmt"

	"github.com/cespare/xxhash"
)

var (
	xxh = xxhash.New()
)

func XXH64Sum(s string) (uint64, error) {
	xxh.Reset()
	if _, err := xxh.Write([]byte(s)); err != nil {
		return 0, err
	}

	return xxh.Sum64(), nil
}

func XXH64SumString(s string) (string, error) {
	xxh.Reset()
	if _, err := xxh.Write([]byte(s)); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", xxh.Sum64()), nil
}
