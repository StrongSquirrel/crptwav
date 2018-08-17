package crptwav

import "testing"

func TestIsValidBitcoinGold(t *testing.T) {
	tt := []struct {
		currency string
		address  string
		network  string
	}{
		{address: "GW3JrQyHtoVfEFES3Y9JagiX3VSKQStLwj", currency: "bitcoingold", network: "prod"},
		{address: "GUDWdeMyAXQbrNFFivAhkJQ1GfBCFdc7JF", currency: "bitcoingold", network: "prod"},
		{address: "mvww6DEJ18dbyQUukpVQXvLgrNDJazZn1Y", currency: "bitcoingold", network: "testnet"},
		{address: "mn3mdEE6cf1snxVsknNz4GRTdSrWXqYp7c", currency: "BTG", network: "testnet"},
		{address: "GSNFPRsdaM3MXrU5HW1AxgFwmUQC8HXK9F", currency: "BTG", network: "prod"},
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
