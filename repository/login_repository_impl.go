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

func (repository *loginRepositoryImpl) LoginByEmail(db *sql.DB, login *domain.LoginDomain) (*domain.LoginDomain, error) {
	SQL := `SELECT email, password FROM users WHERE email = ?`
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

func (repository *loginRepositoryImpl) Register(db *sql.DB, login *domain.LoginDomain) (*domain.LoginDomain, error) {
	SQL := `INSERT INTO users (username, email, password) VALUES (?,?,?)`
	_, err := db.Exec(SQL, login.Username, Login.Email login.Password)

	if err != nil {
		return nil, errors.New("username already exist")
	}

	return login, nil
}
*/
