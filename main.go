package main

import (
	"fmt"
	"go-gin-example/routers"
	"net/http"

	"go-gin-example/pkg/setting"
)

func main() {

	router := routers.InitRouter()

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
