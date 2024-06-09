package repositories

import (
	"discourse-notification/model"
	"gorm.io/gorm"
)

type postRepository struct {
	db *gorm.DB
}

// GetFirstPostByTopicID get first post by topic id (first post is content of topic)
func (repo postRepository) GetFirstPostByTopicID(topicID int) (*model.Post, error) {
	var post model.Post
	if err := repo.db.Where("topic_id = ? and post_number = 1", topicID).First(&post).Error; err != nil {
		return nil, err
	}

	return &post, nil
}
