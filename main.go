package main

import (
	"discourse-notification/gapo"
	"discourse-notification/internal/hooks"
	"discourse-notification/internal/repositories"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
)

type DiscourseNotification struct {
	Id               int       `json:"id"`
	UserId           int       `json:"user_id"`
	NotificationType int       `json:"notification_type"`
	Read             bool      `json:"read"`
	HighPriority     bool      `json:"high_priority"`
	CreatedAt        time.Time `json:"created_at"`
	PostNumber       int       `json:"post_number"`
	TopicId          int       `json:"topic_id"`
	FancyTitle       string    `json:"fancy_title"`
	Slug             string    `json:"slug"`
	Data             struct {
		TopicTitle       string      `json:"topic_title"`
		OriginalPostId   int         `json:"original_post_id"`
		OriginalPostType int         `json:"original_post_type"`
		OriginalUsername string      `json:"original_username"`
		RevisionNumber   interface{} `json:"revision_number"`
		DisplayUsername  string      `json:"display_username"`
	} `json:"data"`
}

type DiscourseUser struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (DiscourseUser) TableName() string {
	return "users"
}

func main() {
	// load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	discourseHost := os.Getenv("DISCOURSE_HOST")

	// connect DB
	dsn := os.Getenv("DATABASE_DSN")
	if dsn == "" {
		dsn = "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Bangkok"
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// create Gapo client
	host := os.Getenv("GAPO_HOST")
	botID := os.Getenv("GAPO_BOT_ID")
	botToken := os.Getenv("GAPO_BOT_TOKEN")
	apiKey := os.Getenv("GAPO_API_KEY")

	gapoClient := gapo.Client{
		Host:     host,
		BotID:    botID,
		BotToken: botToken,
		ApiKey:   apiKey,
	}

	// repositories
	topicRepository := repositories.NewTopicRepository(db)
	userRepository := repositories.NewUserRepository(db)

	// services
	topicHooker := &hooks.TopicHooker{
		DiscourseHost:   discourseHost,
		GapoWorkClient:  &gapoClient,
		TopicRepository: topicRepository,
		UserRepository:  userRepository,
	}

	// create fiber app
	app := fiber.New()

	app.Get("/hooks/discourse", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/hooks/discourse", func(c *fiber.Ctx) error {
		var req struct {
			Notification DiscourseNotification `json:"notification"`
		}
		if err := c.BodyParser(&req); err != nil {
			log.Errorf("failed to parse request: %v", err)
			return err
		}
		// event user mention
		if req.Notification.NotificationType == 1 {
			var user DiscourseUser
			if err := db.Where("id = ?", req.Notification.UserId).First(&user).Error; err != nil {
				log.Errorf("failed to query user: %v", err)
				return err
			}
			log.Infof("User %s has been mentioned", user.Username)

			message := `**📢 Thông báo: Góp ý cần trả lời!**
Bạn đã được nhắc để trả lời một góp ý quan trọng. Hãy nhanh chóng truy cập vào YOKAIZEN để đưa ra ý kiến của mình:
[Trả lời ngay](%s)
`
			redirectUrl := fmt.Sprintf("%s/t/%s/%d/%d", discourseHost, req.Notification.Slug, req.Notification.TopicId, req.Notification.PostNumber)
			if err := gapoClient.SendNotification(user.Username, fmt.Sprintf(message, redirectUrl)); err != nil {
				log.Errorf("failed to send Gapo notification for `user mension`: %v", err)
				return err
			}
		}
		return c.SendString("Success!")
	})

	app.Post("/hooks/discourse/topic-created", func(c *fiber.Ctx) error {
		var req struct {
			Topic struct {
				ID         uint64 `json:"id"`
				Title      string `json:"title"`
				Slug       string `json:"slug"`
				CategoryID uint64 `json:"category_id"`
				CreatedBy  struct {
					Id       uint64 `json:"id"`
					Username string `json:"username"`
					Name     string `json:"name"`
				} `json:"created_by"`
			} `json:"topic"`
		}
		if err := c.BodyParser(&req); err != nil {
			log.Errorf("failed to parse request: %v", err)
			return err
		}

		if err := topicHooker.TopicCreated(req.Topic.CategoryID, req.Topic.ID, req.Topic.Slug, req.Topic.Title, req.Topic.CreatedBy.Id, req.Topic.CreatedBy.Name); err != nil {
			log.Errorf("hook topic created failed: %v", err)
			return err
		}

		return c.SendString("Success!")
	})

	app.Listen("0.0.0.0:8000")
}
