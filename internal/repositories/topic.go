package repositories

import (
	"discourse-notification/model"
	"gorm.io/gorm"
)

type TopicRepository interface {
	GetCategoryByID(id uint64) (*model.Category, error)
}

type topicRepository struct {
	db *gorm.DB
}

func (t topicRepository) GetCategoryByID(id uint64) (*model.Category, error) {
	var category model.Category
	if err := t.db.Where("id = ?", id).First(&category).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func NewTopicRepository(db *gorm.DB) TopicRepository {
	return &topicRepository{db: db}
}
