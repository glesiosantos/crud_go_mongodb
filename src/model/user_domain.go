package model

import (
	"encoding/json"
)

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

func (ud *userDomain) GetId() string       { return ud.ID }
func (ud *userDomain) GetEmail() string    { return ud.Email }
func (ud *userDomain) GetPassword() string { return ud.Password }
func (ud *userDomain) GetName() string     { return ud.Name }
func (ud *userDomain) GetAge() int         { return ud.Age }
