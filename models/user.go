/*

 */

package models

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"time"

	"gorm.io/datatypes"
)

// UserAccount maps the users table in the database. It stores the username
// password and the name of the user.

type Battlenet struct {
	Battletag string `json:"battletag"`
	Region    string `json:"region"`
}

type Epic struct {
	EpicId string `json:"epicId"`
}

type Steam struct {
	SteamId string `json:"steamId"`
}

type Paypal struct {
	PaypalId string `json:"paypalId"`
}

type GameAccount struct {
	Battlenet Battlenet `json:"battlenet"`
	Epic      Epic      `json:"epic"`
	Steam     Steam     `json:"steam"`
}

type PaymentAccount struct {
	Paypal Paypal `json:"paypal"`
}

type Favourite struct {
	FavouriteGame []string `json:"favouriteGame"`
	// FavouriteTeam []models.Club `json:"favouriteTeam"`
}

func (ga GameAccount) Value() (driver.Value, error) {
	valueString, err := json.Marshal(ga)
	return string(valueString), err
}

func (ga *GameAccount) Scan(value interface{}) error {
	if err := json.Unmarshal(value.([]byte), &ga); err != nil {
		return err
	}
	return nil
}

func (pa PaymentAccount) Value() (driver.Value, error) {
	valueString, err := json.Marshal(pa)
	return string(valueString), err
}

func (pa *PaymentAccount) Scan(value interface{}) error {
	if err := json.Unmarshal(value.([]byte), &pa); err != nil {
		return err
	}
	return nil
}

func (favt Favourite) Value() (driver.Value, error) {
	valueString, err := json.Marshal(favt)
	return string(valueString), err
}

func (favt *Favourite) Scan(value interface{}) error {
	if err := json.Unmarshal(value.([]byte), &favt); err != nil {
		return err
	}
	return nil
}

// type JSONB map[string]interface{}

// func (j JSONB) Value() (driver.Value, error) {
// 	valueString, err := json.Marshal(j)
// 	return string(valueString), err
// }

// func (j *JSONB) Scan(value interface{}) error {
// 	if err := json.Unmarshal(value.([]byte), &j); err != nil {
// 		return err
// 	}
// 	return nil
// }

type UserAccount struct {
	Sub           string         `json:"sub" gorm:"type:text;primary_key"`
	Username      string         `json:"username,omitempty" gorm:"type:text"`
	Favourite     datatypes.JSON `json:"favourite" sql:"type:jsonb"`
	DisplayName   string         `json:"displayName" gorm:"unique;not null"`
	IsActive      bool           `json:"isActive"`
	AvatarUrl     string         `json:"avatarUrl"`
	ProfileBanner string         `json:"profileBanner"`
	CreatedAt     time.Time      `json:"createdAt"`
	UpdatedAt     time.Time      `json:"updatedAt"`
	OnlineTime    time.Time      `json:"onlineTime"`
	DeletedAt     *time.Time     `json:"-" sql:"index"`
	EthGiftedAt   time.Time      `json:"-"`
	PublicAddress string         `json:"publicAddress"`
	PhoneNumber   string         `json:"phoneNumber,omitempty"`
	Enabled2FA    bool           `json:"enabled2FA"`
}

type UserGameAccount struct {
	Sub         string      `json:"sub" gorm:"type:text;primary_key"`
	GameAccount GameAccount `json:"gameAccount" gorm:"type:jsonb"`
}

type UserPaymentAccount struct {
	Sub            string         `json:"sub" gorm:"type:text;primary_key"`
	PaymentAccount PaymentAccount `json:"paymentAccount" gorm:"type:jsonb"`
}

type UserFeed struct {
	ID        uint       `gorm:"primary_key"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updateAt"`
	DeletedAt *time.Time `sql:"index" json:"deletedAt"`
	Sub       string     `json:"sub"`
	Icon      string     `json:"icon"`
	Link      string     `json:"link"`
	Message   string     `json:"message"`
}

