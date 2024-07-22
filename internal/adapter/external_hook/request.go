package external_hook

type TopicReq struct {
	CreatorUsername string `json:"creator_username"`
	CreatorFullName string `json:"creator_full_name"`
	TopicID         int    `json:"topic_id"`
	TopicTitle      string `json:"topic_title"`
	TopicContent    string `json:"topic_content"`
	Views           uint   `json:"views"`
	Likes           uint   `json:"likes"`
	Event           string `json:"event"`
	CategoryID      int    `json:"category_id"`
	CategoryName    string `json:"category_name"`
}
