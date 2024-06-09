package api

import (
	"discourse-notification/internal/adapter/external_hook"
	"discourse-notification/internal/core/service"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type PostHandler struct {
	topicService *service.TopicService
	workflowHook *external_hook.WorkflowHook
}

func NewPostHandler(topicService *service.TopicService, workflowHook *external_hook.WorkflowHook) *PostHandler {
	return &PostHandler{
		topicService: topicService,
		workflowHook: workflowHook,
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
		topic, err := h.topicService.GetTopicByID(c.Context(), req.Post.TopicId)
		if err != nil {
			return err
		}

		// send topic information to Workflow hook
		err = h.workflowHook.SendTopic(c.Context(), external_hook.TopicReq{
			CreatorUsername: req.Post.Username,
			CreatorFullName: req.Post.Name,
			TopicID:         topic.ID,
			TopicTitle:      topic.Title,
			TopicContent:    req.Post.Raw,
			Views:           topic.Views,
			Likes:           topic.LikeCount,
			Event:           "topic_created",
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
		topic, err := h.topicService.GetTopicByID(c.Context(), req.Post.TopicId)
		if err != nil {
			return err
		}

		if topic.DeletedAt != nil {
			return c.SendString("topic is deleted")
		}

		// send topic information to Workflow hook
		err = h.workflowHook.SendTopic(c.Context(), external_hook.TopicReq{
			CreatorUsername: req.Post.Username,
			CreatorFullName: req.Post.Name,
			TopicID:         topic.ID,
			TopicTitle:      topic.Title,
			TopicContent:    req.Post.Raw,
			Views:           topic.Views,
			Likes:           topic.LikeCount,
			Event:           "topic_edited",
		})
		if err != nil {
			return fmt.Errorf("failed to send topic to workflow: %w", err)
		}

		return c.SendString("sent hook workflow")
	}

	return c.SendString("nothing change")
}
