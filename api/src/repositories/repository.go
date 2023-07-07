package repositories

import (
	"api/src/entities"

	"gorm.io/gorm"
)

type Repository[T entities.Entity] interface {
	Create(model T) (T, error)
	Update(model T) (T, error)
	Delete(model T) error
	FindAll() ([]T, error)
	FindByID(id uint64) (T, error)
}

type GenericRepository[T entities.Entity] struct {
	db *gorm.DB
}

func NewGenericRepository[T entities.Entity](db *gorm.DB) *GenericRepository[T] {
	return &GenericRepository[T]{db}
}

func (repository GenericRepository[T]) Create(entity T) (*T, error) {
	result := repository.db.Create(&entity)
	if result.Error != nil {
		return nil, result.Error
	}

	return &entity, nil
}

func (repository GenericRepository[T]) Update(entity T) (*T, error) {
	result := repository.db.Save(&entity)
	if result.Error != nil {
		return nil, result.Error
	}

	return &entity, nil
}

func (repository GenericRepository[T]) Delete(entity T) error {
	result := repository.db.Delete(&entity)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repository GenericRepository[T]) FindAll() (*[]T, error) {
	var entities []T

	result := repository.db.Find(&entities)

	if result.Error != nil {
		return nil, result.Error
	}

	return &entities, nil
}

func (repository GenericRepository[T]) FindByID(id uint) (*T, error) {
	var entity T

	result := repository.db.First(&entity, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &entity, nil
}

func (repository GenericRepository[T]) FindOneBy(field string, value any) (*T, error) {
	var entity T

	result := repository.db.Where(field+" = ?", value).First(&entity)

	if result.Error != nil {
		return nil, result.Error
	}

	return &entity, nil
}

func (repository GenericRepository[T]) FindAllBy(field string, value any) (*[]T, error) {
	var entities []T

	result := repository.db.Where(field+" = ?", value).Find(&entities)

	if result.Error != nil {
		return nil, result.Error
	}

	return &entities, nil
}
