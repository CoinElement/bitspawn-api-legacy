package config

import (
	"time"

	"github.com/bitspawngg/bitspawn-api/models"
	"github.com/sirupsen/logrus"
)

type Service struct {
	DB  *models.DB
	log *logrus.Logger

	conf *models.AdminConfig
	quit chan int
}

func NewService(db *models.DB, logger *logrus.Logger) *Service {
	s := &Service{
		DB:  db,
		log: logger,
	}

	s.updateConf()

	ticker := time.NewTicker(5 * time.Minute)
	quit := make(chan int)
	go func() {
		for {
			select {
			case <-ticker.C:
				s.updateConf()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
	s.quit = quit
	return s
}

func (s *Service) updateConf() {
	adminConfig, err := models.ReadAdminConfig(s.DB)

	if err != nil {
		s.log.Error("failed to fetch latest config")
	}
	s.conf = adminConfig
}

func (s *Service) GetConfig() *models.AdminConfig {
	return s.conf
}

func (s *Service) Close() {
	s.quit <- 0
}
