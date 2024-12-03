package store

import (
	"CoursesBack/internal/models"

	"github.com/jmoiron/sqlx"
)

type LessonStore struct {
	db *sqlx.DB
}

func NewLessonStore(db *sqlx.DB) *LessonStore {
	return &LessonStore{
		db: db,
	}
}

func (r *LessonStore) GetLesson(id, userId int) (models.LessonFull, error) {
	var lesson models.LessonFull

	tx, err := r.db.Begin()

	if err != nil {
		tx.Rollback()
		return models.LessonFull{}, err
	}

	var isLiked, isPassed bool

	query := `
		select isPassed, isLiked from users_lessons
		where lessonId = $1 and userId = $2
	`

	row := tx.QueryRow(query, id, userId)

	if err := row.Scan(&isPassed, &isLiked); err != nil {
		tx.Rollback()
		return models.LessonFull{}, err
	}

	lesson.IsLiked = isLiked
	lesson.IsPassed = isPassed

	var querySecond string

	if isPassed {
		querySecond = `
			select name, description, likes, question, answers, correctAnswer, url
			from lessons 
			join test on test.testId = lessons.testId
			join video ont video.videoId = lessons.videoId
			where lessons.lessonId = $1
		`

		row := tx.QueryRow(querySecond, id)

		if err := row.Scan(&lesson.Name, &lesson.Description, &lesson.Likes, &lesson.Test.Questiong, &lesson.Test.Answers, &lesson.Test.CurrentAnser, &lesson.Video.Url); err != nil {
			tx.Rollback()
			return models.LessonFull{}, err
		}

	} else {
		querySecond = `
			select name, description, likes, question, url
			from lessons 
			join test on test.testId = lessons.testId
			join video ont video.videoId = lessons.videoId
			where lessons.lessonId = $1
		`

		row := tx.QueryRow(querySecond, id)

		if err := row.Scan(&lesson.Name, &lesson.Description, &lesson.Likes, &lesson.Test.Questiong, &lesson.Video.Url); err != nil {
			tx.Rollback()
			return models.LessonFull{}, err
		}

	}

	return lesson, tx.Commit()

}

func (r *LessonStore) AnswerTest(lessonId, userId int, answer string) (models.ResultOfAnswering, error) {

	var result models.ResultOfAnswering

	query := `
		select correctAnswer == $1 from test
		join lessons on lessons.testId = test.testId
		join users_lessons on users_lessons.lessonId = lessons.lessonId
		where lessons.lessonId = $2 and users_lessons.userId = $3
	`

	row := r.db.QueryRow(query, answer, lessonId, userId)

	if err := row.Scan(&result.IsCorrect); err != nil {
		return models.ResultOfAnswering{}, err
	}

	return result, nil
}
