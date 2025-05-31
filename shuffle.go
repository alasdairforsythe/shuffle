// Package shuffle provides extremely fast pseudo-random shuffling
// using a Linear Congruential Generator (LCG) for maximum performance.
package shuffle

import (
	"time"
)

var seed uint64 = uint64(time.Now().UnixNano())

func Seed() uint64 {
	seed = seed * 1103515245 + 12345
	return seed
}

func Shuffle[T any](seed uint64, slice []T) {
	n := len(slice)
	for i := uint64(n - 1); i > 0; i-- {
		seed = seed*1664525 + 1013904223
		j := seed % (i+1)
		slice[i], slice[j] = slice[j], slice[i]
	}
}
