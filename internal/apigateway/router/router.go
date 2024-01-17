package router

import (
	"iman-task/internal/apigateway/controller"
	_ "iman-task/internal/apigateway/controller/docs"
	"iman-task/internal/apigateway/service"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Option struct {
	Service service.Service
}

// @title           IMAN-TASK
// @description     Created by Otajonov Quvonchbek
// @in header
func New(option *Option) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	handler := controller.New(option.Service)

	api := router.Group("/v1")

	// Post
	api.POST("/create-posts", handler.CreatePosts)
	api.GET("/get-posts", handler.GetPosts)
	api.GET("/get-post", handler.GetPostById)
	api.PUT("/update-post", handler.UpdatePost)
	api.DELETE("/delete-post", handler.DeletePost)

	url := ginSwagger.URL("swagger/doc.json")
	api.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}