func (db *DB) UserSignup(username, displayName, publicAddress, sub string, favourite json.RawMessage) error {
	u := UserAccount{
		Username:      username,
		DisplayName:   displayName,
		IsActive:      true,
		Sub:           sub,
		PublicAddress: publicAddress,
		EthGiftedAt:   time.Now().UTC(),
		Favourite:     datatypes.JSON(favourite),
	}

	g := UserGameAccount{
		Sub: sub,
	}

	p := UserPaymentAccount{
		Sub: sub,
	}

	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Save(&u).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Save(&g).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Save(&p).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (db *DB) GetUser(sub string) (*UserAccount, error) {
	u := UserAccount{}
	err := db.Where(&UserAccount{Sub: sub}).First(&u).Error
	if err != nil {
		db.log.Error(err)
		return nil, err
	}
	return &u, nil
}

func (db *DB) GetUserGameAccount(sub string) (*GameAccount, error) {
	u := UserGameAccount{}
	err := db.Where(UserGameAccount{Sub: sub}).First(&u).Error
	if err != nil {
		db.log.Error(err)
		return nil, err
	}
	return &u.GameAccount, nil
}

func (db *DB) FetchExistUserNumber() (int64, error) {
	var count int64

	err := db.Table("user_accounts").Count(&count).Error
	if err != nil {
		db.log.Error(err)
		return 0, err
	}

	count = int64(count)

	return count, nil
}

func (db *DB) GetUserProfileByDisplayName(displayname string) (*UserAccount, error) {
	u := UserAccount{}
	err := db.Where(&UserAccount{DisplayName: displayname}).First(&u).Error
	if err != nil {
		db.log.Error(err)
		return nil, err
	}
	return &u, nil
}

func (db *DB) GetUserSubByUserName(username string) (*UserAccount, error) {
	u := UserAccount{}

	err := db.Select("sub").Where(&UserAccount{Username: username}).First(&u).Error
	if err != nil {
		db.log.Error(err)
		return nil, err
	}

	return &u, nil
}

func (db *DB) CountDisplayNameUsage(displayname string) (int64, error) {

	var count int64
	err := db.Table("user_accounts").
		Where(&UserAccount{DisplayName: displayname}).
		Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (db *DB) GetUserProfile(username string) (*UserAccount, error) {
	u := UserAccount{}
	err := db.Where(&UserAccount{Username: username}).First(&u).Error
	if err != nil {
		db.log.Error(err)
		return nil, err
	}
	return &u, nil
}

func (db *DB) UpdateUserProfile(sub string, displayName string) (*UserAccount, error) {
	u, err := db.GetUser(sub)
	if err != nil {
		return nil, err
	}
	return nil, db.Model(u).Updates(map[string]interface{}{"display_name": displayName}).Error
}

func (db *DB) UpdateUserPhoneNumber(sub string, phoneNumber string) error {
	u, err := db.GetUser(sub)
	if err != nil {
		return err
	}
	return db.Model(u).Updates(UserAccount{PhoneNumber: phoneNumber}).Error
}

func (db *DB) Disable2FA(sub string) error {
	u, err := db.GetUser(sub)
	if err != nil {
		return err
	}
	return db.Model(u).Update("enabled2_fa", false).Error
}

func (db *DB) Enable2FA(sub string) error {
	u, err := db.GetUser(sub)
	if err != nil {
		return err
	}
	return db.Model(u).Updates(UserAccount{Enabled2FA: true}).Error
}

func (db *DB) UpdateOnlineTime(sub string) (*UserAccount, error) {
	u, err := db.GetUser(sub)
	if err != nil {
		return nil, err
	}
	time := time.Now().UTC()
	return u, db.Model(u).Updates(map[string]interface{}{"online_time": time}).Error
}

func (db *DB) UpdateUserAvatar(sub string, avatar_url string) (*UserAccount, error) {
	u, err := db.GetUser(sub)
	if err != nil {
		return nil, err
	}
	return nil, db.Model(u).Updates(map[string]interface{}{"avatar_url": avatar_url}).Error
}

func (db *DB) UpdateUserProfileBanner(sub string, profile_banner string) (*UserAccount, error) {
	u, err := db.GetUser(sub)
	if err != nil {
		return nil, err
	}
	return nil, db.Model(u).Updates(map[string]interface{}{"profile_banner": profile_banner}).Error
}

func (db *DB) UpdateEthGiftedAt(sub string) error {
	u, err := db.GetUser(sub)
	if err != nil {
		return err
	}
	return db.Model(u).Update("eth_gifted_at", time.Now().UTC()).Error
}

func (db *DB) UpdateUserFavouriteData(sub string, favouriteData json.RawMessage) error {
	// return db.Exec("UPDATE public.user_accounts SET favourite = jsonb_set(favourite, $1, $2) WHERE username = $3", favouriteField, favouriteValue, username).Error
	u, err := db.GetUser(sub)
	if err != nil {
		return err
	}
	return db.Model(u).Update("favourite", datatypes.JSON(favouriteData)).Error
}

func (db *DB) ListFriendByDisplayName(friendName string) ([]UserAccount, error) {
	friendSearchList := []UserAccount{}
	var rows *sql.Rows
	var err error
	rows, err = db.Table("user_accounts").
		Select("user_accounts.sub,user_accounts.display_name,user_accounts.avatar_url").
		Where("display_name LIKE (?)", friendName).
		Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var userInfo UserAccount
		_ = db.ScanRows(rows, &userInfo)
		friendSearchList = append(friendSearchList, userInfo)
	}
	return friendSearchList, nil
}

func (db *DB) CreateUserFeed(feed *UserFeed) error {
	return db.Create(&feed).Error
}

func (db *DB) ReadUserFeed(sub string) ([]UserFeed, error) {
	feeds := []UserFeed{}
	err := db.Where(&UserFeed{Sub: sub}).Find(&feeds).Error
	if err != nil {
		return nil, err
	}
	return feeds, nil
}
