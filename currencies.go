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
			name:   "auroracoin",
			symbol: "aur",
			addressTypes: addressTypes{
				prod:    []string{"17", "05"},
				testnet: []string{"6f", "c4"},
			},
		},
		validator: isValidBitcoinAddress,
	},
	currency{
		params: currencyParams{
			name:   "beavercoin",
			symbol: "bvc",
			addressTypes: addressTypes{
				prod:    []string{"19", "05"},
				testnet: []string{"6f", "c4"},
			},
		},
		validator: isValidBitcoinAddress,
	},
	currency{
		params: currencyParams{
			name:   "biocoin",
			symbol: "bio",
			addressTypes: addressTypes{
				prod:    []string{"19", "14"},
				testnet: []string{"6f", "c4"},
			},
		},
		validator: isValidBitcoinAddress,
	},
	currency{
		params: currencyParams{
			name:   "bitcoingold",
			symbol: "btg",
			addressTypes: addressTypes{
				prod:    []string{"26", "17"},
				testnet: []string{"6f", "c4"},
			},
		},
		validator: isValidBitcoinAddress,
	},
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
			name:   "dash",
			symbol: "dash",
			addressTypes: addressTypes{
				prod:    []string{"4c", "10"},
				testnet: []string{"8c", "13"},
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
			name:   "peercoin",
			symbol: "ppc",
			addressTypes: addressTypes{
				prod:    []string{"37", "75"},
				testnet: []string{"6f", "c4"},
			},
		},
		validator: isValidBitcoinAddress,
	},
	currency{
		params: currencyParams{
			name:   "dogecoin",
			symbol: "doge",
			addressTypes: addressTypes{
				prod:    []string{"1e", "16"},
				testnet: []string{"71", "c4"},
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
