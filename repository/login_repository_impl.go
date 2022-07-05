package repository

/*
import (
	"database/sql"
	"errors"
)

type loginRepositoryImpl struct {
}

func NewLoginRepository() LoginRepository {
	return &loginRepositoryImpl{}
}

func (repository *loginRepositoryImpl) Login(db *sql.DB, login *domain.LoginDomain) (*domain.LoginDomain, error) {
	SQL := `SELECT username, password FROM public.login WHERE username = $1 AND password = $2`
	row, err := db.Query(SQL, login.Username, login.Password)

	domain := new(domain.LoginDomain)

	if err != nil {
		return nil, err
	}

	defer row.Close()

	if row.Next() {
		err = row.Scan(&domain.Username, &domain.Password)
		if err != nil {
			return nil, err
		}

		return domain, nil
	}

	return nil, errors.New("wrong username or password")
}

func (repository *loginRepositoryImpl) CreateLogin(db *sql.DB, login *domain.LoginDomain) (*domain.LoginDomain, error) {
	SQL := `INSERT INTO public.login (username, password) VALUES ($1, $2)`
	_, err := db.Exec(SQL, login.Username, login.Password)

	if err != nil {
		return nil, errors.New("username already exist")
	}

	return login, nil
}
*/
