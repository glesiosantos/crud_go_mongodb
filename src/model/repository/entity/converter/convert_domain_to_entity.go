package converter

import (
	"github.com/glesiosantos/crud_go_mongodb/src/model"
	"github.com/glesiosantos/crud_go_mongodb/src/model/repository/entity"
)

func ConverterDomainToEntity(domain model.UserDomainInterface) *entity.UserEntity {
	return &entity.UserEntity{
		Email:    domain.GetEmail(),
		Name:     domain.GetName(),
		Password: domain.GetPassword(),
		Age:      domain.GetAge(),
	}
}
