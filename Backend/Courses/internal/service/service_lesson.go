package service

import (
	"CoursesBack/internal/models"
	"CoursesBack/internal/store"
)

type LessonService struct {
	repo store.Lesson
}

func NewLessonService(repo store.Lesson) *LessonService {
	return &LessonService{
		repo: repo,
	}
}

func (s *LessonService) GetLesson(id, userId int) (models.LessonFull, error) {
	return s.repo.GetLesson(id, userId)
}

func (s *LessonService) AnswerTest(lessonId, userId int, answer string) (models.ResultOfAnswering, error) {
	return s.repo.AnswerTest(lessonId, userId, answer)
}
