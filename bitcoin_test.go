package crptwav

import "testing"

func TestIsValidBitcoin(t *testing.T) {
	tt := []struct {
		currency string
		address  string
		network  string
	}{
		{address: "12KYrjTdVGjFMtaxERSk3gphreJ5US8aUP", currency: "bitcoin", network: "prod"},
		{address: "12QeMLzSrB8XH8FvEzPMVoRxVAzTr5XM2y", currency: "bitcoin", network: "prod"},
		{address: "12QeMLzSrB8XH8FvEzPMVoRxVAzTr5XM2y", currency: "BTC", network: "prod"},
		{address: "12QeMLzSrB8XH8FvEzPMVoRxVAzTr5XM2y", currency: "Bitcoin", network: "prod"},
		{address: "12QeMLzSrB8XH8FvEzPMVoRxVAzTr5XM2y", currency: "btc", network: "prod"},
		{address: "12QeMLzSrB8XH8FvEzPMVoRxVAzTr5XM2y", currency: "btc", network: "prod"},
		{address: "12QeMLzSrB8XH8FvEzPMVoRxVAzTr5XM2y", currency: "btc", network: "both"},
		{address: "1oNLrsHnBcR6dpaBpwz3LSwutbUNkNSjs", currency: "bitcoin", network: "prod"},
		{address: "mzBc4XEFSdzCDcTxAgf6EZXgsZWpztRhef", currency: "bitcoin", network: "testnet"},
		{address: "mzBc4XEFSdzCDcTxAgf6EZXgsZWpztRhef", currency: "bitcoin", network: "both"},
		{address: "1SQHtwR5oJRKLfiWQ2APsAd9miUc4k2ez", currency: "bitcoin", network: "prod"},
		{address: "116CGDLddrZhMrTwhCVJXtXQpxygTT1kHd", currency: "bitcoin", network: "prod"},

		// p2sh addresses
		// invalid {address: "3NJZLcZEEYBpxYEUGewU4knsQRn1WM5Fkt", currency: "bitcoint", network: "prod"},
		{address: "2MxKEf2su6FGAUfCEAHreGFQvEYrfYNHvL7", currency: "bitcoin", network: "testnet"},
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
