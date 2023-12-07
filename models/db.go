/*

 */

package models

import (
	"errors"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DB struct {
	gorm.DB
	dbType           string
	dbConnectionPath string

	log *logrus.Entry
}

// NewDB initializes a DB object
func NewDB(dpType, dbConnectionPath string, log *logrus.Logger) *DB {
	return &DB{
		dbType:           dpType,
		dbConnectionPath: dbConnectionPath,
		log:              log.WithField("package", "models"),
	}
}

// Connect initiates a new connection with the given connection parameters
// for the database. It returns an error in case the connection fails.
func (db *DB) Connect() error {
	var dbConnection *gorm.DB
	var err error
	if db.dbType == "sqlite3" {
		dbConnection, err = gorm.Open(sqlite.Open(db.dbConnectionPath), &gorm.Config{})
	} else if db.dbType == "postgres" {
		dbConnection, err = gorm.Open(postgres.Open(db.dbConnectionPath), &gorm.Config{})
	} else {
		err = errors.New("invalid dbtype")
	}
	if err != nil {
		return err
	}
	_ = dbConnection.AutoMigrate(&UserAccount{})
	_ = dbConnection.AutoMigrate(&Badge{})
	_ = dbConnection.AutoMigrate(&TournamentData{})
	_ = dbConnection.AutoMigrate(&ChallengeData{})
	_ = dbConnection.AutoMigrate(&PlayRecord{})
	_ = dbConnection.AutoMigrate(&ChallengeRecordGorm{})
	_ = dbConnection.AutoMigrate(&GameType{})
	_ = dbConnection.AutoMigrate(&GameSubType{})
	_ = dbConnection.AutoMigrate(&ChallengeSubType{})
	_ = dbConnection.AutoMigrate(&TournamentFormat{})
	_ = dbConnection.AutoMigrate(&DepositRecord{})
	_ = dbConnection.AutoMigrate(&Match{})
	_ = dbConnection.AutoMigrate(&Team{})
	_ = dbConnection.AutoMigrate(&Bot{})
	_ = dbConnection.AutoMigrate(&Friend{})
	_ = dbConnection.AutoMigrate(&Featured{})
	_ = dbConnection.AutoMigrate(&AdminConfig{})
	_ = dbConnection.AutoMigrate(&Notification{})
	_ = dbConnection.AutoMigrate(&UserFeed{})
	_ = dbConnection.AutoMigrate(&UserGameAccount{})
	_ = dbConnection.AutoMigrate(&UserPaymentAccount{})
	_ = dbConnection.AutoMigrate(&GamePlatform{})
	_ = dbConnection.AutoMigrate(&Console{})
	_ = dbConnection.AutoMigrate(&TeamMachine{})
	_ = dbConnection.AutoMigrate(&TeamMember{})
	_ = dbConnection.AutoMigrate(&TournamentRequest{})
	_ = dbConnection.AutoMigrate(&TransferLog{})
	_ = dbConnection.AutoMigrate(&MatchGame{})
	_ = dbConnection.AutoMigrate(&SocialLink{})
	db.DB = *dbConnection

	return nil
}
