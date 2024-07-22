package api

import (
	"discourse-notification/internal/adapter/external_hook"
	"discourse-notification/internal/core/service"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type PostHandler struct {
	topicService    *service.TopicService
	categoryService *service.CategoryService
	workflowHook    *external_hook.WorkflowHook
}

func NewPostHandler(topicService *service.TopicService, categoryService *service.CategoryService, workflowHook *external_hook.WorkflowHook) *PostHandler {
	return &PostHandler{
		topicService:    topicService,
		categoryService: categoryService,
		workflowHook:    workflowHook,
	}
}

func (h PostHandler) PostCreated(c *fiber.Ctx) error {
	var req PostCreatedReq
	if err := c.BodyParser(&req); err != nil {
		log.Errorf("failed to parse request: %v", err)
		return err
	}

	// when first post in topic (topic's content) created send to workflow
	if req.Post.PostNumber == 1 {
		topic, err := h.topicService.GetTopicByID(c.UserContext(), req.Post.TopicId)
		if err != nil {
			return err
		}

		// get category information
		category, err := h.categoryService.GetCategoryByID(c.UserContext(), topic.CategoryID)
		if err != nil {
			return err
		}

		// send topic information to Workflow hook
		err = h.workflowHook.SendTopic(c.UserContext(), external_hook.TopicReq{
			CreatorUsername: req.Post.Username,
			CreatorFullName: req.Post.Name,
			TopicID:         topic.ID,
			TopicTitle:      topic.Title,
			TopicContent:    req.Post.Raw,
			Views:           topic.Views,
			Likes:           topic.LikeCount,
			Event:           "topic_created",
			CategoryID:      int(category.ID),
			CategoryName:    category.Name,
		})
		if err != nil {
			return fmt.Errorf("failed to send topic to workflow: %w", err)
		}

		return c.SendString("sent hook workflow")
	}

	return c.SendString("nothing change")
}

func (h PostHandler) PostEdited(c *fiber.Ctx) error {
	var req PostEditedReq
	if err := c.BodyParser(&req); err != nil {
		log.Errorf("failed to parse request: %v", err)
		return err
	}

	// when first post in topic (topic's content) created send to workflow
	if req.Post.PostNumber == 1 {
		topic, err := h.topicService.GetTopicByID(c.UserContext(), req.Post.TopicId)
		if err != nil {
			return err
		}

		// get category information
		category, err := h.categoryService.GetCategoryByID(c.UserContext(), topic.CategoryID)
		if err != nil {
			return err
		}

		if topic.DeletedAt != nil {
			return c.SendString("topic is deleted")
		}

		// send topic information to Workflow hook
		err = h.workflowHook.SendTopic(c.UserContext(), external_hook.TopicReq{
			CreatorUsername: req.Post.Username,
			CreatorFullName: req.Post.Name,
			TopicID:         topic.ID,
			TopicTitle:      topic.Title,
			TopicContent:    req.Post.Raw,
			Views:           topic.Views,
			Likes:           topic.LikeCount,
			Event:           "topic_edited",
			CategoryID:      int(category.ID),
			CategoryName:    category.Name,
		})
		if err != nil {
			return fmt.Errorf("failed to send topic to workflow: %w", err)
		}

		return c.SendString("sent hook workflow")
	}

	return c.SendString("nothing change")
}
