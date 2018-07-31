package crypto

import "golang.org/x/crypto/sha3"

// Keccak256 calculates and returns the Keccak256 hash of the input data.
func Keccak256(src []byte) []byte {
	h := sha3.NewLegacyKeccak256()
	h.Write(src)
	return h.Sum(nil)
}
