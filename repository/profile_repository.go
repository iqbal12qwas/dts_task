package repository

import (
	"gin-api/models"
	"gin-api/helper"
)

func AllProfiles() ([]models.Profile, error) {
	var profiles []models.Profile 
	if err := models.SetupDB().Find(&profiles).Error; err != nil {
		return profiles, err
	}
	return profiles, nil
}

func GetAllProfiles(profile *models.Profile, pagination *helper.Pagination) (*[]models.Profile, error) {
	var profiles []models.Profile
	offset := (pagination.Page - 1) * pagination.Limit
	queryBuider := models.SetupDB().Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)
	result := queryBuider.Model(&models.Profile{}).Where(profile).Where("name like ?", "%"+pagination.Search+"%").Find(&profiles)
	if result.Error != nil {
		msg := result.Error
		return nil, msg
	}
	return &profiles, nil
}

func FindProfileById(id string) (models.Profile, error) {
	var profile models.Profile 
	if err := models.SetupDB().Where("id = ?", id).First(&profile).Error; err != nil {
		return profile, err
	}
	return profile, nil
}

func UpdateProfile(profile models.Profile, updateProfile models.Profile) (models.Profile, error) {
	if err := models.SetupDB().Model(&profile).Updates(updateProfile).Error; err != nil {
		return updateProfile, err
	}
	return updateProfile, nil
}

func DeleteProfile(profile models.Profile) (error){
	if err := models.SetupDB().Delete(profile).Error; err != nil {
		return err
	}
	return nil
}