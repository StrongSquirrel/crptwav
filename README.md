# crptwav

Simple wallet address validator for validating Bitcoin and other altcoins addresses.

It is a port of JavaScript library [https://github.com/ognus/wallet-address-validator](https://github.com/ognus/wallet-address-validator).

## Documentation

[https://godoc.org/github.com/StrongSquirrel/crptwav](https://godoc.org/github.com/StrongSquirrel/crptwav)

## Example

```go
package main

import (
    "fmt"
    "github.com/StrongSquirrel/crptwav"
)

func main() {
    fmt.Println(crptwav.IsValid("0xE37c0D48d68da5c5b14E5c1a9f1CFE802776D9FF", "ETH", "both"))
}
```

### Supported crypto currencies (todo)

- [ ] Auroracoin/AUR, `'auroracoin'` or `'AUR'`
- [ ] BeaverCoin/BVC, `'beavercoin'` or `'BVC'`
- [ ] Biocoin/BIO, `'biocoin'` or `'BIO'`
- [x] Bitcoin/BTC, `'bitcoin'` or `'BTC'`
- [x] BitcoinCash/BCH, `'bitcoincash'` or `'BCH'`
- [ ] BitcoinGold/BTG, `'bitcoingold'` or `'BTG'`
- [ ] BitcoinPrivate/BTCP, `'bitcoinprivate'` or `'BTCP'`
- [ ] BitcoinZ/BTCZ, `'bitcoinz'` or `'BTCZ'`
- [x] Callisto/CLO, `'callisto'` or `'CLO'`
- [ ] Dash, `'dash'` or `'DASH'`
- [ ] Decred/DCR, `'decred'` or `'DCR'`
- [ ] Digibyte/DGB, `'digibyte'` or `'DGB'`
- [ ] Dogecoin/DOGE, `'dogecoin'` or `'DOGE'`
- [x] Ethereum/ETH, `'ethereum'` or `'ETH'`
- [x] EthereumClassic/ETH, `'ethereumclassic'` or `'ETC'`
- [x] EthereumZero/ETZ, `'etherzero'` or `'ETZ'`
- [ ] Freicoin/FRC, `'freicoin'` or `'FRC'`
- [ ] Garlicoin/GRLC, `'garlicoin'` or `'GRLC'`
- [ ] Hush/HUSH, `'hush'` or `'HUSH'`
- [ ] Komodo/KMD, `'komodo'` or `'KMD'`
- [ ] Litecoin/LTC, `'litecoin'` or `'LTC'`
- [ ] Megacoin/MEC, `'megacoin'` or `'MEC'`
- [ ] Namecoin/NMC, `'namecoin'` or `'NMC'`
- [ ] NEO/NEO, `'NEO'` or `'NEO'`
- [ ] NeoGas/GAS, `'neogas'` or `'GAS'`
- [ ] Peercoin/PPCoin/PPC, `'peercoin'` or `'PPC'`
- [ ] Primecoin/XPM, `'primecoin'` or `'XPM'`
- [ ] Protoshares/PTS, `'protoshares'` or `'PTS'`
- [ ] Qtum/QTUM, `'qtum'` or `'QTUM'`
- [ ] Ripple/XRP, `'ripple'` or `'XRP'`
- [ ] Snowgem/SNG, `'snowgem'` or `'SNG'`
- [ ] Vertcoin/VTC, `'vertcoin'` or `'VTC'`
- [ ] Votecoin/VTC, `'votecoin'` or `'VOT'`
- [ ] Zcash/ZEC, `'zcash'` or `'ZEC'`
- [ ] Zclassic/ZCL, `'zclassic'` or `'ZCL'`
- [ ] ZenCash/ZEN, `'zencash'` or `'ZEN'`

## License

[MIT](LICENSE)
