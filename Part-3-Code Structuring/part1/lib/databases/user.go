package databases

import (
	"part1/config"
	"part1/models"

	"github.com/labstack/echo"
)

func GetUsers() (interface{}, error) {
	var users []models.User
	if err := config.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func GetUser(id int) (interface{}, error) {
	var user models.User
	if err := config.DB.Find(&user, "id=?", id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func UpdateUser(id int, user interface{}) (interface{}, error) {
	if err := config.DB.Find(&user, "id=?", id).Save(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func DeleteUser(id int) (interface{}, error) {
	var user models.User
	if err := config.DB.Find(&user, "id=?", id).Delete(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func CreateUser(c echo.Context) (interface{}, error) {
	user := models.User{}
	c.Bind(&user)
	if err := config.DB.Save(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
