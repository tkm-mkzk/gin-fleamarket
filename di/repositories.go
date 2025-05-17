package di

import (
	"gin-fleamarket/repositories"

	"gorm.io/gorm"
)

func NewItemRepository(db *gorm.DB) repositories.IItemRepository {
	return repositories.NewItemRepository(db)
}

func NewAuthRepository(db *gorm.DB) repositories.IAuthRepository {
	return repositories.NewAuthRepository(db)
}
