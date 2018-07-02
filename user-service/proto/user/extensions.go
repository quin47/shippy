package go_micro_srv_user

import (
	"github.com/jinzhu/gorm"
	"github.com/pborman/uuid"
)

func ( *User)BeforeCreate (scope *gorm.Scope) error {
	id := uuid.NewRandom()
	return scope.SetColumn("Id", id.String())
}
