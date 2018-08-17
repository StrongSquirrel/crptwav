package crptwav

import "testing"

func TestIsValidDash(t *testing.T) {
	tt := []struct {
		currency string
		address  string
		network  string
	}{
		{address: "Xx4dYKgz3Zcv6kheaqog3fynaKWjbahb6b", currency: "dash", network: "prod"},
		{address: "XcY4WJ6Z2Q8w7vcYER1JypC8s2oa3SQ1b1", currency: "DASH", network: "prod"},
		{address: "XqMkVUZnqe3w4xvgdZRtZoe7gMitDudGs4", currency: "dash", network: "prod"},
		{address: "yPv7h2i8v3dJjfSH4L3x91JSJszjdbsJJA", currency: "dash", network: "testnet"},
	}
	for _, tc := range tt {
		switch tc.network {
		case NetworkProd:
			if !IsValidProdAddress(tc.address, tc.currency) {
				t.Errorf("Address %s should be valid %s %s address", tc.address, tc.currency, tc.network)
			}
		case NetworkTest:
			if !IsValidTestnetAddress(tc.address, tc.currency) {
				t.Errorf("Address %s should be valid %s %s address", tc.address, tc.currency, tc.network)
			}
		default:
			if !IsValidAddress(tc.address, tc.currency) {
				t.Errorf("Address %s should be valid %s %s address", tc.address, tc.currency, tc.network)
			}
		}
	}
}
