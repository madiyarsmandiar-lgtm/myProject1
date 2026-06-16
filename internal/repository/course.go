package repository

import (
	"errors"
	"time"

	"github.com/madiyarsmandiar-lgtm/myProject1/internal/model"
)

type CourseRepository struct {
	courseMap map[int]*model.Course
}

func NewCourseRepository() *CourseRepository {
	CourseRepo := &CourseRepository{
		courseMap: make(map[int]*model.Course),
	}

	CourseRepo.courseMap[1] = &model.Course{ID: 1, Title: "Go", Price: 2000, IsActive: true, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	CourseRepo.courseMap[2] = &model.Course{ID: 2, Title: "Java", Price: 3000, IsActive: true, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	CourseRepo.courseMap[3] = &model.Course{ID: 3, Title: "Ruby", Price: 1000, IsActive: true, CreatedAt: time.Now(), UpdatedAt: time.Now()}

	return CourseRepo
}

func (cr *CourseRepository) GetAll() ([]model.Course, error) {
	var courses []model.Course

	for _, course := range cr.courseMap {
		courses = append(courses, *course)
	}
	return courses, nil
}

func (cr *CourseRepository) GetById(id int) (*model.Course, error) {
	course, ok := cr.courseMap[id]
	if !ok {
		return &model.Course{}, errors.New("course not found")
	}
	return course, nil
}
