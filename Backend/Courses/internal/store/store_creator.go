package store

import (
	"CoursesBack/internal/models"
	"log/slog"

	"github.com/jmoiron/sqlx"
)

type CreatorRepo struct {
	DB *sqlx.DB
}

func NewCreatorRepo(db *sqlx.DB) *CreatorRepo {
	return &CreatorRepo{DB: db}
}

func (r *CreatorRepo) CreateCourse(req *models.Course, userId int) (int, error) {
	tx, err := r.DB.Begin()
	if err != nil {
		return 0, err
	}

	var courseId int
	query := `
    insert into courses(name, description, creatorId, likes) values($1, $2, $3, $4) returning courseId
`
	row := tx.QueryRow(query, req.Name, req.Description, userId, 0)

	if err := row.Scan(&courseId); err != nil {
		tx.Rollback()
		slog.Error("error with inserting course", "error", err.Error())
		return 0, err
	}

	for _, lesson := range req.Lessons {
		var testId, videoId, lessonId int

		// Insert test
		query := `
        insert into test(question, answers, correctAnswer) values($1, $2, $3) returning testId
    `
		row := tx.QueryRow(query, lesson.Test.Questiong, lesson.Test.Answers, lesson.Test.CurrentAnser)

		if err := row.Scan(&testId); err != nil {
			tx.Rollback()
			slog.Error("error with scanning testId", "error", err.Error())
			return 0, err
		}

		// Insert video
		query = `
        insert into video(url) values($1) returning videoId
    `
		row = tx.QueryRow(query, lesson.Video.Url)

		if err := row.Scan(&videoId); err != nil {
			tx.Rollback()
			slog.Error("error with scanning videoId", "error", err.Error())
			return 0, err
		}

		// Insert lesson
		query = `
        insert into lessons(description, testId, videoId, name, likes) values($1, $2, $3, $4, 0) returning lessonId
    `
		row = tx.QueryRow(query, lesson.Description, testId, videoId, lesson.Name)
		if err := row.Scan(&lessonId); err != nil {
			tx.Rollback()
			slog.Error("error with inserting lesson", "error", err.Error())
			return 0, err
		}

		// Link lesson to course
		query = `
        insert into course_lesson(courseId, lessonId) values($1, $2)
    `
		_, err = tx.Exec(query, courseId, lessonId)
		if err != nil {
			tx.Rollback()
			slog.Error("error with inserting course_lesson", "error", err.Error())
			return 0, err
		}
	}

	return courseId, tx.Commit()

}

func (r *CreatorRepo) GetCreatorCourses(userId int, searchCourse string) ([]models.Course, error) {
	courses := make([]models.Course, 0)

	queryFirst := `
	select courseId, name, description, likes from Courses where creatorId = $1
	` //todo проверить как лайк работает
	if searchCourse != "" {
		key := "%" + searchCourse + "%"

		queryFirst += "and name like $2"

		err := r.DB.Select(&courses, queryFirst, userId, key)

		if err != nil {
			return nil, err
		}

		courses, err = r.setCount(courses)

		if err != nil {
			return nil, err
		}

		return courses, nil
	}

	err := r.DB.Select(&courses, queryFirst, userId)

	if err != nil {
		return nil, err
	}

	courses, err = r.setCount(courses)

	if err != nil {
		return nil, err
	}

	return courses, nil

}

func (r *CreatorRepo) setCount(courses []models.Course) ([]models.Course, error) {
	for i := range courses {

		querySecond := `
			select count(*) from courses_users where courseId = $1
		`

		var count int
		if err := r.DB.QueryRow(querySecond, courses[i].CourseId).Scan(&count); err != nil {
			return nil, err
		}

		courses[i].CourseUsers = count

		queryThird := `
			select count(*) from course_lesson where courseId = $1
		`

		if err := r.DB.QueryRow(queryThird, courses[i].CourseId).Scan(&count); err != nil {
			return nil, err
		}

		courses[i].CourseLessons = count

	}

	return courses, nil
}

func (r *CreatorRepo) DeleteCourse(userId, courseId int) (bool, error) {
	var count int

	query := `
	delete from course where creatorId = $1 and courseId = $2 select @@rowcount
	`

	//todo проверить работает ли @@rowcount

	row := r.DB.QueryRow(query, userId, courseId)

	if err := row.Scan(&count); err != nil {
		return false, err
	}

	return count == 1, nil
}
