/*

 */

package userdata

import (
	"strings"
	"time"

	"github.com/bitspawngg/bitspawn-api/models"
	"github.com/dgryski/dgoogauth"
	"github.com/sirupsen/logrus"
)

// TODO: remove DB, log if they are not needed
// TODO: removing all commented codes
type UserService struct {
	DB                 *models.DB
	log                *logrus.Entry
	marketplaceHandler *models.MarketplaceHandler
}

func NewUserService(db *models.DB, log *logrus.Logger) *UserService {
	return &UserService{
		DB:                 db,
		log:                log.WithField("service", "userdata"),
		marketplaceHandler: models.NewMarketplaceHandler(db),
	}
}

type UserAccountOutPut struct {
	Username      string     `json:"username"`
	SteamId       string     `json:"steamId"`
	PaypalId      string     `json:"paypalId"`
	FavouriteGame []string   `json:"favouriteGame"`
	DisplayName   string     `json:"displayName"`
	FirstName     string     `json:"firstName"`
	LastName      string     `json:"lastName"`
	Country       string     `json:"country"`
	Timezone      string     `json:"timezone"`
	IsActive      bool       `json:"isActive"`
	AvatarUrl     string     `json:"avatarUrl"`
	ProfileBanner string     `json:"profileBanner"`
	CreatedAt     time.Time  `json:"createdAt"`
	UpdatedAt     time.Time  `json:"updatedAt"`
	OnlineTime    time.Time  `json:"onlineTime"`
	DeletedAt     *time.Time `json:"deletedAt"`
	EthGiftedAt   time.Time  `json:"ethGiftedAt"`
	PublicAddress string     `json:"publicAddress"`
	Sub           string     `json:"sub"`
}

func (usvc *UserService) Compute2FACode(sub string) int {
	// Refer to: totp.danhersam.com/
	inputBase32 := strings.NewReplacer("-", "", "0", "O", "1", "I", "8", "B", "9", "6").Replace(sub)
	inputNoSpacesUpper := strings.ToUpper(inputBase32)
	//TODO: put 600 into config
	return dgoogauth.ComputeCode(inputNoSpacesUpper, time.Now().Unix()/600)
}

// GetMarketplaceOrderRecords is returning the list of marketplace order records
func (usvc *UserService) GetMarketplaceOrderRecords(publicAddress string, page, per_page int) ([]*models.MarketplaceOrderRecord, error) {
	return usvc.marketplaceHandler.FetchMarketplaceRecords(publicAddress, page, per_page)
}

func (usvc *UserService) CountMyMarketplaceRecordPages(publicAddress string, per_page int) (int64, error) {
	return usvc.marketplaceHandler.CountMyMarketplaceRecordPages(publicAddress, per_page)
}
