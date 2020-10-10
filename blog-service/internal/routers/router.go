package router

import (
	"github.com/gin-gonic/gin"
	tag "github.com/go-programming-tour-book/blog-service/internal/routers/api/v1"
)

func NewRouter() *gin.Engine{
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	tag := tag.NewTag()

	apiv1:=r.Group("/api/v1")
	{
		apiv1.GET("/tags/:id",tag.Get)
		apiv1.POST("/tags",tag.Create)
	}

	return r
}