package crptwav

type currency struct {
	name      string
	symbol    string
	validator func(address, network string) bool
}

// It defines P2PKH and P2SH address types for standard (prod)
// and testnet networks.
var currencies = []currency{
	currency{
		name:      "ethereum",
		symbol:    "eth",
		validator: ethValidator,
	},
	currency{
		name:      "etherzero",
		symbol:    "etz",
		validator: ethValidator,
	},
	currency{
		name:      "ethereumclassic",
		symbol:    "etc",
		validator: ethValidator,
	},
	currency{
		name:      "callisto",
		symbol:    "clo",
		validator: ethValidator,
	},
}

func ethValidator(address, network string) bool {
	return isValidETH(address)
}
