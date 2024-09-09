package repository

import (
	"context"
	"gorm.io/gorm"
)

type Repository[T any] interface {
	Add(*T, context.Context) error
	AddRange(*[]T, context.Context) error
	GetById(int, context.Context) (*T, error)
	GetFirstOrDefault(*T, context.Context) (*T, error)
	Get(*T, context.Context) (*[]T, error)
	GetAll(context.Context) (*[]T, error)
	Update(*T, context.Context) error
	Delete(*T, context.Context) error
	Count(context.Context) (int64, error)
	Exists(*T, context.Context) (bool, error)
}

type ImplementRepository[T any] struct {
	db *gorm.DB
}

func NewImplementRepository[T any](db *gorm.DB) *ImplementRepository[T] {
	return &ImplementRepository[T]{
		db: db,
	}
}

func (repository *ImplementRepository[T]) Add(entity *T, ctx context.Context) error {
	return repository.db.WithContext(ctx).Create(entity).Error
}

func (repository *ImplementRepository[T]) AddRange(entities *[]T, ctx context.Context) error {
	return repository.db.WithContext(ctx).Create(&entities).Error
}

func (repository *ImplementRepository[T]) GetById(id int, ctx context.Context) (*T, error) {
	var entity T
	err := repository.db.WithContext(ctx).Model(&entity).Where("id = ?", id).FirstOrInit(&entity).Error
	return &entity, err
}

func (repository *ImplementRepository[T]) GetFirstOrDefault(params *T, ctx context.Context) (*T, error) {
	var entity T
	err := repository.db.WithContext(ctx).Where(params).FirstOrInit(&entity).Error
	return &entity, err
}

func (repository *ImplementRepository[T]) Get(params *T, ctx context.Context) (*[]T, error) {
	var entities []T
	err := repository.db.WithContext(ctx).Where(params).Find(&entities).Error
	return &entities, err
}

func (repository *ImplementRepository[T]) GetAll(ctx context.Context) (*[]T, error) {
	var entities []T
	err := repository.db.WithContext(ctx).Find(&entities).Error
	return &entities, err
}

func (repository *ImplementRepository[T]) Update(entity *T, ctx context.Context) error {
	return repository.db.WithContext(ctx).Save(entity).Error
}

func (repository *ImplementRepository[T]) Delete(entity *T, ctx context.Context) error {
	return repository.db.WithContext(ctx).Delete(entity).Error
}

func (repository *ImplementRepository[T]) Count(ctx context.Context) (int64, error) {
	var count int64
	var entity T
	err := repository.db.WithContext(ctx).Model(&entity).Count(&count).Error
	return count, err
}

func (repository *ImplementRepository[T]) Exists(entity *T, ctx context.Context) (bool, error) {
	var count int64
	err := repository.db.WithContext(ctx).Model(&entity).Where(entity).Count(&count).Error
	return count > 0, err
}
