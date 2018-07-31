package crypto

import (
	"encoding/hex"
	"strings"
	"testing"
)

func TestKeccak256(t *testing.T) {
	tt := []struct {
		desc   string
		input  []byte
		repeat int // input will be concatenated the input this many times.
		want   string
	}{
		// Inputs of 8, 248, and 264 bits from http://keccak.noekeon.org/ are included below.
		{
			desc:   "short-8b",
			input:  decodeHex("CC"),
			repeat: 1,
			want:   "EEAD6DBFC7340A56CAEDC044696A168870549A6A7F6F56961E84A54BD9970B8A",
		},
		{
			desc:   "short-248b",
			input:  decodeHex("84FB51B517DF6C5ACCB5D022F8F28DA09B10232D42320FFC32DBECC3835B29"),
			repeat: 1,
			want:   "D477FB02CAAA95B3280EC8EE882C29D9E8A654B21EF178E0F97571BF9D4D3C1C",
		},
		{
			desc:   "short-264b",
			input:  decodeHex("DE8F1B3FAA4B7040ED4563C3B8E598253178E87E4D0DF75E4FF2F2DEDD5A0BE046"),
			repeat: 1,
			want:   "E78C421E6213AFF8DE1F025759A4F2C943DB62BBDE359C8737E19B3776ED2DD2",
		},
		// The computed test vector is 64 MiB long and is a truncated version
		// of the ExtremelyLongMsgKAT taken from http://keccak.noekeon.org/.
		{
			desc:   "long-64MiB",
			input:  []byte("abcdefghbcdefghicdefghijdefghijkefghijklfghijklmghijklmnhijklmno"),
			repeat: 1024 * 1024,
			want:   "5015A4935F0B51E091C6550A94DCD262C08998232CCAA22E7F0756DEAC0DC0D0",
		},
	}
	for _, tc := range tt {
		input := []byte{}
		for i := 0; i < tc.repeat; i++ {
			input = append(input, tc.input...)
		}
		got := strings.ToUpper(hex.EncodeToString(Keccak256(input)))
		if got != tc.want {
			t.Errorf("%s, got %q, want %q", tc.desc, got, tc.want)
		}
	}
}

// decodeHex converts an hex-encoded string into a raw byte string.
func decodeHex(s string) []byte {
	b, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return b
}
