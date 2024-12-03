package store

import (
	"CoursesBack/internal/models"

	"github.com/jmoiron/sqlx"
)

type CoursesRepo struct {
	DB *sqlx.DB
}

func NewCoursesRepo(db *sqlx.DB) *CoursesRepo {
	return &CoursesRepo{
		DB: db,
	}
}

func (r *CoursesRepo) GetCourses(search string) ([]models.Course, error) {
	courses := make([]models.Course, 0)

	query := `
		select courseId, name, description, likes, creatorId from courses where 1 = 1
	`

	if search != "" {
		key := "%" + search + "%"

		query += "and name like $1"

		err := r.DB.Select(&courses, query, key)

		if err != nil {
			return nil, err
		}

		courses, err = r.setCount(courses)

		if err != nil {
			return nil, err
		}

		return courses, nil
	}

	err := r.DB.Select(&courses, query)

	if err != nil {
		return nil, err
	}

	courses, err = r.setCount(courses)

	if err != nil {
		return nil, err
	}

	return courses, nil

}

func (r *CoursesRepo) setCount(courses []models.Course) ([]models.Course, error) {
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

func (r *CoursesRepo) GetFullCourse(courseId int) (models.Course, error) {
	var course models.Course

	query := `
		select courseId, name, description, likes, creatorId from courses where courseId = $1
	`

	row := r.DB.QueryRow(query, courseId)

	if err := row.Scan(&course.CourseId, &course.Name, &course.Description, &course.Likes, &course.CreatorId); err != nil {
		return models.Course{}, err
	}

	var lessons []models.Lesson

	querySecond := `
		select name, description, likes, url, question, answers
	    from lessons join course_lesson on course_lesson.lessonId = lessons.lessonId
		join test on test.testId = lessons.testId
		join video on video.videoId = lessons.videoId
	    where course_lesson.courseId = $1
	`

	rows, err := r.DB.Query(querySecond, courseId)

	if err != nil {
		return models.Course{}, err
	}

	for rows.Next() {

		var lesson models.Lesson

		if err := rows.Scan(&lesson.Name, &lesson.Description, &lesson.Likes, &lesson.Video.Url, &lesson.Test.Questiong, &lesson.Test.Answers); err != nil {
			return models.Course{}, err
		}

		lessons = append(lessons, lesson)
	}

	queryThird := `
			select count(*) from courses_users where courseId = $1
		`

	var count int
	if err := r.DB.QueryRow(queryThird, course.CourseId).Scan(&count); err != nil {
		return models.Course{}, err
	}

	course.CourseUsers = count

	queryFourth := `
			select count(*) from course_lesson where courseId = $1
		`

	if err := r.DB.QueryRow(queryFourth, course.CourseId).Scan(&count); err != nil {
		return models.Course{}, err
	}

	course.CourseLessons = count

	course.Lessons = lessons

	return course, nil
}

func (r *CoursesRepo) LikeCourse(userId, courseId int) (bool, error) {

	tx, err := r.DB.Begin()

	if err != nil {
		tx.Rollback()
		return false, err
	}

	queryFirst := `
		select count(*) from courses_users where userId = $1 and courseId = $2 and isLiked = true
	`

	var count int
	if err := tx.QueryRow(queryFirst, userId, courseId).Scan(&count); err != nil {
		tx.Rollback()
		return false, err
	}

	if count == 1 {
		query := `
			update courses set likes = likes - 1 where courseId = $1
		`

		_, err = tx.Exec(query, courseId)

		if err != nil {
			tx.Rollback()
			return false, err
		}

		querySecond := `
			update courses_users set isLiked = false where userId = $1 and courseId = $2 
		`

		_, err = tx.Exec(querySecond, userId, courseId)

		if err != nil {
			tx.Rollback()
			return false, err
		}

		return true, tx.Commit()
	} else {
		query := `
		update courses set likes = likes + 1 where courseId = $1
	`

		_, err = tx.Exec(query, courseId)

		if err != nil {
			tx.Rollback()
			return false, err
		}

		querySecond := `
		update courses_users set isLiked = true where userId = $1 and courseId = $2 
	`

		_, err = tx.Exec(querySecond, userId, courseId)

		if err != nil {
			tx.Rollback()
			return false, err
		}

		return true, tx.Commit()
	}
}

func (r *CoursesRepo) SubscribeCourse(userId, courseId int) (bool, error) {
	var lessonIds []int

	tx, err := r.DB.Begin()

	if err != nil {
		tx.Rollback()
		return false, err
	}

	queryFirst := `
		select lessonId from course_lesson where courseId = $1
	`

	rows, err := tx.Query(queryFirst, courseId)

	if err != nil {
		tx.Rollback()
		return false, err
	}

	for rows.Next() {
		var lessonId int
		if err := rows.Scan(&lessonId); err != nil {
			tx.Rollback()
			return false, err
		}
		lessonIds = append(lessonIds, lessonId)
	}

	if err := rows.Err(); err != nil {
		tx.Rollback()
		return false, err
	}

	for _, lessonId := range lessonIds {
		querySecond := `
			insert into users_lessons(userId, lessonId) values($1, $2)
		`

		_, err = tx.Exec(querySecond, userId, lessonId)

		if err != nil {
			tx.Rollback()
			return false, err
		}
	}

	query := `
		insert into courses_users(userId, courseId) values($1, $2)
	`
	_, err = tx.Exec(query, userId, courseId)

	if err != nil {
		tx.Rollback()
		return false, err
	}

	return true, tx.Commit()
}
