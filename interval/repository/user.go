package repository

import (
	"gin-blog-newest/interval/model"
	"gin-blog-newest/pkg/logger"
	"gorm.io/gorm"
)

type UserRepository interface {
	// “继承”BaseRepository的方法
	BaseRepository[model.User]
	FindByUsername(username string) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	FindByPhone(phone string) (*model.User, error)
}

type UserRepositoryImpl struct {
	// 继承BaseRepositoryImpl的方法，不用再实现BaseRepository的方法了
	BaseRepositoryImpl[model.User]
}

func NewUserRepository(db *gorm.DB, log *logger.Logger) UserRepository {
	// 初始化BaseRepositoryImpl，传入db
	// 这样就可以使用BaseRepositoryImpl的方法了，定义接受为interface，返回可以是UserRepositoryImpl，也可以是其他的实现
	return &UserRepositoryImpl{
		BaseRepositoryImpl: BaseRepositoryImpl[model.User]{
			DB:  db,
			log: log,
		},
	}
}

func (u UserRepositoryImpl) FindByUsername(username string) (*model.User, error) {
	var user model.User
	err := u.DB.Where("username = ?", username).First(&user).Error
	u.log.Info().Msg("find user by username")
	return &user, err
}

func (u UserRepositoryImpl) FindByEmail(email string) (*model.User, error) {
	var user model.User
	err := u.DB.Where("email =?", email).First(&user).Error
	u.log.Info().Msg("find user by email")
	return &user, err
}

func (u UserRepositoryImpl) FindByPhone(phone string) (*model.User, error) {
	var user model.User
	err := u.DB.Where("phone =?", phone).First(&user).Error
	return &user, err
}
