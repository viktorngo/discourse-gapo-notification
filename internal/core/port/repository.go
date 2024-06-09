package port

import "discourse-notification/model"

type PostRepository interface {
	GetFirstPostByTopicID(topicID int) (*model.Post, error)
}

type TopicRepository interface {
	GetTopicByID(id int) (*model.Topic, error)
	GetCategoryByID(id uint64) (*model.Category, error)
}

type UserRepository interface {
	GetUsersByGroupID(groupID uint64) ([]model.User, error)
}

type CategoryRepository interface {
	GetCategoryByID(id int) (*model.Category, error)
}
