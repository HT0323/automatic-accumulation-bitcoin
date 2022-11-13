package bitflyer

type ProductCode int

const (
	Btcjpy ProductCode = iota
	Ethjpy
	Fxbtcjpy
	Ethbtc
	bchbtc
)

func (code ProductCode) String() string {
	switch code {
	case Btcjpy:
		return "BTC_JPY"
	case Ethjpy:
		return "ETH_JPY"
	case Fxbtcjpy:
		return "FX_BTC_JPY"
	case Ethbtc:
		return "ETH_BTC"
	case bchbtc:
		return "BCH_BTC"
	default:
		return "BTC_JPY"
	}
}

type OrderType int

const (
	Limit  OrderType = iota //指値
	Market                  //成行
)

func (or OrderType) String() string {
	switch or {
	case Limit:
		return "LIMIT"
	case Market:
		return "MARKET"
	default:
		return "LIMIT"
	}
}

type Side int

const (
	Buy  Side = iota //買い
	Sell             //売り
)

func (s Side) String() string {
	switch s {
	case Buy:
		return "BUY"
	case Sell:
		return "SELL"
	default:
		return "BUY"
	}
}

type TimeInForce int

const (
	Gtc TimeInForce = iota
	Ioc
	Fok
)

func (tif TimeInForce) String() string {
	switch tif {
	case Gtc:
		return "GTC"
	case Ioc:
		return "IOC"
	case Fok:
		return "FOK"
	default:
		return "GTC"
	}
}
