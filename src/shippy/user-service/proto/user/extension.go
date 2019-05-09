package go_micro_srv_user

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/log"
	uuid "github.com/satori/go.uuid"
)

func (user *User) BeforeCreate(scope *gorm.Scope) error {
	uuid, err := uuid.NewV4()
	if err != nil {
		log.Fatalf("created uuid error: %v\n", err)
	}
	return scope.SetColumn("Id", uuid.String())
}
