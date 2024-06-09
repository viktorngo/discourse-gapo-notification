package model

import "time"

type Post struct {
	ID                    int32      `json:"id"`
	UserID                int32      `json:"user_id"`
	TopicID               int32      `json:"topic_id"`
	PostNumber            int32      `json:"post_number"`
	Raw                   string     `json:"raw"`
	Cooked                string     `json:"cooked"`
	CreatedAt             time.Time  `json:"created_at"`
	UpdatedAt             time.Time  `json:"updated_at"`
	ReplyToPostNumber     *int32     `json:"reply_to_post_number,omitempty"`
	ReplyCount            int32      `json:"reply_count"`
	QuoteCount            int32      `json:"quote_count"`
	DeletedAt             *time.Time `json:"deleted_at,omitempty"`
	OffTopicCount         int32      `json:"off_topic_count"`
	LikeCount             int32      `json:"like_count"`
	IncomingLinkCount     int32      `json:"incoming_link_count"`
	BookmarkCount         int32      `json:"bookmark_count"`
	Score                 float64    `json:"score"`
	Reads                 int32      `json:"reads"`
	PostType              int32      `json:"post_type"`
	SortOrder             *int32     `json:"sort_order,omitempty"`
	LastEditorID          *int32     `json:"last_editor_id,omitempty"`
	Hidden                bool       `json:"hidden"`
	HiddenReasonID        *int32     `json:"hidden_reason_id,omitempty"`
	NotifyModeratorsCount int32      `json:"notify_moderators_count"`
	SpamCount             int32      `json:"spam_count"`
	IllegalCount          int32      `json:"illegal_count"`
	InappropriateCount    int32      `json:"inappropriate_count"`
	LastVersionAt         time.Time  `json:"last_version_at"`
	UserDeleted           bool       `json:"user_deleted"`
	ReplyToUserID         *int32     `json:"reply_to_user_id,omitempty"`
	PercentRank           float64    `json:"percent_rank"`
	NotifyUserCount       int32      `json:"notify_user_count"`
	LikeScore             int32      `json:"like_score"`
	DeletedByID           *int32     `json:"deleted_by_id,omitempty"`
	EditReason            string     `json:"edit_reason,omitempty"`
	WordCount             int32      `json:"word_count"`
	Version               int32      `json:"version"`
	CookMethod            int32      `json:"cook_method"`
	Wiki                  bool       `json:"wiki"`
	BakedAt               *time.Time `json:"baked_at,omitempty"`
	BakedVersion          *int32     `json:"baked_version,omitempty"`
	HiddenAt              *time.Time `json:"hidden_at,omitempty"`
	SelfEdits             int32      `json:"self_edits"`
	ReplyQuoted           bool       `json:"reply_quoted"`
	ViaEmail              bool       `json:"via_email"`
	RawEmail              string     `json:"raw_email,omitempty"`
	PublicVersion         int32      `json:"public_version"`
	ActionCode            string     `json:"action_code,omitempty"`
	LockedByID            *int32     `json:"locked_by_id,omitempty"`
	ImageUploadID         *int64     `json:"image_upload_id,omitempty"`
	OutboundMessageID     string     `json:"outbound_message_id,omitempty"`
}
