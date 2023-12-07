package social

import (
	"github.com/bitspawngg/bitspawn-api/models"
	"github.com/sirupsen/logrus"
)

// SocialService is used to handler social link related request services
type SocialService struct {
	sh  *models.SocialHandler
	log *logrus.Entry
}

// NewSocialService create a social link instance
func NewSocialService(db *models.DB, log *logrus.Entry) *SocialService {
	return &SocialService{
		sh:  models.NewSocialHandler(db),
		log: log.WithField("service", "social"),
	}
}

func (ss *SocialService) CreateSocialLink(socialLink *models.SocialLink) error {
	if err := ss.sh.CreateSocialLink(socialLink); err != nil {
		ss.log.Error("error in creating social link", err)
		return err
	}
	return nil
}
func (ss *SocialService) DeleteSocialLink(userID string, socialType models.SocialType) error {
	if err := ss.sh.DeleteSocialLink(userID, socialType); err != nil {
		ss.log.Error("error in deleting social link", err)
		return err
	}
	return nil
}

func (ss *SocialService) UpdateSocialLink(userID, socialID string, socialType models.SocialType) error {
	if err := ss.sh.UpdateSocialLink(userID, socialID, socialType); err != nil {
		ss.log.Error("error in updating social link", err)
		return err
	}
	return nil
}
func (ss *SocialService) GetSocialLink(userID string, socialType models.SocialType) (*models.SocialLink, error) {
	sl, err := ss.sh.FetchSocialLink(userID, socialType)
	if err != nil {
		ss.log.Error("error in fetching social link", err)
		return nil, err
	}
	return sl, nil
}

func (ss *SocialService) GetAllSocialLink(userID string) ([]*models.SocialLink, error) {
	sl, err := ss.sh.FetchAllSocialLink(userID)
	if err != nil {
		ss.log.Error("error in fetching social link", err)
		return nil, err
	}
	return sl, nil
}
