package repositories

import (
	"discourse-notification/internal/core/port"
	"discourse-notification/model"
	"gorm.io/gorm"
)

type categoryRepository struct {
	db *gorm.DB
}

func (repo categoryRepository) GetCategoryByID(id int) (*model.Category, error) {
	var category model.Category
	if err := repo.db.First(&category, id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func NewcategoryRepository(db *gorm.DB) port.CategoryRepository {
	return &categoryRepository{db: db}
}
