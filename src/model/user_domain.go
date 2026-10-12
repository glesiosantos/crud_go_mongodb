package model

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/glesiosantos/crud_go_mongodb/src/config/errorsc"
)

func NewUserDomain(email, password, name string, age int) UserDomainInterface {
	return &userDomain{email, password, name, age}
}

type userDomain struct {
	email    string
	password string
	name     string
	age      int
}

func (ud *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.password))
	ud.password = hex.EncodeToString(hash.Sum(nil))
}

type UserDomainInterface interface {
	CreateUser() *errorsc.RestErr
	UpdateUser(string) *errorsc.RestErr
	FindUser(string) (*userDomain, *errorsc.RestErr)
	DeleteUser(string) (*userDomain, *errorsc.RestErr)
}
