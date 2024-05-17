package repository

import (
	"github.com/jordanmarcelino/backend-pplbo/internal/entity"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserRepository[T entity.User] struct {
	DB  *gorm.DB
	Log *logrus.Logger
}

func NewUserRepository(DB *gorm.DB, log *logrus.Logger) *UserRepository[entity.User] {
	return &UserRepository[entity.User]{DB: DB, Log: log}
}

func (u *UserRepository[T]) Create(db *gorm.DB, entity *T) error {
	return db.Create(entity).Error
}

func (u *UserRepository[T]) Update(db *gorm.DB, entity *T) error {
	return db.Save(entity).Error
}

func (u *UserRepository[T]) Delete(db *gorm.DB, entity *T) error {
	return db.Delete(entity).Error
}

func (u *UserRepository[T]) FindById(db *gorm.DB, entity *T, id any) error {
	return db.Take(entity, "id = ?", id).Error
}

func (r *UserRepository[T]) CountById(db *gorm.DB, id any) (int64, error) {
	var total int64
	err := db.Model(new(T)).Where("id = ?", id).Count(&total).Error
	return total, err
}

func (r *UserRepository[T]) CountByEmail(db *gorm.DB, email string) (int64, error) {
	var total int64
	err := db.Model(new(entity.User)).Where("email = ?", email).Count(&total).Error

	return total, err
}

func (r *UserRepository[T]) FindByEmail(db *gorm.DB, user *entity.User, email string) error {
	return db.Where("email = ?", email).Take(user).Error
}
