package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go-gin-example/pkg/setting"
)

func main() {

	router := gin.Default()
	router.GET("/test", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "test",
		})
	})

	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", setting.HTTPPort),
		Handler: router,

		ReadTimeout:  setting.ReadTimeOut,
		WriteTimeout: setting.WriteTimeOut,
	}
	err := s.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
