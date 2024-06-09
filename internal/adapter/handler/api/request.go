package api

import "time"

type Topic struct {
	Tags             []interface{} `json:"tags"`
	TagsDescriptions struct {
	} `json:"tags_descriptions"`
	Id                int         `json:"id"`
	Title             string      `json:"title"`
	FancyTitle        string      `json:"fancy_title"`
	PostsCount        int         `json:"posts_count"`
	CreatedAt         time.Time   `json:"created_at"`
	Views             int         `json:"views"`
	ReplyCount        int         `json:"reply_count"`
	LikeCount         int         `json:"like_count"`
	LastPostedAt      time.Time   `json:"last_posted_at"`
	Visible           bool        `json:"visible"`
	Closed            bool        `json:"closed"`
	Archived          bool        `json:"archived"`
	Archetype         string      `json:"archetype"`
	Slug              string      `json:"slug"`
	CategoryId        int         `json:"category_id"`
	WordCount         int         `json:"word_count"`
	DeletedAt         interface{} `json:"deleted_at"`
	UserId            int         `json:"user_id"`
	FeaturedLink      interface{} `json:"featured_link"`
	PinnedGlobally    bool        `json:"pinned_globally"`
	PinnedAt          interface{} `json:"pinned_at"`
	PinnedUntil       interface{} `json:"pinned_until"`
	Unpinned          interface{} `json:"unpinned"`
	Pinned            bool        `json:"pinned"`
	HighestPostNumber int         `json:"highest_post_number"`
	DeletedBy         interface{} `json:"deleted_by"`
	HasDeleted        bool        `json:"has_deleted"`
	Bookmarked        bool        `json:"bookmarked"`
	ParticipantCount  int         `json:"participant_count"`
	Thumbnails        interface{} `json:"thumbnails"`
	CreatedBy         User        `json:"created_by"`
	LastPoster        User        `json:"last_poster"`
}
type Post struct {
	Id                          int         `json:"id"`
	Name                        string      `json:"name"`
	Username                    string      `json:"username"`
	AvatarTemplate              string      `json:"avatar_template"`
	CreatedAt                   time.Time   `json:"created_at"`
	Cooked                      string      `json:"cooked"`
	PostNumber                  int         `json:"post_number"`
	PostType                    int         `json:"post_type"`
	UpdatedAt                   time.Time   `json:"updated_at"`
	ReplyCount                  int         `json:"reply_count"`
	ReplyToPostNumber           interface{} `json:"reply_to_post_number"`
	QuoteCount                  int         `json:"quote_count"`
	IncomingLinkCount           int         `json:"incoming_link_count"`
	Reads                       int         `json:"reads"`
	Score                       int         `json:"score"`
	TopicId                     int         `json:"topic_id"`
	TopicSlug                   string      `json:"topic_slug"`
	TopicTitle                  string      `json:"topic_title"`
	CategoryId                  int         `json:"category_id"`
	DisplayUsername             string      `json:"display_username"`
	PrimaryGroupName            interface{} `json:"primary_group_name"`
	FlairName                   interface{} `json:"flair_name"`
	FlairGroupId                interface{} `json:"flair_group_id"`
	Version                     int         `json:"version"`
	UserTitle                   interface{} `json:"user_title"`
	Bookmarked                  bool        `json:"bookmarked"`
	Raw                         string      `json:"raw"`
	Moderator                   bool        `json:"moderator"`
	Admin                       bool        `json:"admin"`
	Staff                       bool        `json:"staff"`
	UserId                      int         `json:"user_id"`
	Hidden                      bool        `json:"hidden"`
	TrustLevel                  int         `json:"trust_level"`
	DeletedAt                   interface{} `json:"deleted_at"`
	UserDeleted                 bool        `json:"user_deleted"`
	EditReason                  interface{} `json:"edit_reason"`
	Wiki                        bool        `json:"wiki"`
	ReviewableId                interface{} `json:"reviewable_id"`
	ReviewableScoreCount        int         `json:"reviewable_score_count"`
	ReviewableScorePendingCount int         `json:"reviewable_score_pending_count"`
	TopicPostsCount             int         `json:"topic_posts_count"`
	TopicFilteredPostsCount     int         `json:"topic_filtered_posts_count"`
	TopicArchetype              string      `json:"topic_archetype"`
	CategorySlug                string      `json:"category_slug"`
}

type User struct {
	Id             int    `json:"id"`
	Username       string `json:"username"`
	Name           string `json:"name"`
	AvatarTemplate string `json:"avatar_template"`
}

type PostLikeReq struct {
	Like struct {
		Post Post `json:"post"`
		User User `json:"user"`
	} `json:"like"`
}

type TopicCreatedReq struct {
	Topic Topic `json:"topic"`
}

type PostCreatedReq struct {
	Post Post `json:"post"`
}
type PostEditedReq struct {
	Post Post `json:"post"`
}
