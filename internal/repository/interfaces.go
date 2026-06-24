package repository

import "github.com/madiyarsmandiar-lgtm/myProject1/internal/model"

type CourseRepositoryI interface {
	GetAll() ([]model.Course, error)
	GetById(id int) (*model.Course, error)
	Delete(id int) error
	Create(course *model.Course) (int, error)
	Update(id int, input *model.UpdateCourse) (model.Course, error)
}
