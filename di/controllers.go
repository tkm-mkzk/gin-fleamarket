package di

import (
	"gin-fleamarket/controllers"

	"gorm.io/gorm"
)

func NewItemController(db *gorm.DB) controllers.IItemController {
	return controllers.NewItemController(
		NewItemService(db),
	)
}

func NewAuthController(db *gorm.DB) controllers.IAuthController {
	return controllers.NewAuthController(
		NewAuthService(db),
	)
}
