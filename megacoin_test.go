package crptwav

import "testing"

func TestIsValidMegacoin(t *testing.T) {
	tt := []struct {
		currency string
		address  string
		network  string
	}{
		{address: "MMGut25spuzpNZuWjXnmZaSW8Ym4QaeNEx", currency: "megacoin", network: "prod"},
		{address: "MRCEC5KG7dDwiDhnR4CkbZFBnQW3DnkE9P", currency: "megacoin", network: "prod"},
		{address: "MH17hNLVPbcDjUAppyV385zAwYXCNSNuBj", currency: "megacoin", network: "prod"},
		{address: "M9P8kjJRAA1XyWMC5wm1UAEZD2BMBxLWvH", currency: "MEC", network: "prod"},
		{address: "MSLs5yKxFzsBjyxUWqftx2mcBsDPXibeTE", currency: "MEC", network: "prod"},
		{address: "MEhNXKErNV1zMzA9dApxuGBzGs7tqWJRQr", currency: "MEC", network: "prod"},
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
