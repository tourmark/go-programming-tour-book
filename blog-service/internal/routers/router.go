package routers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/go-programming-tour-book/blog-service/internal/routers/api/v1"
)

func NewRouter(db *sql.DB) *gin.Engine{
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	user := v1.NewUser()

	apiv1:=r.Group("/api/v1")
	{
		apiv1.GET("/users",func(c *gin.Context) {
			c.JSON(200, user.List(c,db))
		})
	}

	return r
}