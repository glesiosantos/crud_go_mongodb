package converter

import (
	"github.com/glesiosantos/crud_go_mongodb/src/model"
	"github.com/glesiosantos/crud_go_mongodb/src/model/repository/entity"
)

func ConverterEntityToDomain(entity entity.UserEntity) model.UserDomainInterface {
	domain := model.NewUserDomain(
		entity.Email,
		entity.Password,
		entity.Name,
		entity.Age,
	)

	domain.SetID(entity.ID.Hex())
	return domain
}
