package service

import (
	"CoursesBack/internal/models"
	"CoursesBack/internal/store"

	"github.com/segmentio/kafka-go"
)

type Service struct {
	Auth
	Creator
	Courses
	Lesson
}

func NewService(store store.Store, salt string, kafka *kafka.Conn) *Service {
	return &Service{
		Auth:    NewAuthService(store.Auth, salt, kafka),
		Creator: NewCreatorService(store.Creator),
		Courses: NewCourseSerivec(store.Courses),
		Lesson:  NewLessonService(store.Lesson),
	}
}

type Auth interface {
	CheckKey(key, mail string) (bool, error)
	CheckCreator(id int) (bool, error)
	SignUp(request models.SignUpRequest) (bool, error)
	SignIn(req models.SignUpRequest) (bool, string, error)
	GetUserById(id int) (models.User, error)
	SetCreator(id int) (bool, error)
	GetIsConfirmed(id int) (bool, error)
}

type Creator interface {
	CreateCourse(req *models.Course, userId int) (int, error)
	GetCreatorCourses(userId int, searchCourse string) ([]models.Course, error)
	DeleteCourse(userId, courseId int) (bool, error)
}

type Courses interface {
	GetCourses(search string) ([]models.Course, error)
	GetFullCourse(courseId int) (models.Course, error)
	LikeCourse(userId, courseId int) (bool, error)
	SubscribeCourse(userId, courseId int) (bool, error)
}

type Lesson interface {
	GetLesson(id, userId int) (models.LessonFull, error)
	AnswerTest(lessonId, userId int, answer string) (models.ResultOfAnswering, error)
}
