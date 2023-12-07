package models

import (
	"math"
	"time"
)

// DepositRecord maps the users deposit record table in the database.
type DepositRecord struct {
	TxHash     string    `json:"txHash" gorm:"primary_key;"`
	TxID       string    `json:"txId"`
	TxSender   string    `json:"txSender"`
	TxReceiver string    `json:"txReceiver"`
	PoaAddress string    `json:"poaAddress"`
	CoinType   string    `json:"coinType"`
	CoinAmount string    `json:"coinAmount" gorm:"type:bigint"`
	SpwnAmount string    `json:"spwnAmount"`
	MintDate   time.Time `json:"mintDate"`
	Action     string    `json:"action"`
	TxStatus   string    `json:"txStatus"`
	Fee        string    `json:"fee"`
	Remark     string    `json:"remark"`
}

func (db *DB) GetUserDepositRecord(publicAddress string, page, per_page int) ([]DepositRecord, error) {
	depositRecord := []DepositRecord{}

	offset := (page - 1) * per_page
	limit := per_page

	err := db.Where(&DepositRecord{PoaAddress: publicAddress}).
		Order("mint_Date desc").
		Limit(limit).
		Offset(offset).
		Find(&depositRecord).Error
	if err != nil {
		return nil, err
	}

	return depositRecord, nil
}

func (db *DB) CountMyDepositRecordPages(publicAddress string, per_page int) (int64, error) {
	var count int64

	err := db.Table("deposit_records").
		Where("poa_address = ?", publicAddress).
		Count(&count).Error

	if err != nil {
		return 0, err
	}
	pagefloat := math.Ceil(float64(count) / float64(per_page))
	page := int64(pagefloat)

	return page, nil
}
