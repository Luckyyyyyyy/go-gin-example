package routers

import (
	"github.com/gin-gonic/gin"
	"go-gin-example/pkg/setting"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)
	r.GET("/test", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "test",
		})
	})
	//
	//apiv1 := r.Group("/api/v1")
	//{
	//	apiv1.GET("/tags", v1.GetTags)
	//	apiv1.POST("/tags", v1.AddTag)
	//	apiv1.PUT("/tags", v1.EditTag)
	//	apiv1.DELETE("/tags", v1.DeleteTag)
	//
	//}

	return r
}
