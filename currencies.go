package crptwav

type addressTypes struct {
	prod    []string
	testnet []string
}

type currencyParams struct {
	name           string
	symbol         string
	addressTypes   addressTypes
	hashFunc       string
	expectedLength int
}

type currency struct {
	params    currencyParams
	validator func(address, network string, params currencyParams) bool
}

// It defines P2PKH and P2SH address types for standard (prod)
// and testnet networks.
var currencies = []currency{
	currency{
		params: currencyParams{
			name:   "bitcoin",
			symbol: "btc",
			addressTypes: addressTypes{
				prod:    []string{"00", "05"},
				testnet: []string{"6f", "c4"},
			},
		},
		validator: isValidBitcoinAddress,
	},
	currency{
		params: currencyParams{
			name:   "bitcoincash",
			symbol: "bch",
			addressTypes: addressTypes{
				prod:    []string{"00", "05"},
				testnet: []string{"6f", "c4"},
			},
		},
		validator: isValidBitcoinAddress,
	},
	currency{
		params: currencyParams{
			name:   "litecoin",
			symbol: "ltc",
			addressTypes: addressTypes{
				prod:    []string{"30", "05", "32"},
				testnet: []string{"6f", "c4", "3a"},
			},
		},
		validator: isValidBitcoinAddress,
	},
	currency{
		params: currencyParams{
			name:   "ethereum",
			symbol: "eth",
		},
		validator: ethValidator,
	},
	currency{
		params: currencyParams{
			name:   "etherzero",
			symbol: "etz",
		},
		validator: ethValidator,
	},
	currency{
		params: currencyParams{
			name:   "ethereumclassic",
			symbol: "etc",
		},
		validator: ethValidator,
	},
	currency{
		params: currencyParams{
			name:   "callisto",
			symbol: "clo",
		},
		validator: ethValidator,
	},
}

func ethValidator(address, network string, params currencyParams) bool {
	return isValidEthereumAddress(address)
}
