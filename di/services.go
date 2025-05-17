package di

import (
	"gin-fleamarket/services"

	"gorm.io/gorm"
)

func NewItemService(db *gorm.DB) services.IItemService {
	return services.NewItemService(
		NewItemRepository(db),
	)
}

func NewAuthService(db *gorm.DB) services.IAuthService {
	return services.NewAuthService(
		NewAuthRepository(db),
	)
}
