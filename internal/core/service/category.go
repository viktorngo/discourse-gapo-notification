package service

import (
	"context"
	"discourse-notification/internal/core/port"
	"discourse-notification/model"
)

type CategoryService struct {
	CategoryRepo port.CategoryRepository
}

func (service CategoryService) GetCategoryByID(ctx context.Context, id int) (*model.Category, error) {
	category, err := service.CategoryRepo.GetCategoryByID(id)
	if err != nil {
		return nil, err
	}

	return category, nil
}
