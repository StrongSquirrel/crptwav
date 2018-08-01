package crptwav

import (
	"encoding/hex"
	"errors"

	"github.com/StrongSquirrel/crptwav/crypto/base58"
)

func isValidBitcoinAddress(address, network string, params currencyParams) bool {
	t, err := bitcoinAddressType(address, params)
	if err != nil {
		return false
	}

	var addressTypes []string
	switch network {
	case NetworkProd:
		addressTypes = params.addressTypes.prod
	case NetworkTest:
		addressTypes = params.addressTypes.testnet
	default:
		addressTypes = append(params.addressTypes.prod, params.addressTypes.testnet...)
	}

	for _, v := range addressTypes {
		if v == t {
			return true
		}
	}

	return false
}

func bitcoinAddressType(address string, params currencyParams) (string, error) {
	expectedLength := 25
	if params.expectedLength != 0 {
		expectedLength = params.expectedLength
	}

	decoded := base58.Decode(address)

	length := len(decoded)
	if length != expectedLength {
		return "", errors.New("invalid decoded address length")
	}

	decoded, version, err := base58.CheckDecode(address)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString([]byte{version}), nil
}
