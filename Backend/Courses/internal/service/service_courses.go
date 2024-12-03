package service

import (
	"CoursesBack/internal/models"
	"CoursesBack/internal/store"
)

type CourseService struct {
	repo store.Courses
}

func NewCourseSerivec(repo store.Courses) *CourseService {
	return &CourseService{
		repo: repo,
	}
}

func (r *CourseService) GetCourses(search string) ([]models.Course, error) {
	return r.repo.GetCourses(search)
}

func (r *CourseService) GetFullCourse(courseId int) (models.Course, error) {
	return r.repo.GetFullCourse(courseId)
}

func (r *CourseService) LikeCourse(userId, courseId int) (bool, error) {
	return r.repo.LikeCourse(userId, courseId)
}

func (r *CourseService) SubscribeCourse(userId, courseId int) (bool, error) {
	return r.repo.SubscribeCourse(userId, courseId)
}
