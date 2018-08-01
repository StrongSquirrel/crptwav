package crptwav

import "strings"

const (
	// NetworkProd is used to validate main network addresses.
	NetworkProd = "prod"

	// NetworkTest is used to validate test network addresses.
	NetworkTest = "testnet"

	// NetworkBoth is used to validate all network addresses.
	NetworkBoth = "both"
)

// IsValidAddress validates the given currency main or test network address.
// Currency can be a name or a symbol, e.g. "bitcoin", "litecoin" or "ETH".
func IsValidAddress(address, currency string) bool {
	return isValid(address, currency, NetworkBoth)
}

// IsValidProdAddress validates the given currency main network address.
// Currency can be a name or a symbol, e.g. "bitcoin", "litecoin" or "ETH".
func IsValidProdAddress(address, currency string) bool {
	return isValid(address, currency, NetworkProd)
}

// IsValidTestnetAddress validates the given currency test network address.
// Currency can be a name or a symbol, e.g. "bitcoin", "litecoin" or "ETH".
func IsValidTestnetAddress(address, currency string) bool {
	return isValid(address, currency, NetworkTest)
}

// isValid validates the given currency address.
// Currency can be a name or a symbol, e.g. "bitcoin", "litecoin" or "ETH".
// Network can be "prod" to enforce standard address, "testnet" to enforce
// testnet address and "both" to enforce nothing.
func isValid(address, currency, network string) bool {
	nameOrSymbol := strings.ToLower(currency)
	for _, c := range currencies {
		if c.params.name == nameOrSymbol || c.params.symbol == nameOrSymbol {
			return c.validator(address, network, c.params)
		}
	}
	return false
}
