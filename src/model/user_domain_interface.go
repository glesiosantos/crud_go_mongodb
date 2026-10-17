package model

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetName() string
	GetAge() int
	GetId() string
	SetID(string)

	GetJsonValue() (string, error)
	EncryptPassword()
}
