package external

import (
	"fmt"
	"strconv"
)

const ItemType = "item"

type Item struct {
	Id          string      `json:"id"`
	Name        string      `json:"name"`
	Payment     Payment     `json:"payment"`
	Products    []IdType    `json:"products"`
	Campaigns   []IdType    `json:"campaigns"`
	CreatedTime Iso8601Time `json:"createdTime"`
	Type        string      `json:"type"`
}

type Items []Item

func (self Items) GetReferencedProductIds() []string {
	productIds := []string{}
	for _, item := range self {
		productIds = append(productIds, IdTypes(item.Products).Ids()...)
	}
	return productIds
}

type Payment struct {
	Quantity            int                        `json:"quantity"`
	Expires             int64                      `json:"expires"`
	Pricing             PaymentPricing             `json:"pricing"`
	PricingCNY          PaymentPricing             `json:"-"`
	PricingUSD          PaymentPricing             `json:"-"`
	AffiliateProductIds PaymentAffiliateProductIds `json:"affiliateProductIds"`
}

type PaymentPricing struct {
	CurrencyCode   string  `json:"currencyCode"`
	CurrencySymbol string  `json:"currencySymbol"`
	Price          Price   `json:"price"`
	UnitPeriod     int64   `json:"unitPeriod"`
	UnitPrice      Price   `json:"unitPrice"`
	DiscountPrice  Price   `json:"discountPrice"`
	Discount       float64 `json:"discount"`
}

type PaymentAffiliateProductIds struct {
	Alipay   PaymentAffiliate `json:"alipay"`
	AppStore PaymentAffiliate `json:"appstore"`
	Wechat   PaymentAffiliate `json:"wechat"`
}

type PaymentAffiliate struct {
	AutoRenewable string `json:"autoRenewable"`
	NonRenewing   string `json:"nonRenewing"`
	Consumable    string `json:"consumable"`
	NonConsumable string `json:"nonConsumable"`
}

type Price float64

func (p Price) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%.2f", p)), nil
}

func (p *Price) UnmarshalJSON(data []byte) error {
	val, err := strconv.ParseFloat(string(data), 64)
	if err != nil {
		return err
	}
	*p = Price(val)

	return nil
}
