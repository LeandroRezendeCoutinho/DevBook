package repositories

import "api/src/entities"

type Repository interface {
	Create(model entities.Entity) (entities.Entity, error)
	Update(model entities.Entity) (entities.Entity, error)
	Delete(model entities.Entity) error
	FindAll() ([]entities.Entity, error)
	FindByID(id uint64) (entities.Entity, error)
}
