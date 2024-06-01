package hooks

import (
	"discourse-notification/gapo"
	"discourse-notification/internal/repositories"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
)

type TopicHooker struct {
	DiscourseHost   string
	GapoWorkClient  *gapo.Client
	TopicRepository repositories.TopicRepository
	UserRepository  repositories.UserRepository
}

func (hooker TopicHooker) TopicCreated(categoryID uint64, topicID uint64, topicSlug string, title string, createdByName string) error {
	// get category information
	category, err := hooker.TopicRepository.GetCategoryByID(categoryID)
	if err != nil {
		return err
	}

	// get users in reviewable group by group id
	users, err := hooker.UserRepository.GetUsersByGroupID(category.ReviewableByGroupID)
	if err != nil {
		return err
	}

	// send notification to users
	redirectURL := fmt.Sprintf("%s/t/%s/%d", hooker.DiscourseHost, topicSlug, topicID)
	for _, user := range users {
		if err := hooker.GapoWorkClient.SendTopicCreatedNotification(user.Username, title, createdByName, category.Name, redirectURL); err != nil {
			log.Errorf("failed to send Gapo notification for `topic created` to user `%s`: %v", user.Username, err)
			continue
		}
	}
	return nil
}
