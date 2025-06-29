package routers

import (
	"gin-fleamarket/di"
	"gin-fleamarket/infra"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	// インフラの初期化
	infra.Initialize()
	db := infra.SetupDB()

	r := gin.Default()

	// 依存性注入（インターフェースとして取得）
	itemController := di.NewItemController(db)
	authController := di.NewAuthController(db)

	// API v1 グループ
	v1 := r.Group("/v1")
	{
		// アイテム関連のルート
		items := v1.Group("/items")
		{
			items.GET("", itemController.FindAll)
			items.GET("/:id", itemController.FindById)
			items.POST("", itemController.Create)
			items.PUT("/:id", itemController.Update)
			items.DELETE("/:id", itemController.Delete)
		}

		// 認証関連のルート
		auth := v1.Group("/auth")
		{
			auth.POST("/signup", authController.SignUp)
			auth.POST("/login", authController.Login)
		}
	}

	return r
}
