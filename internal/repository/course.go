package repository

import (
	"time"

	"github.com/madiyarsmandiar-lgtm/myProject1/internal/apperrors"
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
		return &model.Course{}, apperrors.NotFound("course not found", nil)
	}
	return course, nil
}

func (cr *CourseRepository) Delete(id int) error {
	_, ok := cr.courseMap[id]
	if !ok {
		return apperrors.NotFound("course is not found", nil)
	}
	delete(cr.courseMap, id)
	return nil
}

func (cr *CourseRepository) Create(course *model.Course) (int, error) {
	course.ID = len(cr.courseMap) + 1
	cr.courseMap[course.ID] = course

	return course.ID, nil
}

func (cr *CourseRepository) Update(id int, input *model.UpdateCourse) (model.Course, error) {
	course, ok := cr.courseMap[id]
	if !ok {
		return model.Course{}, apperrors.NotFound("course is not found", nil)
	}

	if input.Title != nil {
		course.Title = *input.Title
	}
	if input.Price != nil {
		course.Price = *input.Price
	}
	if input.IsActive != nil {
		course.IsActive = *input.IsActive
	}

	cr.courseMap[id] = course
	return *course, nil
}
