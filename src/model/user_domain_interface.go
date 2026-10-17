package model

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetName() string
	GetAge() int
	GetId() string
	SetID(string)

	// GetJsonValue() (string, error)
	EncryptPassword()
}

func NewUserDomain(email, password, name string, age int) UserDomainInterface {
	return &userDomain{
		email:    email,
		password: password,
		name:     name,
		age:      age,
	}
}
