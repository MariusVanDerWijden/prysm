package fuzz_utils

import "math/rand"

// /////////////////////////////////////////////////////////////////////////////
// Mutators
// /////////////////////////////////////////////////////////////////////////////

func MutateNBytes(b *[]byte, n int) *[]byte {
	// make a copy of the slice
	//new_slice := make([]byte, len(*b))
	//new_ptr := &new_slice
	//copy(*new_ptr, *b)

	// mutate the copy
	for i := 0; i < n; i++ {
		// iterate through bytes, 50% of the time mutate the byte
		if rand.Intn(2) == 0 {
			(*b)[i] = byte(rand.Intn(256))
		}
	}
	return b
}