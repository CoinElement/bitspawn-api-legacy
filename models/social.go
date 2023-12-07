package models

import (
	"errors"
	"gorm.io/gorm"
)

// SocialHandler is used to handle social data handler and maniupulate social data in database
type SocialHandler struct {
	*DB
}

// NewSocialHandler is used to create a new SocialHandler
func NewSocialHandler(db *DB) *SocialHandler {
	return &SocialHandler{db}
}

type SocialType string

const (
	Twitter SocialType = "TWITTER"
	Twitch  SocialType = "TWITCH"
	Discord SocialType = "DISCORD"
)

func (s SocialType) Valid() bool {
	if s != Twitter && s != Twitch && s != Discord {
		return false
	}
	return true
}

// SocialLink is used to save the related user's social medias
type SocialLink struct {
	UserID     string     `json:"userID" gorm:"primaryKey;"`
	SocialType SocialType `json:"socialType" gorm:"primaryKey;"`
	SocialID   string     `json:"socialID"`
}

// CreateSocialLink is used to create a social link in database
// it will replace it if there is a existing object in database
func (sh *SocialHandler) CreateSocialLink(sl *SocialLink) error {
	social, err := sh.FetchSocialLink(sl.UserID, sl.SocialType)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return convError(err)
	}
	if social != nil {
		return convError(sh.UpdateSocialLink(sl.UserID, sl.SocialID, sl.SocialType))
	}
	return convError(sh.Create(sl).Error)
}

// DeleteSocialLink is used to delete the given social link in database
func (sh *SocialHandler) DeleteSocialLink(userID string, socialType SocialType) error {
	sl := &SocialLink{}
	if err := sh.Where(&SocialLink{UserID: userID, SocialType: socialType}).First(&sl).Error; err != nil {
		return convError(err)
	}
	return convError(sh.Delete(sl).Error)
}

// UpdateSocialLink is used to edit the given social link in database
func (sh *SocialHandler) UpdateSocialLink(userID, socialID string, socialType SocialType) error {
	sl := &SocialLink{}
	if err := sh.Where(&SocialLink{UserID: userID, SocialType: socialType}).First(&sl).Error; err != nil {
		return convError(err)
	}
	return sh.Model(sl).Update("social_id", socialID).Error
}

// FetchSocialLink is used to fetch the given social link in database
func (sh *SocialHandler) FetchSocialLink(userID string, socialType SocialType) (*SocialLink, error) {
	sl := &SocialLink{UserID: userID, SocialType: socialType}
	if err := sh.Where(sl).First(&sl).Error; err != nil {
		return nil, convError(err)
	}
	return sl, nil
}

// FetchAllSocialLink is used to fetch all socials for the given user
func (sh *SocialHandler) FetchAllSocialLink(userID string) ([]*SocialLink, error) {
	sl := []*SocialLink{}
	if err := sh.Where("user_id=?", userID).Find(&sl).Error; err != nil {
		return nil, convError(err)
	}
	return sl, nil
}
