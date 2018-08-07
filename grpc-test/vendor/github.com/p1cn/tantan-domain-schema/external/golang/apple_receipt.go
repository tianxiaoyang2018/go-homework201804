package external

import (
	"time"
)

var (
	AppleReceiptTypeConsumable    = "consumable"
	AppleReceiptTypeNonConsumable = "non-consumable"
	AppleReceiptTypeSubscription  = "non-renewing"
	AppleReceiptTypeAutoRenewable = "auto-renewable"
)

type AppleReceipt struct {
	ID                    string
	TransactionID         string
	OriginalTransactionID string
	ProductID             string
	Type                  string
	Quantity              string
	ExpireTime            *time.Time
	CancellationTime      *time.Time
	PurchaseTime          time.Time
	OriginalPurchaseTime  time.Time
	Receipt               string
}
