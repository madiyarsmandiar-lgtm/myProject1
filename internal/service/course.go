package service

import (
	"time"

	"github.com/madiyarsmandiar-lgtm/myProject1/internal/apperrors"
	"github.com/madiyarsmandiar-lgtm/myProject1/internal/model"
	"github.com/madiyarsmandiar-lgtm/myProject1/internal/repository"
)

type CourseService struct {
	repo repository.CourseRepositoryI
}

func NewCourseService(CourseRepo repository.CourseRepositoryI) *CourseService {
	service := &CourseService{
		repo: CourseRepo,
	}
	return service
}

func (cs *CourseService) GetAll() ([]model.Course, error) {
	courses, err := cs.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return courses, nil
}

func (cs *CourseService) GetByID(id int) (*model.Course, error) {
	course, err := cs.repo.GetById(id)
	if err != nil {
		return nil, err
	}
	return course, nil
}

func (cs *CourseService) Delete(id int) error {
	if id <= 0 {
		return apperrors.BadRequest("invalid id", nil)
	}
	return cs.repo.Delete(id)
}

func (cs *CourseService) Create(input model.CreateCourse) (int, error) {
	if len(input.Title) < 3 {
		return 0, apperrors.BadRequest("title should be greater than 3", nil)
	}
	if input.Price < 0 {
		return 0, apperrors.BadRequest("price should be greater than 0", nil)
	}

	course := model.Course{
		Title:     input.Title,
		Price:     input.Price,
		IsActive:  input.IsActive,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return cs.repo.Create(&course)
}

func (cs *CourseService) Update(id int, input model.UpdateCourse) (model.Course, error) {
	_, err := cs.repo.GetById(id)
	if err != nil {
		return model.Course{}, apperrors.NotFound("course is not found", nil)
	}

	if input.Title != nil && len(*input.Title) < 3 {
		return model.Course{}, apperrors.BadRequest("title should be greater than 3", nil)
	}
	if input.Price != nil && *input.Price < 0 {
		return model.Course{}, apperrors.BadRequest("price should be greater than 0", nil)
	}

	return cs.repo.Update(id, &input)
}
