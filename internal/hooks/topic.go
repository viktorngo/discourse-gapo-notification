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

func (hooker TopicHooker) TopicCreated(categoryID uint64, topicID uint64, topicSlug string, title string, createdByID uint64, createdByName string) error {
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
	message := `**📢 Thông báo: Chủ đề mới đã được mở!**
**%s** đã tạo chủ đề **%s** thuộc danh mục **%s** do bạn quản lý.
[Xem chủ đề!](%s)`

	redirectURL := fmt.Sprintf("%s/t/%s/%d", hooker.DiscourseHost, topicSlug, topicID)
	msg := fmt.Sprintf(message, createdByName, title, category.Name, redirectURL)
	for _, user := range users {
		if user.ID != createdByID {
			if err := hooker.GapoWorkClient.SendNotification(user.Username, msg); err != nil {
				log.Errorf("failed to send Gapo notification for `topic created` to user `%s`: %v", user.Username, err)
				continue
			}
		}
	}
	return nil
}
