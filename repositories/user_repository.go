package repositories

import (
	"github.com/ertantorizkyf/money-tracker-go/constants"
	"github.com/ertantorizkyf/money-tracker-go/helpers"
	"github.com/ertantorizkyf/money-tracker-go/initializers"
	"github.com/ertantorizkyf/money-tracker-go/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		DB: initializers.DB,
	}
}

func (r *UserRepository) GetAll() ([]models.User, error) {
	var users []models.User

	err := r.DB.Find(&users).Error
	if err != nil {
		helpers.LogWithSeverity(constants.LOGGER_SEVERITY_ERROR, err)
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) GetFirst(where models.UserWhere) (models.User, error) {
	var user models.User

	query := r.DB

	if where.Email != "" {
		query.Where("email = ?", where.Email)
	}
	if where.Phone != "" {
		query.Where("phone = ?", where.Phone)
	}
	if where.Username != "" {
		query.Where("username = ?", where.Username)
	}

	err := query.Find(&user).Error
	if err != nil {
		helpers.LogWithSeverity(constants.LOGGER_SEVERITY_ERROR, err)
		return user, err
	}

	return user, nil
}

func (r *UserRepository) Create(user models.User) error {
	tx := r.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	err := tx.Create(&user).Error
	if err != nil {
		tx.Rollback()
		helpers.LogWithSeverity(constants.LOGGER_SEVERITY_ERROR, err)
		return err
	}

	if err = tx.Commit().Error; err != nil {
		helpers.LogWithSeverity(constants.LOGGER_SEVERITY_ERROR, err)
		return err
	}

	return nil
}
