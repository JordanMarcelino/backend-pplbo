package usecase

import (
	"context"
	"github.com/jordanmarcelino/backend-pplbo/internal/entity"
	"github.com/jordanmarcelino/backend-pplbo/internal/models"
	"github.com/jordanmarcelino/backend-pplbo/internal/models/converter"
	"github.com/jordanmarcelino/backend-pplbo/internal/repository"
	"github.com/jordanmarcelino/backend-pplbo/internal/util"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserUseCase struct {
	DB             *gorm.DB
	Log            *logrus.Logger
	UserRepository *repository.UserRepository[entity.User]
}

func NewUserUseCase(DB *gorm.DB, log *logrus.Logger, userRepository *repository.UserRepository[entity.User]) *UserUseCase {
	return &UserUseCase{DB: DB, Log: log, UserRepository: userRepository}
}

func (c *UserUseCase) Login(ctx context.Context, request *models.UserLogin) error {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	total, err := c.UserRepository.CountByEmail(tx, request.Email)
	if err != nil {
		c.Log.Warnf("failed to count user from database : %+v", err)
		return util.InternalServerError
	}

	if total == 0 {
		c.Log.Warnf("user doesn't exist : %+v", err)
		return util.NotFound
	}

	user := new(entity.User)
	if err := c.UserRepository.FindByEmail(tx, user, request.Email); err != nil {
		c.Log.Warnf("failed to get user : %+v", err)
		return util.InternalServerError
	}

	if authorize := util.CheckPasswordHash(request.Password, user.Password); !authorize {
		c.Log.Warnf("password doesn't match")
		return util.Unauthorized
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("failed commit transaction : %+v", err)
		return util.InternalServerError
	}

	return nil
}

func (c *UserUseCase) Create(ctx context.Context, request *models.UserRegister) (*models.UserResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	total, err := c.UserRepository.CountByEmail(tx, request.Email)
	if err != nil {
		c.Log.Warnf("failed to count user from database : %+v", err)
		return nil, util.InternalServerError
	}

	if total != 0 {
		c.Log.Warn("user already exist")
		return nil, util.BadRequest
	}

	password, err := util.HashPassword(request.Password)
	if err != nil {
		c.Log.Warnf("failed to hash password : %+v", err)
		return nil, util.InternalServerError
	}

	user := &entity.User{
		Username: request.Username,
		Email:    request.Email,
		Password: password,
	}

	if err := c.UserRepository.Create(tx, user); err != nil {
		c.Log.Warnf("failed create user to database : %+v", err)
		return nil, util.InternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("failed commit transaction : %+v", err)
		return nil, util.InternalServerError
	}

	return converter.UserToResponse(user), nil
}

func (c *UserUseCase) Update(ctx context.Context, request *models.UserUpdate) (*models.UserResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	total, err := c.UserRepository.CountById(tx, request.Id)

	if err != nil {
		c.Log.Warnf("failed to count user from database : %+v", err)
		return nil, util.InternalServerError
	}

	if total == 0 {
		c.Log.Warnf("user doesn't exist : %+v", err)
		return nil, util.NotFound
	}

	user := new(entity.User)

	c.UserRepository.FindById(tx, user, request.Id)

	if request.Password != "" {
		password, err := util.HashPassword(request.Password)
		if err != nil {
			c.Log.Warnf("failed to hash password : %+v", err)
			return nil, util.InternalServerError
		}
		user.Password = password
	}

	if request.Username != "" {
		user.Username = request.Username
	}

	if request.Email != "" {
		user.Email = request.Email
	}

	if err := c.UserRepository.Update(tx, user); err != nil {
		c.Log.Warnf("failed update user to database : %+v", err)
		return nil, util.InternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("failed commit transaction : %+v", err)
		return nil, util.InternalServerError
	}

	return converter.UserToResponse(user), nil
}

func (c *UserUseCase) FindById(ctx context.Context, id int) (*models.UserResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	user := new(entity.User)
	total, err := c.UserRepository.CountById(tx, id)

	if err != nil {
		c.Log.Warnf("failed to count user from database : %+v", err)
		return nil, util.InternalServerError
	}

	if total == 0 {
		c.Log.Warnf("user doesn't exist : %+v", err)
		return nil, util.InternalServerError
	}

	if err := c.UserRepository.FindById(tx, user, id); err != nil {
		c.Log.Warnf("failed find user from database : %+v", err)
		return nil, util.InternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("failed commit transaction : %+v", err)
		return nil, util.InternalServerError
	}

	return converter.UserToResponse(user), nil
}
