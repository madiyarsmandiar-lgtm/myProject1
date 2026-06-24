package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/madiyarsmandiar-lgtm/myProject1/internal/config"
	"github.com/madiyarsmandiar-lgtm/myProject1/internal/handler"
	"github.com/madiyarsmandiar-lgtm/myProject1/internal/repository"
	"github.com/madiyarsmandiar-lgtm/myProject1/internal/service"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}
	courseRepo := repository.NewCourseRepository()
	courseService := service.NewCourseService(courseRepo)
	courseHandler := handler.NewCourseHandler(courseService)

	router := gin.New()

	router.GET("api/courses", courseHandler.GetAll)
	router.GET("api/courses/:id", courseHandler.GetByID)
	router.DELETE("api/courses/:id", courseHandler.Delete)
	router.POST("api/courses/", courseHandler.Create)
	router.PUT("api/courses/:id", courseHandler.Update)

	srv := &http.Server{Addr: ":" + cfg.Port, Handler: router}

	if err := srv.ListenAndServe(); err != nil {
		fmt.Println(err)
	}
}
