package user

import (
	"sirclo/layered/relation/entities"
)

type UserInterface interface{
	GetUsers() ([]entities.User, error)
	PostUser(user entities.User) error 
	GetUser(id int) (entities.User, error) 
	DeleteUser(id int) error
	UpdateUser(id int, user entities.User) error
}