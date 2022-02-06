package auth

import (
	"sirclo/layered/relation/delivery/common"
	"sirclo/layered/relation/entities"
)

type AuthInterface interface{
	Login(identity common.LoginRequestFormat) (entities.User, error)
}