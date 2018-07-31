package crptwav

import "strings"

// IsValid validates the given currency address.
// address - Wallet address to validate.
// currency - Currency name or symbol, e.g. "bitcoin", "litecoin" or "ETH".
// network - Use "prod" to enforce standard address, "testnet" to enforce
// testnet address and "both" to enforce nothing.
func IsValid(address, currency, network string) bool {
	nameOrSymbol := strings.ToLower(currency)
	for _, c := range currencies {
		if c.name == nameOrSymbol || c.symbol == nameOrSymbol {
			return c.validator(address, network)
		}
	}
	return false
}
