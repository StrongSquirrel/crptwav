package crptwav

import "testing"

func TestIsValidBitcore(t *testing.T) {
	tt := []struct {
		currency string
		address  string
		network  string
	}{
		{address: "2Wt3K47XEAANg9vMSWPmLSgfDFzBVNdS3T", currency: "bitcore", network: "prod"},
		{address: "2Yr4uUHsunFATv75RSvbbxS2YDpU6iCu9E", currency: "bitcore", network: "prod"},
		{address: "mx8qEXzx9Aeog4c4WbhM3eKMp3Uoh89UA4", currency: "bitcore", network: "testnet"},
		{address: "mkHJdVjP3yZuz9ZK8kwJ16kxUC94iLSoWp", currency: "BTX", network: "testnet"},
		{address: "2Vkxa2FHBdhnhBTVEGdKQgiDvCbY9iCBqc", currency: "BTX", network: "prod"},
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
