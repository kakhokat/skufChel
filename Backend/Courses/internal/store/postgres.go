package store

import (
	"CoursesBack/internal/models"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Config struct {
	Host     string
	Port     int
	Username string
	Password string
	Dbname   string
	Sslmode  string
}

func NewConfig(Host string, Port int, Username, Password, Dbname, Sslmode string) Config {
	return Config{
		Host:     Host,
		Port:     Port,
		Username: Username,
		Password: Password,
		Dbname:   Dbname,
		Sslmode:  Sslmode,
	}
}

func InitPostgres(cfg Config) (*sqlx.DB, error) {
	conn, err := sqlx.Connect("postgres", fmt.Sprintf("host =%s port =%d user =%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.Dbname, cfg.Password, cfg.Sslmode))
	if err != nil {
		return nil, err
	}
	return conn, nil
}

type Store struct {
	Auth
	Creator
	Courses
	Lesson
}

func NewStore(DB *sqlx.DB) Store {
	return Store{
		Auth:    NewAuthRepo(DB),
		Creator: NewCreatorRepo(DB),
		Courses: NewCoursesRepo(DB),
		Lesson:  NewLessonStore(DB),
	}
}

type Auth interface {
	GetIsConfirmed(id int) (bool, error)
	CheckKey(key, mail string) (bool, error)
	SignUp(request models.SignUpRequest, randomInt int) (bool, error)
	SignIn(email, password string) (int, bool, error)
	CheckCreator(id int) (bool, error)
	GetUserById(id int) (models.User, error)
	SetCreator(id int) (bool, error)
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
