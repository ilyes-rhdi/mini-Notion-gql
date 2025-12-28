package services

import (
	"context"
	"errors"

	"github.com/ilyes-rhdi/buildit-Gql/internal/models"
	"github.com/ilyes-rhdi/buildit-Gql/pkg/utils"
	"gorm.io/gorm"
)

type AuthService struct{}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (s *AuthService) CreateUser(name string, email string, password string, gender *bool) (*models.User, error) {
	ctx := context.Background()

	encryptedPassword, err := utils.Encrypt(password)
	if err != nil {
		return nil, err
	}

	u := &models.User{
		Name:     name,
		Email:    email,
		Password: encryptedPassword,
		Gender:   gender,

		Bio:   "",
		Image: "uploads/profiles/default.jpg",
		BgImg: "uploads/bgs/default.jpg",
		// Active false par défaut si ton model le définit
	}

	if err := getDB().WithContext(ctx).Create(u).Error; err != nil {
		return nil, err
	}

	return u, nil
}

func (s *AuthService) CheckUser(email string, password string) (*models.User, error) {
	ctx := context.Background()

	var u models.User
	if err := getDB().WithContext(ctx).Where("email = ?", email).First(&u).Error; err != nil {
		return nil, errors.New("wrong credentials")
	}

	if err := utils.CheckPassword(u.Password, password); err != nil {
		return nil, errors.New("wrong credentials")
	}

	return &u, nil
}

func (s *AuthService) GetUserByEmail(email string) (*models.User, error) {
	ctx := context.Background()

	var u models.User
	err := getDB().WithContext(ctx).Where("email = ?", email).First(&u).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &u, nil
}

func (s *AuthService) ActivateUser(userID string) error {
	ctx := context.Background()

	res := getDB().WithContext(ctx).Model(&models.User{}).
		Where("id = ?", userID).
		Update("active", true)

	return res.Error
}
