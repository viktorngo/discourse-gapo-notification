package repositories

import (
	"discourse-notification/internal/core/port"
	"discourse-notification/model"
	"gorm.io/gorm"
)

type topicRepository struct {
	db *gorm.DB
}

func (repo topicRepository) GetTopicByID(id int) (*model.Topic, error) {
	var topic model.Topic
	if err := repo.db.Where("id = ?", id).First(&topic).Error; err != nil {
		return nil, err
	}
	return &topic, nil
}

func (repo topicRepository) GetCategoryByID(id uint64) (*model.Category, error) {
	var category model.Category
	if err := repo.db.Where("id = ?", id).First(&category).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func NewTopicRepository(db *gorm.DB) port.TopicRepository {
	return &topicRepository{db: db}
}
