package main

import (
	"net/http"

	_ "tistory/codes/code210309/docs"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// User struct
type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// UserMap type
type UserMap map[string]User

var (
	userMap UserMap
)

// @title Wookiist Sample Swagger API
// @version 1.0
// @host localhost:30000
// @BasePath /api/v1
func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/api/v1/user/:name", getUser)
	e.POST("/api/v1/user", createUser)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":30000"))
}

// @Summary Get user
// @Description Get user's info
// @Accept json
// @Produce json
// @Param name path string true "name of the user"
// @Success 200 {object} User
// @Router /user/{name} [get]
func getUser(c echo.Context) error {
	userName := c.Param("name")
	if val, ok := userMap[userName]; ok {
		return c.JSONPretty(http.StatusOK, val, "  ")
	}
	defaultUser := &User{
		Name: "default",
		Age:  0,
	}
	return c.JSONPretty(http.StatusBadRequest, defaultUser, "  ")
}

// @Summary Create user
// @Description Create new user
// @Accept json
// @Produce json
// @Param userBody body User true "User Info Body"
// @Success 200 {object} User
// @Router /user [post]
func createUser(c echo.Context) error {
	user := new(User)
	if err := c.Bind(&user); err != nil {
		return err
	}
	userMap[user.Name] = *user
	return c.JSONPretty(http.StatusOK, userMap[user.Name], "  ")
}
