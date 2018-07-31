package keccak256

import "golang.org/x/crypto/sha3"

// Encode calculates and returns the Keccak256 hash of the input data.
func Encode(src []byte) []byte {
	h := sha3.NewLegacyKeccak256()
	h.Write(src)
	return h.Sum(nil)
}
