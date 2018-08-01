package crptwav

import "testing"

func TestIsValidLitecoin(t *testing.T) {
	tt := []struct {
		currency string
		address  string
		network  string
	}{
		{address: "LVg2kJoFNg45Nbpy53h7Fe1wKyeXVRhMH9", currency: "litecoin", network: "prod"},
		{address: "LVg2kJoFNg45Nbpy53h7Fe1wKyeXVRhMH9", currency: "LTC", network: "prod"},
		{address: "LTpYZG19YmfvY2bBDYtCKpunVRw7nVgRHW", currency: "litecoin", network: "prod"},
		{address: "Lb6wDP2kHGyWC7vrZuZAgV7V4ECyDdH7a6", currency: "litecoin", network: "prod"},
		{address: "mzBc4XEFSdzCDcTxAgf6EZXgsZWpztRhef", currency: "litecoin", network: "testnet"},

		// p2sh addresses
		{address: "3NJZLcZEEYBpxYEUGewU4knsQRn1WM5Fkt", currency: "litecoin", network: "prod"},
		{address: "2MxKEf2su6FGAUfCEAHreGFQvEYrfYNHvL7", currency: "litecoin", network: "testnet"},
		{address: "QW2SvwjaJU8LD6GSmtm1PHnBG2xPuxwZFy", currency: "litecoin", network: "testnet"},
	}
	for _, tc := range tt {
		if !IsValid(tc.address, tc.currency, tc.network) {
			t.Errorf("Address %s should be valid %s %s address", tc.address, tc.currency, tc.network)
		}
	}
}
