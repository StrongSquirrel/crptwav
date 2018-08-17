package crptwav

import "testing"

func TestIsValidAuroracoin(t *testing.T) {
	tt := []struct {
		currency string
		address  string
		network  string
	}{
		{address: "ARM3GLZXF1PDTZ5vz3wh5MVahbK9BHTWAN", currency: "auroracoin", network: "prod"},
		{address: "AUtfc6ThCLb7FuEu7QPrWpJuaXaJRPciDF", currency: "auroracoin", network: "prod"},
		{address: "AUN1oaj5hjispGnczt8Aruw3TxgGyRqB3V", currency: "auroracoin", network: "prod"},
		{address: "AXGcBkGX6NiaDXj85C5dCrhTRUgwxSkKDK", currency: "auroracoin", network: "prod"},
		{address: "AXGcBkGX6NiaDXj85C5dCrhTRUgwxSkKDK", currency: "AUR", network: "prod"},
		{address: "mzBc4XEFSdzCDcTxAgf6EZXgsZWpztRhef", currency: "auroracoin", network: "testnet"},

		// p2sh addresses
		{address: "3NJZLcZEEYBpxYEUGewU4knsQRn1WM5Fkt", currency: "auroracoin", network: "prod"},
		{address: "2MxKEf2su6FGAUfCEAHreGFQvEYrfYNHvL7", currency: "auroracoin", network: "testnet"},
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
