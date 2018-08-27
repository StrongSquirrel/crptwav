package crptwav

import "testing"

func TestIsValidBitcloud(t *testing.T) {
	tt := []struct {
		currency string
		address  string
		network  string
	}{
		{address: "B5xbPCtDHBksZwDsgSssnmuSx6ydNLBQEG", currency: "bitcloud", network: "prod"},
		{address: "BTiVBSbBZ2AkvMt8nZ3WZ9Vpb4qEPDtDGz", currency: "bitcloud", network: "prod"},
		{address: "BDt8xFZo2Y8XMzpqQRWrtci3v3MU9ijVmh", currency: "bitcloud", network: "prod"},
		{address: "BTMTdv3KQ73h3jkBCCRGH9yMwAC1bDF2Vc", currency: "BTDX", network: "prod"},
		{address: "B6mVF3ijTvAqu3U1KX1xuF8CkeqMUqGeDf", currency: "BTDX", network: "prod"},
		{address: "B7Dhb12dBdJhRkGEDNKcjSau8h1e6bEw7o", currency: "BTDX", network: "prod"},
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
