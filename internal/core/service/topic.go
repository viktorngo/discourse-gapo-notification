package service

import (
	"context"
	"discourse-notification/internal/core/port"
	"discourse-notification/model"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
)

type TopicService struct {
	DiscourseHost      string
	Notification       port.Notification
	TopicRepository    port.TopicRepository
	UserRepository     port.UserRepository
	CategoryRepository port.CategoryRepository
}

func (service TopicService) SendNoticeTopicCreated(ctx context.Context, topic *model.Topic, creator *model.User) error {
	// get category information
	category, err := service.CategoryRepository.GetCategoryByID(topic.CategoryID)
	if err != nil {
		return err
	}

	// get users in reviewable group by group id
	users, err := service.UserRepository.GetUsersByGroupID(category.ReviewableByGroupID)
	if err != nil {
		return err
	}

	// send notification to users
	redirectURL := fmt.Sprintf("%s/t/%s/%d", service.DiscourseHost, topic.Slug, topic.ID)
	for _, user := range users {
		if user.ID != creator.ID {
			if err := service.Notification.SendTopicCreatedNotification(user.Username, topic.Title, creator.Name, category.Name, redirectURL); err != nil {
				log.Errorf("failed to send Gapo notification for `topic created` to user `%s`: %v", user.Username, err)
				continue
			}
		}
	}
	return nil
}

func (service TopicService) GetTopicByID(ctx context.Context, id int) (*model.Topic, error) {
	topic, err := service.TopicRepository.GetTopicByID(id)
	if err != nil {
		return nil, err
	}

	return topic, nil
}
