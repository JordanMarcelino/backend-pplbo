package converter

import (
	"github.com/jordanmarcelino/backend-pplbo/internal/entity"
	"github.com/jordanmarcelino/backend-pplbo/internal/models"
)

func UserToResponse(user *entity.User) *models.UserResponse {
	return &models.UserResponse{
		Id:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Unix(),
	}
}
