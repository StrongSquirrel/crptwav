package crptwav

import "strings"

// IsValid validates the given currency address.
// Currency can be a name or a symbol, e.g. "bitcoin", "litecoin" or "ETH".
// Network can be "prod" to enforce standard address, "testnet" to enforce
// testnet address and "both" to enforce nothing.
func IsValid(address, currency, network string) bool {
	nameOrSymbol := strings.ToLower(currency)
	for _, c := range currencies {
		if c.params.name == nameOrSymbol || c.params.symbol == nameOrSymbol {
			return c.validator(address, network, c.params)
		}
	}
	return false
}
