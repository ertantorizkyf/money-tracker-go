package usecases

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/ertantorizkyf/money-tracker-go/constants"
	"github.com/ertantorizkyf/money-tracker-go/dto"
	"github.com/ertantorizkyf/money-tracker-go/helpers"
	"github.com/ertantorizkyf/money-tracker-go/initializers"
	"github.com/ertantorizkyf/money-tracker-go/models"
	"github.com/ertantorizkyf/money-tracker-go/repositories"
	"github.com/gin-gonic/gin"
	"github.com/willf/bloom"
	"gorm.io/gorm"
)

type UserUseCase struct {
	BloomFilter *bloom.BloomFilter
	UserRepo    *repositories.UserRepository
}

func NewUserUsecase(userRepo *repositories.UserRepository) *UserUseCase {
	return &UserUseCase{
		BloomFilter: initializers.BloomFilter,
		UserRepo:    userRepo,
	}
}

func (uc *UserUseCase) RegisterUser(c *gin.Context, req dto.RegisterReq) (*string, error) {
	// VALIDATE PASSWORD HASH WITH BLOOM FILTER
	sha1HashedPassword := sha1.Sum([]byte(req.Password))
	if uc.BloomFilter.Test(sha1HashedPassword[:]) {
		helpers.LogWithSeverity(constants.LOGGER_SEVERITY_ERROR, fmt.Errorf("an error has occurred: %s", constants.ERR_MESSAGE_BAD_REQUEST))
		return nil, fmt.Errorf("an error has occurred: %s", constants.ERR_MESSAGE_BAD_REQUEST)
	}

	user, err := uc.UserRepo.GetFirst(models.UserWhere{
		Username: req.Username,
		OrEmail:  req.Email,
		OrPhone:  req.Phone,
	})

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		helpers.LogWithSeverity(constants.LOGGER_SEVERITY_ERROR, err)
		return nil, err
	}

	if user.ID > 0 {
		helpers.LogWithSeverity(constants.LOGGER_SEVERITY_ERROR, constants.ERR_MESSAGE_DATA_TAKEN)
		return nil, fmt.Errorf("an error has occurred: %s", constants.ERR_MESSAGE_DATA_TAKEN)
	}

	hashedPassword, err := helpers.HashPassword(req.Password)
	if err != nil {
		helpers.LogWithSeverity(constants.LOGGER_SEVERITY_ERROR, err)
		return nil, err
	}

	timeParsedDOB, err := time.Parse("2006-01-02", req.DOB)
	if err != nil {
		helpers.LogWithSeverity(constants.LOGGER_SEVERITY_ERROR, err)
		return nil, err
	}

	newUser := models.User{
		Name:     req.Name,
		Username: req.Username,
		Email:    req.Email,
		Phone:    req.Phone,
		Password: hashedPassword,
		DOB:      timeParsedDOB,
	}

	err = uc.UserRepo.Create(newUser)
	if err != nil {
		helpers.LogWithSeverity(constants.LOGGER_SEVERITY_ERROR, err)
		return nil, err
	}

	jwtToken, err := helpers.GenerateToken(newUser)
	if err != nil {
		helpers.LogWithSeverity(constants.LOGGER_SEVERITY_ERROR, err)
		return nil, err
	}

	return &jwtToken, nil
}

func (uc *UserUseCase) Login(c *gin.Context, req dto.LoginReq) (*string, error) {
	user, err := uc.UserRepo.GetFirst(models.UserWhere{
		Username: req.UsernameOrEmail,
		OrEmail:  req.UsernameOrEmail,
	})

	if err != nil {
		helpers.LogWithSeverity(constants.LOGGER_SEVERITY_ERROR, err)
		return nil, err
	}

	isPasswordMatch := helpers.CheckPasswordHash(req.Password, user.Password)
	if !isPasswordMatch {
		return nil, fmt.Errorf("an error has occurred: %s", constants.ERR_MESSAGE_INVALID_CREDENTIALS)
	}

	jwtToken, err := helpers.GenerateToken(user)
	if err != nil {
		helpers.LogWithSeverity(constants.LOGGER_SEVERITY_ERROR, err)
		return nil, err
	}

	return &jwtToken, nil
}
