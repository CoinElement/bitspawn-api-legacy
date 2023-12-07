package models

import (
	"math"
	"math/big"
)

// MarketplaceHandler is used to handle marketplace-order-record data handler in database
type MarketplaceHandler struct {
	*DB
}

// NewMarketplaceHandler is used to create a new MarketplaceHandler
func NewMarketplaceHandler(db *DB) *MarketplaceHandler {
	return &MarketplaceHandler{db}
}

// MarketplaceOrderRecord copied from deposit-manager
// TODO: this should be done in deposit-manager
type MarketplaceOrderRecord struct {
	OrderId                string   `json:"orderId"` // v4 uuid that refers to the orderId
	PoaAddr                string   `json:"poaAddr"`
	TokenAmount            *big.Int `json:"tokenAmount"`            // token amount that got burnt, either credit or platform, in min unit
	TokenType              string   `json:"tokenType"`              // either credit or spwn
	PurchaseCurrency       string   `json:"purchaseCurrency"`       // USD
	PurchaseCurrencyAmount float64  `json:"purchaseCurrencyAmount"` // in purchaseCurrency unit, e.g: 10.00
	ProductId              string   `json:"productId"`
	Country                string   `json:"country"`
	NativeCurrencyPrice    float64  `json:"nativeCurrencyPrice"`
	NativeCurrencyName     string   `json:"nativeCurrencyName"`
	BrandName              string   `json:"brandName"`
	Statue                 string   `json:"statue"`
	ProductDetail          string   `json:"productDetail"`
}

// FetchTeamMachineByID fetchs records based on the given user
func (mh *MarketplaceHandler) FetchMarketplaceRecords(publicAddress string, page, per_page int) ([]*MarketplaceOrderRecord, error) {
	marketRecords := []*MarketplaceOrderRecord{}

	offset := (page - 1) * per_page
	limit := per_page

	if err := mh.Where(&MarketplaceOrderRecord{PoaAddr: publicAddress}).
		Order("insert_seq desc").
		Limit(limit).
		Offset(offset).
		Find(&marketRecords).Error; err != nil {
		return nil, err
	}
	return marketRecords, nil
}

func (mh *MarketplaceHandler) CountMyMarketplaceRecordPages(publicAddress string, per_page int) (int64, error) {
	var count int64

	err := mh.Table("marketplace_order_records").
		Where("poaAddr = ?", publicAddress).
		Count(&count).Error

	if err != nil {
		return 0, err
	}
	pagefloat := math.Ceil(float64(count) / float64(per_page))
	page := int64(pagefloat)

	return page, nil
}
