/*

 */

package models

// UserAccount maps the users table in the database. It stores the username
// password and the name of the user.
type Badge struct {
	BadgeId  string `json:"badgeId" gorm:"type:text;primary_key"`
	BadgeUrl string `json:"badgeUrl"`
}

func (db *DB) StorageBadgeUrl(badge_url, badge_id string) (*Badge, error) {

	badge := Badge{
		BadgeId:  badge_id,
		BadgeUrl: badge_url,
	}

	db.log.Debug(badge)
	return nil, db.Save(&badge).Error
}

func (db *DB) GetAllBadgeUrl() ([]Badge, error) {
	badge := []Badge{}

	err := db.Find(&badge).Error
	if err != nil {
		db.log.Error(err)
		return nil, err
	}

	db.log.Debug(badge)
	return badge, nil
}

func (db *DB) GetMatchBadgeUrl(badgeId string) (*Badge, error) {
	badge := Badge{}

	err := db.Where(&Badge{BadgeId: badgeId}).First(&badge).Error
	if err != nil {
		db.log.Error(err)
		return nil, err
	}

	db.log.Debug(badge)
	return &badge, nil
}
