package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/madiyarsmandiar-lgtm/myProject1/internal/service"
)

type CourseHandler struct {
	service *service.CourseService
}

func NewCourseHandler(courseServ *service.CourseService) *CourseHandler {
	return &CourseHandler{service: courseServ}
}

func (h *CourseHandler) GetAll(c *gin.Context) {
	courses, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": courses})
}

func (h *CourseHandler) GetByID(c *gin.Context) {
	courseID := c.Param("id")
	id, err := strconv.Atoi(courseID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid course ID"})
	}
	course, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Course Not Found"})
	}
	c.JSON(http.StatusOK, gin.H{"data": course})
}
