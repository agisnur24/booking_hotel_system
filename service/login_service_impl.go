package service

/*
type loginServiceImpl struct {
	LoginRepository loginrepository.LoginRepository
	Db              *sql.DB
}

func NewLoginService(repository loginrepository.LoginRepository, db *sql.DB) LoginService {
	return &loginServiceImpl{
		LoginRepository: repo,
		Db:              db,
	}
}

func (service *loginServiceImpl) Login(request *request.LoginRequest) error {
	login := domain.LoginDomain{
		Username: request.Username,
		Password: request.Password,
	}

	_, err := service.LoginRepository.Login(service.Db, &login)

	if err != nil {
		return err
	}

	return nil
}

func (service *loginServiceImpl) CreateLogin(request *request.LoginRequest) error {
	login := domain.LoginDomain{
		Username: request.Username,
		Password: request.Password,
	}

	_, err := service.LoginRepository.CreateLogin(service.Db, &login)

	if err != nil {
		return err
	}

	return nil
}
*/
