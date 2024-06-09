package model

import "time"

type Topic struct {
	ID                     int        `gorm:"column:id;primaryKey;autoIncrement"`
	Title                  string     `gorm:"column:title;not null"`
	LastPostedAt           time.Time  `gorm:"column:last_posted_at"`
	CreatedAt              time.Time  `gorm:"column:created_at;not null"`
	UpdatedAt              time.Time  `gorm:"column:updated_at;not null"`
	Views                  uint       `gorm:"column:views;not null;default:0"`
	PostsCount             uint       `gorm:"column:posts_count;not null;default:0"`
	UserID                 uint       `gorm:"column:user_id"`
	LastPostUserID         uint       `gorm:"column:last_post_user_id;not null"`
	ReplyCount             uint       `gorm:"column:reply_count;not null;default:0"`
	FeaturedUser1ID        uint       `gorm:"column:featured_user1_id"`
	FeaturedUser2ID        uint       `gorm:"column:featured_user2_id"`
	FeaturedUser3ID        uint       `gorm:"column:featured_user3_id"`
	DeletedAt              *time.Time `gorm:"column:deleted_at"`
	HighestPostNumber      uint       `gorm:"column:highest_post_number;not null;default:0"`
	LikeCount              uint       `gorm:"column:like_count;not null;default:0"`
	IncomingLinkCount      uint       `gorm:"column:incoming_link_count;not null;default:0"`
	CategoryID             int        `gorm:"column:category_id"`
	Visible                bool       `gorm:"column:visible;not null;default:true"`
	ModeratorPostsCount    uint       `gorm:"column:moderator_posts_count;not null;default:0"`
	Closed                 bool       `gorm:"column:closed;not null;default:false"`
	Archived               bool       `gorm:"column:archived;not null;default:false"`
	BumpedAt               time.Time  `gorm:"column:bumped_at;not null"`
	HasSummary             bool       `gorm:"column:has_summary;not null;default:false"`
	Archetype              string     `gorm:"column:archetype;not null;default:regular"`
	FeaturedUser4ID        uint       `gorm:"column:featured_user4_id"`
	NotifyModeratorsCount  uint       `gorm:"column:notify_moderators_count;not null;default:0"`
	SpamCount              uint       `gorm:"column:spam_count;not null;default:0"`
	PinnedAt               *time.Time `gorm:"column:pinned_at"`
	Score                  float64    `gorm:"column:score"`
	PercentRank            float64    `gorm:"column:percent_rank;not null;default:1.0"`
	Subtype                string     `gorm:"column:subtype"`
	Slug                   string     `gorm:"column:slug"`
	DeletedByUserID        uint       `gorm:"column:deleted_by_id"`
	ParticipantCount       int        `gorm:"column:participant_count;default:1"`
	WordCount              uint       `gorm:"column:word_count"`
	Excerpt                string     `gorm:"column:excerpt"`
	PinnedGlobally         bool       `gorm:"column:pinned_globally;not null;default:false"`
	PinnedUntil            *time.Time `gorm:"column:pinned_until"`
	FancyTitle             string     `gorm:"column:fancy_title"`
	HighestStaffPostNumber uint       `gorm:"column:highest_staff_post_number;not null;default:0"`
	FeaturedLink           string     `gorm:"column:featured_link"`
	ReviewableScore        float64    `gorm:"column:reviewable_score;not null;default:0.0"`
	ImageUploadID          int64      `gorm:"column:image_upload_id"`
	SlowModeSeconds        uint       `gorm:"column:slow_mode_seconds;not null;default:0"`
	BanneredUntil          *time.Time `gorm:"column:bannered_until"`
	ExternalID             string     `gorm:"column:external_id"`
}
