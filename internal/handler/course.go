package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/madiyarsmandiar-lgtm/myProject1/internal/apperrors"
	"github.com/madiyarsmandiar-lgtm/myProject1/internal/model"
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
		respondWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": courses})
}

func (h *CourseHandler) GetByID(c *gin.Context) {
	courseID := c.Param("id")
	id, err := strconv.Atoi(courseID)
	if err != nil {
		respondWithError(c, apperrors.BadRequest("invalid id parameter", nil))
		return
	}
	course, err := h.service.GetByID(id)
	if err != nil {
		respondWithError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": course})
}

func (h *CourseHandler) Delete(c *gin.Context) {
	courseID := c.Param("id")
	id, err := strconv.Atoi(courseID)
	if err != nil {
		respondWithError(c, apperrors.BadRequest("invalid id parameter", nil))
		return
	}
	err = h.service.Delete(id)
	if err != nil {
		respondWithError(c, apperrors.BadRequest("invalid id parameter", nil))
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Course was deleted"})
}

func (h *CourseHandler) Create(c *gin.Context) {
	var input model.CreateCourse

	if err := c.ShouldBindJSON(&input); err != nil {
		respondWithError(c, apperrors.BadRequest("invalid JSON	data", nil))
		return
	}

	newId, err := h.service.Create(input)
	if err != nil {
		respondWithError(c, err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": newId})
}

func (h *CourseHandler) Update(c *gin.Context) {
	courseID := c.Param("id")
	id, err := strconv.Atoi(courseID)
	if err != nil {
		respondWithError(c, apperrors.BadRequest("invalid id parameter", nil))
		return
	}

	var input model.UpdateCourse

	if err := c.ShouldBindJSON(&input); err != nil {
		respondWithError(c, apperrors.BadRequest("invalid id parameter", nil))
		return
	}

	course, err := h.service.Update(id, input)
	if err != nil {
		respondWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": course})
}
