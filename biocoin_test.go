package crptwav

import "testing"

func TestIsValidBioCoin(t *testing.T) {
	tt := []struct {
		currency string
		address  string
		network  string
	}{
		{address: "B7xseoLGk7hEpMDDeSvZDKmmiAMHWiccok", currency: "biocoin", network: "prod"},
		{address: "B8zjmYFGhWmiaQSJshfrnefE72xCapCkvo", currency: "biocoin", network: "prod"},
		{address: "muH8LL42DiMs8GEQ6Grfi8KUw2uFvuKr1J", currency: "biocoin", network: "testnet"},
		{address: "muH8LL42DiMs8GEQ6Grfi8KUw2uFvuKr1J", currency: "BIO", network: "testnet"},
		{address: "B8zjmYFGhWmiaQSJshfrnefE72xCapCkvo", currency: "BIO", network: "prod"},
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
