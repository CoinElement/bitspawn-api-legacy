/*

 */

package models

import (
	"time"
)

type AdminConfig struct {
	ID           uint `gorm:"primary_key"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	RpcNodeUrl   string `json:"rpc_node_url"`
	FaucetKey    string `json:"faucet_key"`
	SpwnContract string `json:"spwn_contract"`
	SpwnFee      string `json:"spwn_fee"`
	CredContract string `json:"cred_contract"`
	CredFee      string `json:"cred_fee"`
	UsdcContract string `json:"usdc_contract"`
	UsdcFee      string `json:"usdc_fee"`
	WalletUrl    string `json:"wallet_url"`
	AdapterUrl   string `json:"adapter_url"`
	OrganizerUrl string `json:"organizer_url"`
	AdminUserID  string `json:"admin_user_id"`
	APIRate      uint   `json:"api_rate"`
}

func ReadAdminConfig(db *DB) (*AdminConfig, error) {
	config := AdminConfig{}
	err := db.Table("admin_configs").First(&config).Error
	if err != nil {
		return nil, err
	}
	return &config, nil
}
