package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var users = map[int]*User{}

// ------------------ Controller --------------------

// get all user
func GetUsersConstroller(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users":    users,
	})
}

// get user by id
func GetUserConstroller(c echo.Context) error {
	user, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get user by id",
		"userId":   users[user],
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	user, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	delete(users, user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "delete success",
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	user := new(User)
	if err := c.Bind(user); err != nil {
		return err
	}
	if users[userId].Name != "" {
		users[userId].Name = user.Name
	}
	if users[userId].Email != "" {
		users[userId].Email = user.Email
	}
	if users[userId].Password != "" {
		users[userId].Password = user.Password
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update user",
		"users":   users[userId],
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	// binding dta
	user := User{}
	c.Bind(&user)

	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}
	users[user.Id] = &user
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"user":     user,
	})
}

func main() {
	e := echo.New()
	// routing with query parameter
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world!")
	})
	e.GET("/users", GetUsersConstroller)
	e.GET("/users/:id", GetUserConstroller)
	e.POST("/users", CreateUserController)
	e.PUT("/users/:id", UpdateUserController)
	e.DELETE("/users/:id", DeleteUserController)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
