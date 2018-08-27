package crptwav

import "testing"

func TestIsValidBitsend(t *testing.T) {
	tt := []struct {
		currency string
		address  string
		network  string
	}{
		{address: "i5m5mDHP8amxzPnqqKbk9Fh4xuRUpC77KU", currency: "bitsend", network: "prod"},
		{address: "i8VWuqDJEmQVG7ybxr7LrYN8QXB59ed1vg", currency: "bitsend", network: "prod"},
		{address: "i9JSd3SQmurL8z7YaWi7nZfGkow1pfN3Gt", currency: "bitsend", network: "prod"},
		{address: "iDcDWJ8QywriMGkMaDs5PqxrbrBjvr9rWj", currency: "BSD", network: "prod"},
		{address: "iSmk3ZcRAzM3NF2xB6xP8U58HCwpVnsJDY", currency: "BSD", network: "prod"},
		{address: "iSk6evb7HsRJmStTWeRiAMVDM1yNoGSEFd", currency: "BSD", network: "prod"},
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
