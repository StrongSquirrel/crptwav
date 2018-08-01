package crptwav

import "testing"

func TestIsValidBitcoinCash(t *testing.T) {
	tt := []struct {
		currency string
		address  string
		network  string
	}{
		{address: "12KYrjTdVGjFMtaxERSk3gphreJ5US8aUP", currency: "bitcoincash", network: "prod"},
		{address: "12QeMLzSrB8XH8FvEzPMVoRxVAzTr5XM2y", currency: "bitcoincash", network: "prod"},
		{address: "12QeMLzSrB8XH8FvEzPMVoRxVAzTr5XM2y", currency: "BCH", network: "prod"},
		{address: "12QeMLzSrB8XH8FvEzPMVoRxVAzTr5XM2y", currency: "Bitcoin", network: "prod"},
		{address: "12QeMLzSrB8XH8FvEzPMVoRxVAzTr5XM2y", currency: "bch", network: "prod"},
		{address: "12QeMLzSrB8XH8FvEzPMVoRxVAzTr5XM2y", currency: "bch", network: "prod"},
		{address: "12QeMLzSrB8XH8FvEzPMVoRxVAzTr5XM2y", currency: "bch", network: "both"},
		{address: "1oNLrsHnBcR6dpaBpwz3LSwutbUNkNSjs", currency: "bitcoincash", network: "prod"},
		{address: "mzBc4XEFSdzCDcTxAgf6EZXgsZWpztRhef", currency: "bitcoincash", network: "testnet"},
		{address: "mzBc4XEFSdzCDcTxAgf6EZXgsZWpztRhef", currency: "bitcoincash", network: "both"},

		// p2sh addresses
		{address: "3NJZLcZEEYBpxYEUGewU4knsQRn1WM5Fkt", currency: "bitcoincash", network: "prod"},
		{address: "2MxKEf2su6FGAUfCEAHreGFQvEYrfYNHvL7", currency: "bitcoincash", network: "testnet"},
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
