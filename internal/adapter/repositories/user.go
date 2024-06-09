package repositories

import (
	"discourse-notification/internal/core/port"
	"discourse-notification/model"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func (u userRepository) GetUsersByGroupID(groupID uint64) ([]model.User, error) {
	var users []model.User
	if err := u.db.Model(&model.User{}).
		Joins("join group_users on group_users.user_id = users.id").
		Where("group_users.group_id = ?", groupID).
		Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func NewUserRepository(db *gorm.DB) port.UserRepository {
	return &userRepository{db: db}
}
