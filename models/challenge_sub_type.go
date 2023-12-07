package models

// PlayRecord maps the users play record table in the database.
type ChallengeSubType struct {
	ChallengeType string `json:"challengeType" gorm:"type:text;primary_key"`
	Subtype       string `json:"challengeSubType" gorm:"type:text;primary_key"`
}

func (db *DB) CreateChallengeSubType(challenge_type, challenge_sub_type string) error {
	challenge := ChallengeSubType{
		ChallengeType: challenge_type,
		Subtype:       challenge_sub_type,
	}
	return db.Save(&challenge).Error
}

func (db *DB) GetChallengeTypes() ([]ChallengeSubType, error) {
	challengeTypes := []ChallengeSubType{}

	err := db.Find(&challengeTypes).Error
	if err != nil {
		return nil, err
	}

	return challengeTypes, nil
}
