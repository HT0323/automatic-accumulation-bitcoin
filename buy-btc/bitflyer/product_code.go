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
		return "BTC_JPY"
	case Ethbtc:
		return "BTC_JPY"
	case bchbtc:
		return "BTC_JPY"
	default:
		return "BTC_JPY"
	}
}
