package auth

import (
	"database/sql"
	"fmt"
	"sirclo/layered/relation/delivery/common"
	"sirclo/layered/relation/entities"
)

type AuthRepository struct{
	db *sql.DB
}

func New(db *sql.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (ar AuthRepository)Login(identity common.LoginRequestFormat) (entities.User, error) {
	var user entities.User
	result, err := ar.db.Query("SELECT * FROM users WHERE email = ? AND password = ?", identity.Email, identity.Password)
	if err != nil {
		return user, err
	}
	if isExist := result.Next(); !isExist{
		return user, fmt.Errorf("user not exist")
	}
	ErrorScan := result.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Address)
	if ErrorScan != nil {
		fmt.Print(ErrorScan)
		return user, fmt.Errorf("error scan data")
	}
	// if user exist
	if user.Email == identity.Email {
		return user, nil
	}
	// if user not found
	return user, fmt.Errorf("user not found")
}