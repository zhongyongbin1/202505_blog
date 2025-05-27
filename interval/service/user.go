package service

import (
	"gin-blog-newest/interval/model"
	"gin-blog-newest/interval/repository"
	"gin-blog-newest/pkg/logger"
)

type UserService interface {
	BaseService[model.User]
	FindByUsername(username string) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	FindByPhone(phone string) (*model.User, error)
}

type UserServiceImpl struct {
	BaseServiceImpl[model.User]
	UserRepository repository.UserRepository
	log            *logger.Logger
}

func NewUserService(repo repository.UserRepository, logger *logger.Logger) UserService {
	return &UserServiceImpl{
		BaseServiceImpl: BaseServiceImpl[model.User]{
			BaseRepository: repo,
			log:            logger,
		},
		UserRepository: repo,
	}
}

func (u *UserServiceImpl) FindByUsername(username string) (*model.User, error) {
	user, err := u.UserRepository.FindByUsername(username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserServiceImpl) FindByEmail(email string) (*model.User, error) {
	user, err := u.UserRepository.FindByEmail(email)
	if err != nil {
		return nil, err

	}
	return user, nil
}

func (u *UserServiceImpl) FindByPhone(phone string) (*model.User, error) {
	user, err := u.UserRepository.FindByPhone(phone)
	if err != nil {
		return nil, err
	}
	return user, nil
}
