package crptwav

import "testing"

func TestIsValidEthereum(t *testing.T) {
	tt := []struct {
		currency string
		address  string
	}{
		{address: "0xE37c0D48d68da5c5b14E5c1a9f1CFE802776D9FF", currency: "ethereum"},
		{address: "0xa00354276d2fC74ee91e37D085d35748613f4748", currency: "ethereum"},
		{address: "0xAff4d6793F584a473348EbA058deb8caad77a288", currency: "ETH"},
		{address: "0xc6d9d2cd449a754c494264e1809c50e34d64562b", currency: "ETH"},
		{address: "0x52908400098527886E0F7030069857D2E4169EE7", currency: "ETH"},
		{address: "0x8617E340B3D01FA5F11F306F4090FD50E238070D", currency: "ETH"},
		{address: "0xde709f2102306220921060314715629080e2fb77", currency: "ETH"},
		{address: "0x27b1fdb04752bbc536007a920d24acb045561c26", currency: "ETH"},
		{address: "0x5aAeb6053F3E94C9b9A09f33669435E7Ef1BeAed", currency: "ETH"},
		{address: "0xfB6916095ca1df60bB79Ce92cE3Ea74c37c5d359", currency: "ETH"},
		{address: "0xdbF03B407c01E7cD3CBea99509d93f8DDDC8C6FB", currency: "ETH"},
		{address: "0xD1220A0cf47c7B9Be7A2E6BA89F429762e7b9aDb", currency: "ETH"},
		{address: "0xD1220A0cf47c7B9Be7A2E6BA89F429762e7b9aDb", currency: "ethereumclassic"},
		{address: "0xD1220A0cf47c7B9Be7A2E6BA89F429762e7b9aDb", currency: "ETC"},
		{address: "0xD1220A0cf47c7B9Be7A2E6BA89F429762e7b9aDb", currency: "etherzero"},
		{address: "0xD1220A0cf47c7B9Be7A2E6BA89F429762e7b9aDb", currency: "ETZ"},
		{address: "0xD1220A0cf47c7B9Be7A2E6BA89F429762e7b9aDb", currency: "callisto"},
		{address: "0xD1220A0cf47c7B9Be7A2E6BA89F429762e7b9aDb", currency: "CLO"},
	}
	for _, tc := range tt {
		if !IsValid(tc.address, tc.currency, "both") {
			t.Errorf("Address %s should be valid %s address", tc.address, tc.currency)
		}
	}
}
