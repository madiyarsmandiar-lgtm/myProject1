package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/madiyarsmandiar-lgtm/myProject1/internal/handler"
	"github.com/madiyarsmandiar-lgtm/myProject1/internal/repository"
	"github.com/madiyarsmandiar-lgtm/myProject1/internal/service"
)

func main() {
	courseRepo := repository.NewCourseRepository()
	courseService := service.NewCourseService(courseRepo)
	courseHandler := handler.NewCourseHandler(courseService)

	router := gin.New()

	router.GET("/courses", courseHandler.GetAll)
	router.GET("/courses/:id", courseHandler.GetByID)

	srv := &http.Server{Addr: ":8080", Handler: router}

	if err := srv.ListenAndServe(); err != nil {
		fmt.Println(err)
	}
}
