package ethNotification

type EngineTokenDataSource interface {
	FindTokens(tokenAddress []string) []TokenContract
}

type DefaultTokenDataSource struct {
	Data map[string]TokenContract
}

func (ds DefaultTokenDataSource) FindTokens(tokenAddress []string) []TokenContract {
	result := []TokenContract{}

	for _, addr := range tokenAddress {
		tc, found := ds.Data[addr]
		if found {
			result = append(result, tc)
		}
	}

	return result
}

func newDefaultTokenDataSource() DefaultTokenDataSource {
	return DefaultTokenDataSource{
		Data: map[string]TokenContract{
			"0x86fa049857e0209aa7d9e616f7eb3b3b78ecfdb0": TokenContract{
				Name:     "EOS",
				Symbol:   "EOS",
				Decimals: 18,
			},
			"0xa974c709cfb4566686553a20790685a47aceaa33": TokenContract{
				Name:     "Mixin",
				Symbol:   "XIN",
				Decimals: 18,
			},
			"0xf230b790e05390fc8295f4d3f60332c93bed42e2": TokenContract{
				Name:     "Tronix",
				Symbol:   "TRX",
				Decimals: 6,
			},
			"0xb8c77482e45f1f44de1745f52c74426c631bdd52": TokenContract{
				Name:     "Binance Coin",
				Symbol:   "BNB",
				Decimals: 18,
			},
			"0xdd974d5c2e2928dea5f71b9825b8b646686bd200": TokenContract{
				Name:     "Kyber Network",
				Symbol:   "KNC",
				Decimals: 18,
			},
			"0xd26114cd6ee289accf82350c8d8487fedb8a0c07": TokenContract{
				Name:     "OmiseGO",
				Symbol:   "OMG",
				Decimals: 18,
			},
			"0xb63b606ac810a52cca15e44bb630fd42d8d1d83d": TokenContract{
				Name:     "Monaco",
				Symbol:   "MCO",
				Decimals: 8,
			},
			"0xd850942ef8811f2a866692a623011bde52a462c1": TokenContract{
				Name:     "VeChain",
				Symbol:   "VEN",
				Decimals: 18,
			},
			"0x08d32b0da63e2C3bcF8019c9c5d849d7a9d791e6": TokenContract{
				Name:     "Dentacoin",
				Symbol:   "DCN",
				Decimals: 0,
			},
			"0xe41d2489571d322189246dafa5ebde1f4699f498": TokenContract{
				Name:     "ZRX",
				Symbol:   "ZRX",
				Decimals: 18,
			},
			"0x818fc6c2ec5986bc6e2cbf00939d90556ab12ce5": TokenContract{
				Name:     "Kin",
				Symbol:   "KIN",
				Decimals: 18,
			},
			"0xb7cb1c96db6b22b0d3d9536e0108d062bd488f74": TokenContract{
				Name:     "Walton",
				Symbol:   "WTC",
				Decimals: 18,
			},
			"0xb5a5f22694352c15b00323844ad545abb2b11028": TokenContract{
				Name:     "ICON",
				Symbol:   "ICX",
				Decimals: 18,
			},
			"0xf85feea2fdd81d51177f6b8f35f0e6734ce45f5f": TokenContract{
				Name:     "CyberMiles",
				Symbol:   "CMT",
				Decimals: 18,
			},
			"0x05f4a42e251f2d52b8ed15e9fedaacfcef1fad27": TokenContract{
				Name:     "Zilliqa",
				Symbol:   "ZIL",
				Decimals: 12,
			},
			"0x1f573d6fb3f13d689ff844b4ce37794d79a7ff1c": TokenContract{
				Name:     "Bancor",
				Symbol:   "BNT",
				Decimals: 18,
			},
			"0x5ca9a71b1d01849c0a95490cc00559717fcf0d1d": TokenContract{
				Name:     "Aeternity",
				Symbol:   "AE",
				Decimals: 18,
			},
			"0xf0ee6b27b759c9893ce4f094b49ad28fd15a23e4": TokenContract{
				Name:     "Enigma",
				Symbol:   "ENG",
				Decimals: 8,
			},
			"0x9f8f72aa9304c8b593d555f12ef6589cc3a579a2": TokenContract{
				Name:     "Maker",
				Symbol:   "MKR",
				Decimals: 18,
			},
			"0x0d8775f648430679a709e98d2b0cb6250d2887ef": TokenContract{
				Name:     "BAT",
				Symbol:   "BAT",
				Decimals: 8,
			},
			"0xcb97e65f07da24d46bcdd078ebebd7c6e6e3d750": TokenContract{
				Name:     "Bytom",
				Symbol:   "BTM",
				Decimals: 8,
			},
			"0xa15c7ebe1f07caf6bff097d8a589fb8ac49ae5b3": TokenContract{
				Name:     "Pundi X Token",
				Symbol:   "NPXS",
				Decimals: 18,
			},
			"0xfa1a856cfa3409cfa145fa4e20eb270df3eb21ab": TokenContract{
				Name:     "IOStoken",
				Symbol:   "IOST",
				Decimals: 18,
			},
			"0x4ceda7906a5ed2179785cd3a40a69ee8bc99c466": TokenContract{
				Name:     "Aion",
				Symbol:   "AION",
				Decimals: 8,
			},
			"0x1985365e9f78359a9b6ad760e32412f4a445e862": TokenContract{
				Name:     "Reputation",
				Symbol:   "REP",
				Decimals: 18,
			},
			"0xef68e7c694f40c8202821edf525de3782458639f": TokenContract{
				Name:     "Loopring",
				Symbol:   "LRC",
				Decimals: 18,
			},
			"0xa74476443119a942de498590fe1f2454d7d4ac0d": TokenContract{
				Name:     "Golem",
				Symbol:   "GNT",
				Decimals: 18,
			},
			"0xbf2179859fc6d5bee9bf9158632dc51678a4100e": TokenContract{
				Name:     "aelf",
				Symbol:   "ELF",
				Decimals: 18,
			},
			"0x5d65d971895edc438f465c17db6992698a52318d": TokenContract{
				Name:     "Nebulas",
				Symbol:   "NAS",
				Decimals: 18,
			},
			"0x744d70fdbe2ba4cf95131626614a1763df805b9e": TokenContract{
				Name:     "StatusNetwork",
				Symbol:   "SNT",
				Decimals: 18,
			},
			"0xcbce61316759d807c474441952ce41985bbc5a40": TokenContract{
				Name:     "MoacToken",
				Symbol:   "MOAC",
				Decimals: 18,
			},
			"0x595832f8fc6bf59c85c527fec3740a1b7a361269": TokenContract{
				Name:     "Power Ledger",
				Symbol:   "POWR",
				Decimals: 6,
			},
			"0xd4fa1460f537bb9085d22c7bccb5dd450ef28e3a": TokenContract{
				Name:     "Populous",
				Symbol:   "PPT",
				Decimals: 8,
			},
			"0x419d0d8bdd9af5e606ae2232ed285aff190e711b": TokenContract{
				Name:     "FunFair",
				Symbol:   "FUN",
				Decimals: 8,
			},
			"0x168296bb09e24a88805cb9c33356536b980d3fc5": TokenContract{
				Name:     "RHOC",
				Symbol:   "RHOC",
				Decimals: 8,
			},
			"0x90528aeb3a2b736b780fd1b6c478bb7e1d643170": TokenContract{
				Name:     "XPlay",
				Symbol:   "XPA",
				Decimals: 18,
			},
			"0x9992ec3cf6a55b00978cddf2b27bc6882d88d1ec": TokenContract{
				Name:     "Polymath",
				Symbol:   "POLY",
				Decimals: 18,
			},
			"0x39bb259f66e1c59d5abef88375979b4d20d98022": TokenContract{
				Name:     "WAX",
				Symbol:   "WAX",
				Decimals: 8,
			},
			"0x5af2be193a6abca9c8817001f45744777db30756": TokenContract{
				Name:     "Ethos",
				Symbol:   "ETHOS",
				Decimals: 8,
			},
			"0xb97048628db6b661d4c2aa833e95dbe1a905b280": TokenContract{
				Name:     "TenXPay",
				Symbol:   "PAY",
				Decimals: 18,
			},
			"0xb91318f35bdb262e9423bc7c7c2a3a93dd93c92c": TokenContract{
				Name:     "Nuls",
				Symbol:   "NULS",
				Decimals: 18,
			},
			"0x0f5d2fb29fb7d3cfee444a200298f468908cc942": TokenContract{
				Name:     "Decentraland",
				Symbol:   "MANA",
				Decimals: 18,
			},
			"0xc5bbae50781be1669306b9e001eff57a2957b09d": TokenContract{
				Name:     "Gifto",
				Symbol:   "GTO",
				Decimals: 5,
			},
			"0x8f3470a7388c05ee4e7af3d01d8c722b0ff52374": TokenContract{
				Name:     "Veritaseum",
				Symbol:   "VERI",
				Decimals: 18,
			},
			"0x62a56a4a2ef4d355d34d10fbf837e747504d38d4": TokenContract{
				Name:     "Paypex",
				Symbol:   "PAYX",
				Decimals: 2,
			},
			"0x12480e24eb5bec1a9d4369cab6a80cad3c0a377a": TokenContract{
				Name:     "Substratum",
				Symbol:   "SUB",
				Decimals: 2,
			},
			"0x3883f5e181fccaf8410fa61e12b59bad963fb645": TokenContract{
				Name:     "Theta Token",
				Symbol:   "THETA",
				Decimals: 18,
			},
			"0x618e75ac90b12c6049ba3b27f5d5f8651b0037f6": TokenContract{
				Name:     "QASH",
				Symbol:   "QASH",
				Decimals: 6,
			},
			"0xa4e8c3ec456107ea67d3075bf9e3df3a75823db0": TokenContract{
				Name:     "Loom",
				Symbol:   "LOOM",
				Decimals: 18,
			},
			"0x0e935e976a47342a4aee5e32ecf2e7b59195e82f": TokenContract{
				Name:     "BMB",
				Symbol:   "BMB",
				Decimals: 18,
			},
			"0x177d39ac676ed1c67a2b268ad7f1e58826e5b0af": TokenContract{
				Name:     "CoinDash Token",
				Symbol:   "CDT",
				Decimals: 18,
			},
			"0x0abdace70d3790235af448c88547603b945604ea": TokenContract{
				Name:     "district0x Network Token",
				Symbol:   "DNT",
				Decimals: 18,
			},
			"0x23ccc43365d9dd3882eab88f43d515208f832430": TokenContract{
				Name:     "MidasProtocol",
				Symbol:   "MAS",
				Decimals: 18,
			},
			"0xf8b358b3397a8ea5464f8cc753645d42e14b79ea": TokenContract{
				Name:     "Airbloc",
				Symbol:   "ABL",
				Decimals: 18,
			},
		},
	}
}
