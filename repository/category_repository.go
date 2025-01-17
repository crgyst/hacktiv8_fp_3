package repository

import (
	"context"
	"hacktiv8_fp_2/entity"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	CreateCategory(ctx context.Context, category entity.Category) (entity.Category, error)
	GetCategory(ctx context.Context, userID uint64) ([]entity.Category, error)
	PatchCategory(ctx context.Context, category entity.Category) (entity.Category, error)
	DeleteCategory(ctx context.Context, categoryID uint64) error
}

type categoryConnection struct {
	connection *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryConnection{
		connection: db,
	}
}

// CreateCategory implements CategoryRepository
func (db *categoryConnection) CreateCategory(ctx context.Context, category entity.Category) (entity.Category, error) {
	tx := db.connection.Create(&category)
	if tx.Error != nil {
		return entity.Category{}, tx.Error
	}
	return category, nil
}

// GetCategory implements CategoryRepository
func (db *categoryConnection) GetCategory(ctx context.Context, userID uint64) ([]entity.Category, error) {
	var category []entity.Category
	tx := db.connection.Preload("User").Find(&category)
	if tx.Error != nil {
		return []entity.Category{}, tx.Error
	}
	return category, nil
}

// PatchCategory implements CategoryRepository
func (db *categoryConnection) PatchCategory(ctx context.Context, category entity.Category) (entity.Category, error) {
	panic("unimplemented")
}

// DeleteCategory implements CategoryRepository
func (db *categoryConnection) DeleteCategory(ctx context.Context, categoryID uint64) error {
	tx := db.connection.Delete(&entity.Category{}, categoryID)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
