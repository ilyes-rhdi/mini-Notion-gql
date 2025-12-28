package services

import (
	"context"
	"errors"

	"github.com/ilyes-rhdi/buildit-Gql/internal/models"
	"github.com/ilyes-rhdi/buildit-Gql/pkg/logger"
	"github.com/ilyes-rhdi/buildit-Gql/pkg/types"
	"gorm.io/gorm"
)

type ProfileService struct{}

func NewProfileService() *ProfileService {
	return &ProfileService{}
}

func (s *ProfileService) GetUser(id string) (*models.User, error) {
	logger.LogInfo().Fields(map[string]interface{}{
		"query":  "get profile",
		"params": id,
	}).Msg("DB Query")

	ctx := context.Background()

	var user models.User
	if err := getDB().WithContext(ctx).
		Omit("password").
		First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *ProfileService) GetUserByEmail(email string) (*models.User, error) {
	logger.LogInfo().Fields(map[string]interface{}{
		"query":  "search profile",
		"params": email,
	}).Msg("DB Query")

	ctx := context.Background()

	var user models.User
	if err := getDB().WithContext(ctx).
		Omit("password").
		Where("email = ?", email).
		First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *ProfileService) SearchByName(name string) ([]models.User, error) {
	logger.LogInfo().Fields(map[string]interface{}{
		"query":  "search profile",
		"params": name,
	}).Msg("DB Query")

	ctx := context.Background()

	var users []models.User
	// Si tu es sur Postgres et tu veux case-insensitive, remplace LIKE par ILIKE.
	if err := getDB().WithContext(ctx).
		Omit("password").
		Where("name LIKE ?", "%"+name+"%").
		Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (s *ProfileService) UpdateUser(id string, payload types.ProfileUpdate) (*models.User, error) {
	logger.LogInfo().Fields(map[string]interface{}{
		"query":  "update profile",
		"id":     id,
		"params": payload,
	}).Msg("DB Query")

	ctx := context.Background()

	updates := map[string]any{
		"email":          payload.Email,
		"name":           payload.Name,
		"bio":            payload.Bio,
		"adress":         payload.Adress,
		"phone":          payload.Phone,
		"external_links": payload.Links,
	}

	if err := getDB().WithContext(ctx).
		Model(&models.User{}).
		Where("id = ?", id).
		Updates(updates).Error; err != nil {
		return nil, err
	}

	return s.GetUser(id)
}

func (s *ProfileService) UpdateUserImage(id string, path string) (string, error) {
	ctx := context.Background()

	res := getDB().WithContext(ctx).
		Model(&models.User{}).
		Where("id = ?", id).
		Update("image", path)
	if res.Error != nil {
		return "", res.Error
	}
	if res.RowsAffected == 0 {
		return "", gorm.ErrRecordNotFound
	}

	return path, nil
}

func (s *ProfileService) UpdateUserBg(id string, path string) (string, error) {
	ctx := context.Background()

	res := getDB().WithContext(ctx).
		Model(&models.User{}).
		Where("id = ?", id).
		Update("bg_img", path)
	if res.Error != nil {
		return "", res.Error
	}
	if res.RowsAffected == 0 {
		return "", gorm.ErrRecordNotFound
	}

	return path, nil
}

func (s *ProfileService) DeleteUser(id string) (string, error) {
	logger.LogInfo().Fields(map[string]interface{}{
		"query":  "delete profile",
		"params": id,
	}).Msg("DB Query")

	ctx := context.Background()

	res := getDB().WithContext(ctx).Delete(&models.User{}, "id = ?", id)
	if res.Error != nil {
		return "", res.Error
	}
	if res.RowsAffected == 0 {
		return "", errors.New("user not found")
	}

	return id, nil
}
