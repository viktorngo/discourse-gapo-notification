package port

type Notification interface {
	SendMentionNotification(receiverUsername string, redirectURL string) error
	SendTopicCreatedNotification(receiverUsername string, topicTitle string, createdByName string, categoryName string, redirectUrl string) error
}
