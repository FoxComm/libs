package generators

import (
	"github.com/FoxComm/FoxComm/models"
	"github.com/FoxComm/FoxComm/spree"
)

func Entity() models.Entity {
	user := spree.User{
		Email:     fake.Email(),
		FirstName: fake.FirstName(),
		LastName:  fake.LastName(),
	}
	return models.NewUserEntity(user)
}
