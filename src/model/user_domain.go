package model

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
)

type UserDomainInterface interface {
	SetID(string)
	GetEmail() string
	GetPassword() string
	GetName() string
	GetAge() int
	GetJsonValue() (string, error)
	EncryptPassword()
}

func NewUserDomain(email, password, name string, age int) UserDomainInterface {
	return &userDomain{
		Email:    email,
		Password: password,
		Name:     name,
		Age:      age,
	}
}

type userDomain struct {
	ID       string
	Email    string
	Password string
	Name     string
	Age      int
}

func (ud *userDomain) SetID(id string) {
	ud.ID = id
}

func (ud *userDomain) GetJsonValue() (string, error) {
	b, err := json.Marshal(ud)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (ud *userDomain) GetEmail() string    { return ud.Email }
func (ud *userDomain) GetPassword() string { return ud.Password }
func (ud *userDomain) GetName() string     { return ud.Name }
func (ud *userDomain) GetAge() int         { return ud.Age }

func (ud *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.Password))
	ud.Password = hex.EncodeToString(hash.Sum(nil))
}
