package repository

import (
	"gin-blog-newest/pkg/logger"
	"gorm.io/gorm"
)

type BaseRepository[T any] interface {
	Create(data *T) error
	Update(data *T) error
	Delete(id int64) error
	FindByID(id int64) (*T, error)
	FindAll() ([]*T, error)
}

type BaseRepositoryImpl[T any] struct {
	DB  *gorm.DB
	log *logger.Logger
}

func (b *BaseRepositoryImpl[T]) Create(data *T) error {
	return b.DB.Create(data).Error
}

func (b *BaseRepositoryImpl[T]) Update(data *T) error {
	return b.DB.Save(data).Error
}

func (b *BaseRepositoryImpl[T]) Delete(id int64) error {
	return b.DB.Delete(id).Error
}

func (b *BaseRepositoryImpl[T]) FindByID(id int64) (*T, error) {
	var data T
	err := b.DB.First(&data, id).Error
	return &data, err
}

func (b *BaseRepositoryImpl[T]) FindAll() ([]*T, error) {
	var data []*T
	err := b.DB.Find(&data).Error
	return data, err
}
