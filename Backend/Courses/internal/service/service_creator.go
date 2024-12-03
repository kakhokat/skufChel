package service

import (
	"CoursesBack/internal/models"
	"CoursesBack/internal/store"
)

type CreatorService struct {
	repo store.Creator
}

func NewCreatorService(repo store.Creator) *CreatorService {
	return &CreatorService{
		repo: repo,
	}
}

func (s *CreatorService) CreateCourse(req *models.Course, userId int) (int, error) {
	return s.repo.CreateCourse(req, userId)
}

func (s *CreatorService) GetCreatorCourses(userId int, searchCourse string) ([]models.Course, error) {
	return s.repo.GetCreatorCourses(userId, searchCourse)
}

func (s *CreatorService) DeleteCourse(userId, courseId int) (bool, error) {
	return s.repo.DeleteCourse(userId, courseId)
}
