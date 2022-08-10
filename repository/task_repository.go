package repository

import (
	"gin-api/models"
	"gin-api/helper"
)

func GetAllTasks(task *models.Task, pagination *helper.Pagination) (*[]models.Task, error) {
	var tasks []models.Task
	offset := (pagination.Page - 1) * pagination.Limit
	queryBuider := models.SetupDB().Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)
	result := queryBuider.Model(&models.Task{}).Where(task).Where("assigned_to like ?", "%"+pagination.Search+"%").Find(&tasks)
	if result.Error != nil {
		msg := result.Error
		return nil, msg
	}
	return &tasks, nil
}

func AllTasks() ([]models.Profile, error) {
	var profiles []models.Profile 
	if err := models.SetupDB().Find(&profiles).Error; err != nil {
		return profiles, err
	}
	return profiles, nil
}

func FindTaskById(id string) (models.Task, error) {
	var task models.Task 
	if err := models.SetupDB().Where("id = ?", id).First(&task).Error; err != nil {
		return task, err
	}
	return task, nil
}

func UpdateTask(task models.Task, updateTask models.Task) (models.Task, error) {
	if err := models.SetupDB().Model(&task).Updates(updateTask).Error; err != nil {
		return updateTask, err
	}
	return updateTask, nil
}

func DeleteTask(task models.Task) (error){
	if err := models.SetupDB().Delete(task).Error; err != nil {
		return err
	}
	return nil
}

func DeleteTaskByIdProfile(id string) (error){
	// Query Builder
	if err := models.SetupDB().Exec("DELETE FROM tasks WHERE id_profile = ?", id).Error; err != nil {
		return err
	}
	return nil
}