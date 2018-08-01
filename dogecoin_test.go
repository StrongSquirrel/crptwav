package crptwav

import "testing"

func TestIsValidDogecoin(t *testing.T) {
	tt := []struct {
		currency string
		address  string
		network  string
	}{
		{address: "DPpJVPpvPNP6i6tMj4rTycAGh8wReTqaSU", currency: "dogecoin", network: "prod"},
		{address: "DNzLUN6MyYVS5zf4Xc2yK69V3dXs6Mxia5", currency: "dogecoin", network: "prod"},
		{address: "DPS6iZj7roHquvwRYXNBua9QtKPzigUUhM", currency: "dogecoin", network: "prod"},
		{address: "DPS6iZj7roHquvwRYXNBua9QtKPzigUUhM", currency: "DOGE", network: "prod"},

		// p2sh addresses
		{address: "A7JjzK9k9x5b2MkkQzqt91WZsuu7wTu6iS", currency: "dogecoin", network: "prod"},
		{address: "2MxKEf2su6FGAUfCEAHreGFQvEYrfYNHvL7", currency: "dogecoin", network: "testnet"},
	}
	for _, tc := range tt {
		if !IsValid(tc.address, tc.currency, tc.network) {
			t.Errorf("Address %s should be valid %s %s address", tc.address, tc.currency, tc.network)
		}
	}
}
