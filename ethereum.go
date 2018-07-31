package crptwav

import (
	"encoding/hex"
	"regexp"
	"strconv"
	"strings"

	"github.com/StrongSquirrel/crptwav/crypto"
)

func isValidEthereumAddress(address string) bool {
	baseRequirements, _ := regexp.Compile("^0x[0-9a-fA-F]{40}$")
	if !baseRequirements.MatchString(address) {
		return false
	}

	allSmall, _ := regexp.Compile("^0x[0-9a-f]{40}$")
	if allSmall.MatchString(address) {
		return true
	}

	allCaps, _ := regexp.Compile("^0x?[0-9A-F]{40}$")
	if allCaps.MatchString(address) {
		return true
	}

	// Otherwise check each case
	return isValidChecksum(address)
}

func isValidChecksum(address string) bool {
	address = strings.Replace(address, "0x", "", -1)

	hash := hex.EncodeToString(crypto.Keccak256([]byte(strings.ToLower(address))))

	for i := 0; i < 40; i++ {
		num, _ := strconv.ParseInt(string(hash[i]), 16, 8)
		sign := string(address[i])

		if num > 7 && strings.ToUpper(sign) != sign {
			return false
		}

		if num <= 7 && strings.ToLower(sign) != sign {
			return false
		}
	}

	return true
}
