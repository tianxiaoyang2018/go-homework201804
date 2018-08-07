package currency

const (
	CNY = "CNY"
	USD = "USD"
	JPY = "JPY"
	EUR = "EUR"
	KRW = "KRW"
	TWD = "TWD"
	HKD = "HKD"
)

var symbols map[string]string = map[string]string{
	"CNY": "¥",
	"USD": "$",
	"JPY": "¥",
	"EUR": "€",
	"KRW": "₩",
	"TWD": "NT$",
	"HKD": "$",
}

func Symbol(code string) string {
	return symbols[code]
}
