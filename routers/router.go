package routers

import (
	"gin-fleamarket/di"
	"gin-fleamarket/infra"
	"gin-fleamarket/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRouter() *gin.Engine {
	// インフラの初期化
	infra.Initialize()
	db := infra.SetupDB()

	return NewRouterWithDB(db)
}

// テスト用にDBインスタンスを受け取れる関数
func NewRouterWithDB(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())

	// 依存性注入（インターフェースとして取得）
	itemController := di.NewItemController(db)
	authController := di.NewAuthController(db)
	authService := di.NewAuthService(db)

	// API v1 グループ
	v1 := r.Group("/v1")
	{
		// アイテム関連のルート
		items := v1.Group("/items")
		itemsWithAuth := v1.Group("/items", middlewares.AuthMiddleware((authService)))
		{
			items.GET("", itemController.FindAll)
			itemsWithAuth.GET("/:id", itemController.FindById)
			itemsWithAuth.POST("", itemController.Create)
			itemsWithAuth.PUT("/:id", itemController.Update)
			itemsWithAuth.DELETE("/:id", itemController.Delete)
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
