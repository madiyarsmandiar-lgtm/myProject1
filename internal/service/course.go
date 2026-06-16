package service

import (
	"github.com/madiyarsmandiar-lgtm/myProject1/internal/model"
	"github.com/madiyarsmandiar-lgtm/myProject1/internal/repository"
)

type CourseService struct {
	repo *repository.CourseRepository
}

func NewCourseService(CourseRepo *repository.CourseRepository) *CourseService {
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
