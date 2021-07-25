package databases

import (
	"part1/config"
	"part1/models"

	"github.com/labstack/echo"
)

func GetUsers() (interface{}, error) {
	var books []models.Book
	if err := config.DB.Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func GetUser(id int) (interface{}, error) {
	var book models.Book
	if err := config.DB.Find(&book, "id=?", id).Error; err != nil {
		return nil, err
	}
	return book, nil
}

func UpdateUser(id int, book interface{}) (interface{}, error) {
	if err := config.DB.Find(&book, "id=?", id).Save(&book).Error; err != nil {
		return nil, err
	}
	return book, nil
}

func DeleteUser(id int) (interface{}, error) {
	var book models.Book
	if err := config.DB.Find(&book, "id=?", id).Delete(&book).Error; err != nil {
		return nil, err
	}
	return book, nil
}

func CreateUser(c echo.Context) (interface{}, error) {
	book := models.Book{}
	c.Bind(&book)
	if err := config.DB.Save(&book).Error; err != nil {
		return nil, err
	}
	return book, nil
}
