package crptwav

import "testing"

func TestIsValidBeavercoin(t *testing.T) {
	tt := []struct {
		currency string
		address  string
		network  string
	}{
		{address: "BPPtB4EpPi5wCaGXZuNyKQgng8ya579qUh", currency: "beavercoin", network: "prod"},
		{address: "BC1LLYoE4mTCHTJhVYvLGxhRTwAHyWTQ49", currency: "beavercoin", network: "prod"},
		{address: "BBuyeg2vjtyFdMNj3LTxuVra4wJMKVAY9C", currency: "beavercoin", network: "prod"},
		{address: "BBuyeg2vjtyFdMNj3LTxuVra4wJMKVAY9C", currency: "BVC", network: "prod"},
		{address: "mzBc4XEFSdzCDcTxAgf6EZXgsZWpztRhef", currency: "beavercoin", network: "testnet"},

		// p2sh addresses
		{address: "3NJZLcZEEYBpxYEUGewU4knsQRn1WM5Fkt", currency: "beavercoin", network: "prod"},
		{address: "2MxKEf2su6FGAUfCEAHreGFQvEYrfYNHvL7", currency: "beavercoin", network: "testnet"},
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
