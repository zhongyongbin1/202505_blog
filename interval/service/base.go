package service

import (
	"gin-blog-newest/interval/repository"
	"gin-blog-newest/pkg/logger"
)

type BaseService[T any] interface {
	Create(data *T) error
	Update(data *T) error
	Delete(id int64) error
	FindByID(id int64) (*T, error)
	FindAll() ([]*T, error)
}

type BaseServiceImpl[T any] struct {
	BaseRepository repository.BaseRepository[T]
	log            *logger.Logger
}

func NewBaseService[T any](baseRepository repository.BaseRepository[T], logger *logger.Logger) BaseService[T] {
	return &BaseServiceImpl[T]{
		BaseRepository: baseRepository,
		log:            logger,
	}
}
func (b *BaseServiceImpl[T]) Create(data *T) error {
	return b.BaseRepository.Create(data)
}
func (b *BaseServiceImpl[T]) Update(data *T) error {
	return b.BaseRepository.Update(data)
}
func (b *BaseServiceImpl[T]) Delete(id int64) error {
	return b.BaseRepository.Delete(id)
}
func (b *BaseServiceImpl[T]) FindByID(id int64) (*T, error) {
	return b.BaseRepository.FindByID(id)
}
func (b *BaseServiceImpl[T]) FindAll() ([]*T, error) {
	return b.BaseRepository.FindAll()
}
