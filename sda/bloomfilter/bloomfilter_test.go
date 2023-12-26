package bloomfilter

import (
	"hash"
	"hash/fnv"
	"testing"

	"github.com/spaolacci/murmur3"
)

var testData = []string{
	"YysUAY1DxCPajnnsZj2R",
	"W3iohUanmmOXcBSX9bR3",
	"9zItDdxuIrP5xU56hviq",
	"TVUh0imbSOLULaaLn5dJ",
	"dPwkf8ACBT9vmYnT3jTV",
	"q7MwoJCUiNyo245ZZssJ",
	"z4lIazANeynWbLrNnN8G",
	"CGVuEkxyFfwVjt7mR0oZ",
	"23tGajdiGdq9DAcut5FB",
	"kSFTEwaLclghNK8Nk2da",
	"ySDXqzmMZRTkSdpoyGoI",
	"KHouH0NfVYkSQJUMHmQv",
	"OKZwQMeyVIbBh2C6bFQl",
	"YIGBQomDPtsjOUuC9s8b",
	"qHj9FXePdtGoDSDNpLC5",
	"3aGF4kLGFsIBwqqjJ7or",
	"OtoduV6HhCsIZh3B33cT",
	"mHXUjEEdVw3ZabFjtoDt",
	"vOpFmSRynaCtJKGtoGPL",
	"5EM5h9gdMaWMCRxBGpy8",
	"73zZlES4n6AgusUSk2vt",
	"pA0eem4XPzqjOioaPWyg",
	"6rv155VnUDA5IXcLCvyi",
	"9XkN4jbzRTKlzoYo6dEN",
	"1PTcqPPxh2cE1hj6FbFy",
	"MZRd8htd6OwSGIQPypve",
	"P8obqWR68Lr6vcP6DDT1",
	"Q5ph80ucrznOExG6Arpl",
	"tlloJc4Bw7AM3ZiekCJs",
	"29mlVFGHIJ7hgkc87stK",
	"nQuws60gwEzV9yMvDQcI",
	"oEAEXNKSDIfM2E7XEAWD",
	"caacu3Hd5EgSUAAYYvB1",
	"tifnSMzUKqJcUimPexjs",
	"NxxQxZ2GhvqNddwgtWIV",
	"X6ALicVHXlsYjSOtXfSt",
	"pEySRoe2x0Jqcety0NBh",
	"QeZgENC4z3RiHtbUaQq2",
	"OJn6nzdtmFvsRxC0eEWj",
	"eTZ0dudnXgAFU0NlIkoa",
	"0sBovuJN5Dz56kc4jMia",
	"0SPs8OzMvvRJG6ZUsqCZ",
	"YM8c7Uygk4YXN49wWBpK",
	"hNtZsNWB9Xf6vAlDEFWG",
	"lYNcsxvfih3C7exHa3K9",
	"RE4ZB8si5qIe9MBei41V",
	"34Aeq3Lvwddeqbh35OcF",
	"RrQ7VyGdXemazQYGpMo7",
	"cImZS28oasrk6BEpWyl1",
	"UfsIj4YGjXpljkJGFukR",
	"68atEEeb2uLbvTz9a0Sx",
	"dIZhpeMH2Hya4HblPz6x",
	"U1IM3jwxdEKup1O9tK8P",
	"rDcEu2fIt8DdZV8CbJk4",
	"YYo7s3bvPRf2fOvwwSqy",
	"pqKujlZEscQY0qvtsVau",
	"Z3U7CNcwzG49qvLbvN9U",
	"hBwPbWfr4zzwjLQamkaL",
	"URrq2GYkOGxoB0UGQ20k",
	"EjbCMcaYWyYxXIbhg5V1",
	"KEnTpibui5PTB2tzfIHk",
	"ygPAJZyNBBag5GGVXOXS",
	"dVFmi8jr6UZjvtBm5puV",
	"KuOT4xcgOCR2HHX6rY6o",
	"0n9E35snuNxvPzbEwvi8",
	"1IMB1FHFWfddfD0UGoFX",
	"H3Eg6A0RFKCtbUElAK2w",
	"6rRBIV2njJfm0783mUum",
	"IJR67QrsplaauntdmAYi",
	"OrNGr6cddN2LpqHqYPRh",
	"2qxBzwGfwb8Z1RX9co80",
	"m0HGCJRQ6yB3wt63oLUN",
	"fm5uV53mBmwpIQhjOVi7",
	"LZMaBS2esvzyvXBCBPZx",
	"f3OkIbJvWHBehNuxFdbx",
	"smshsRa95zIo0yGwrMQe",
	"lureUaOYLCosi6knwWaW",
	"s2ynqtCLzMvKBg6TRJcA",
	"qDPzJlsrfb9A8Wl5BmOL",
	"nqUKmLfxgfaJOSyBNl07",
	"iO2W8uFXiIkAiCxI6gIe",
	"CBlRpPTbPs0U1rYIUSF0",
	"IEnlrY3Ch3iEmlIV7mho",
	"RrC6Zd7TYI7Zk0gvAYDt",
	"PtgfzqSDlbIZBAjiNXEG",
	"EzvW7XjHmZ4iEHdroVqx",
	"TYuw8o16cQYGwv2YIrcc",
	"P0K5xmbmsAu3BAhBYB17",
	"OkhKaJTQ43fQ7Ws9mJXT",
	"r2d4RT9MjlxtNY79T0eU",
	"JmEEU8UQ788LYwpscoAx",
	"kO5w4biWLrzE2HALUo1o",
	"kujlAULZHcVJEubt5zjm",
	"ZgC952WuxrOI9BbbybeS",
	"yhUY5dNzJLjjyIMwt716",
	"j5FsWVEqYdD0ZHvCELwf",
	"rzRlBzjJOFWy5JsqNMvE",
	"r2CvGqVgGMiXgLwiGE14",
	"Lvmm19SsbhLCi0JZqQwW",
	"mLbphCBVtaIJKCODPzrR",
}

func TestBloomFilter(t *testing.T) {
	bf := &BloomFilter{
		Bytes: make([]byte, 10000),
		Hashes: []hash.Hash64{
			fnv.New64(),
			murmur3.New64(),
		},
	}

	for _, str := range testData {
		bf.Push([]byte(str))
	}

	for _, str := range testData {
		t.Log(bf.Existed([]byte(str)))
	}

	t.Log(bf.Existed([]byte("mLbphCBVtaIJKCODPzRR")))
}
