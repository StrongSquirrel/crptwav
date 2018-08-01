package crptwav

import "testing"

func TestIsValidPeercoin(t *testing.T) {
	tt := []struct {
		currency string
		address  string
		network  string
	}{
		{address: "PHCEsP6od3WJ8K2WKWEDBYKhH95pc9kiZN", currency: "peercoin", network: "prod"},
		{address: "PSbM1pGoE9dnAuVWvpQqTTYVpKZU41dNAz", currency: "peercoin", network: "prod"},
		{address: "PUULeHrJL2WujJkorc2RsUAR3SardKUauu", currency: "peercoin", network: "prod"},
		{address: "PUULeHrJL2WujJkorc2RsUAR3SardKUauu", currency: "PPC", network: "prod"},
		{address: "mzBc4XEFSdzCDcTxAgf6EZXgsZWpztRhef", currency: "peercoin", network: "testnet"},

		// p2sh addresses
		{address: "pNms4CaWqgZUxbNZaA1yP2gPr3BYnez9EM", currency: "peercoin", network: "prod"},
		{address: "2MxKEf2su6FGAUfCEAHreGFQvEYrfYNHvL7", currency: "peercoin", network: "testnet"},
	}
	for _, tc := range tt {
		if !IsValid(tc.address, tc.currency, tc.network) {
			t.Errorf("Address %s should be valid %s %s address", tc.address, tc.currency, tc.network)
		}
	}
}
