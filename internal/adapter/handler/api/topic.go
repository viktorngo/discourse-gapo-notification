package api

import (
	"discourse-notification/internal/core/service"
	"discourse-notification/model"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type TopicHandler struct {
	topicService *service.TopicService
}

func NewHandlerService(topicService *service.TopicService) TopicHandler {
	return TopicHandler{topicService: topicService}
}

func (h TopicHandler) TopicCreated(c *fiber.Ctx) error {
	var req TopicCreatedReq
	if err := c.BodyParser(&req); err != nil {
		log.Errorf("failed to parse request: %v", err)
		return err
	}

	err := h.topicService.SendNoticeTopicCreated(c.UserContext(),
		&model.Topic{
			ID:         req.Topic.Id,
			CategoryID: req.Topic.CategoryId,
			Slug:       req.Topic.Slug,
			Title:      req.Topic.Title,
		},
		&model.User{
			ID:   uint64(req.Topic.CreatedBy.Id),
			Name: req.Topic.CreatedBy.Name,
		},
	)
	if err != nil {
		return err
	}

	return c.SendString("Success!")
}
