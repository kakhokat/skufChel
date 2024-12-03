package store

import (
	"strconv"

	"github.com/jmoiron/sqlx"

	"CoursesBack/internal/models"
)

type AuthRepo struct {
	Db *sqlx.DB
}

func NewAuthRepo(db *sqlx.DB) *AuthRepo {
	return &AuthRepo{Db: db}
}

func (r *AuthRepo) SignUp(request models.SignUpRequest, randomInt int) (bool, error) {
	if request.Birthday == "" {
		request.Birthday = "08.11.2004"
	}
	var query string
	if request.Photo == nil {
		query = `
	insert into users(email, password, name, birthday, iscreator, checkInt, isConfirmed)
	values ($1, $2, $3, $4, false, $5, false)
	`
		_, err := r.Db.Exec(query, request.Email, request.Password, request.Username, request.Birthday, randomInt)

		if err != nil {
			return false, err
		}
	} else {
		query = `
	insert into users(email, password, name, birthday, image, iscreator, checkInt, isConfirmed)
	values ($1, $2, $3, $4, $5, false, $6, false)
`
		_, err := r.Db.Exec(query, request.Email, request.Password, request.Username, request.Birthday, request.Photo, randomInt)

		if err != nil {
			return false, err
		}
	}

	return true, nil
}

func (r *AuthRepo) SignIn(email, password string) (int, bool, error) {
	query := `
	select id, isconfirmed from users where email = $1 and password = $2
`

	var id int
	var isconfirmed bool

	row := r.Db.QueryRow(query, email, password)

	if err := row.Scan(&id, &isconfirmed); err != nil {
		return 0, false, err
	}

	return id, isconfirmed, nil
}

func (r *AuthRepo) CheckKey(key, mail string) (bool, error) {

	keyInt, err := strconv.Atoi(key)

	if err != nil {
		return false, err
	}

	query := `
		update users set isconfirmed = true where email = $1 and checkint = $2
	`

	rows, err := r.Db.Exec(query, mail, keyInt)

	if err != nil {
		return false, err
	}

	count, err := rows.RowsAffected()

	if err != nil {
		return false, err
	}

	return count == 1, nil
}

func (r *AuthRepo) CheckCreator(id int) (bool, error) {
	query := `
	select isCreator from users where id = $1
	`

	var isCreator bool
	if err := r.Db.Get(&isCreator, query, id); err != nil {
		return false, err
	}

	return isCreator, nil
}

func (r *AuthRepo) GetIsConfirmed(id int) (bool, error) {
	query := `
		select isconfirmed from users where id = $1
	`

	var isConfirmed bool

	row := r.Db.QueryRow(query, id)

	if err := row.Scan(&isConfirmed); err != nil {
		return false, err
	}

	return isConfirmed, nil
}

func (r *AuthRepo) GetUserById(id int) (models.User, error) {
	query := `
		select id, email, name, birthday, isCreator
		from users where id = $1
	`

	var user models.User

	row := r.Db.QueryRow(query, id)

	if err := row.Scan(&user.Id, &user.Email, &user.Name, &user.Birthday, &user.IsCreator); err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *AuthRepo) SetCreator(id int) (bool, error) {
	query := `
		update users set isCreator = true where id = $1
	`

	_, err := r.Db.Exec(query, id)

	if err != nil {
		return false, err
	}

	return true, nil
}
