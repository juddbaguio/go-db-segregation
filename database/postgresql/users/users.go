package users_postgresql

import (
	"encoding/json"
	"go-db-segregation/business/model"
	"go-db-segregation/database/postgresql/entities"

	"gorm.io/gorm"
)

type DB struct {
	dbConn *gorm.DB
}

func New(db *gorm.DB) *DB {
	return &DB{
		dbConn: db,
	}
}

func (db *DB) GetUsers() (*[]model.User, error) {
	var res []model.User
	postgresRes := []entities.User{}

	if err := db.dbConn.Find(&postgresRes).Error; err != nil {
		return nil, err
	}

	// start of deep copy using json struct tag (without looping each of entities.User)
	byteRes, err := json.Marshal(postgresRes)

	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(byteRes, &res); err != nil {
		return nil, err
	}

	// end of deep copy

	return &res, nil
}
