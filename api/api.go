package api

import (
	"github.com/gin-gonic/gin"
	_ "github.com/nurmuhammaddeveloper/Note/api/docs"
	v1 "github.com/nurmuhammaddeveloper/Note/api/v1"
	"github.com/nurmuhammaddeveloper/Note/config"
	"github.com/nurmuhammaddeveloper/Note/storage"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type RouterOptions struct {
	Cfg     *config.Config
	Storage storage.StorageI
}

// @title TodoApp
// @version 1.2.3
// @description Note app desc
// @host localhost:8001
// @BasePath /v1
func New(opt *RouterOptions) *gin.Engine {
	router := gin.Default()
	handlerv1 := v1.New(&v1.Handlerv1option{
		Cfg:     opt.Cfg,
		Storage: opt.Storage,
	})
	aipV1 := router.Group("/v1")
	// User api
	aipV1.GET("/users/:id", handlerv1.GetUser)
	aipV1.POST("/users", handlerv1.CreateUser)
	aipV1.PUT("/users/:id", handlerv1.UpdateUser)
	aipV1.DELETE("/users/:id", handlerv1.DeleteUser)
	// Note api
	aipV1.POST("/notes", handlerv1.CreateNote)




	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}
