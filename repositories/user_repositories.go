package repositories

import (
	"workshop/database"
	"workshop/models"
)

func FindAllUsers() ([]models.User, error) {
	var users []models.User
	err := database.DB.Order("id desc").Find(&users).Error
	return users, err
}

func FindUserByID(id uint) (*models.User, error) {
	var user models.User
	err := database.DB.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func FindUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := database.DB.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func CreateUser(user *models.User) error {
	return database.DB.Create(user).Error
}

func UpdateUser(user *models.User, updates map[string]any) error {
	return database.DB.Model(user).Updates(updates).Error
}

func DeleteUserByID(id uint) error {
	return database.DB.Delete(&models.User{}, id).Error
}
