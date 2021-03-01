package actions

import (
	"crud_rest_api/models"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gobuffalo/buffalo"
)

// HomeHandler is a default handler to serve up
// a home page.

type STReponse struct {
	Status string
	Data   models.User
}

func getDataFromBody(body io.ReadCloser) models.User {

	parsed, _ := ioutil.ReadAll(body)

	bodyJSON := models.User{}
	json.Unmarshal([]byte(parsed), &bodyJSON)

	return (bodyJSON)
}

func AddUser(c buffalo.Context) error {

	data := getDataFromBody(c.Request().Body)

	_, err := models.DB.ValidateAndCreate(&data)
	if err != nil {
		fmt.Println(err)
		return c.Render(http.StatusBadRequest, r.JSON(map[string]string{"error": "server error"}))
	}
	return c.Render(http.StatusCreated, r.JSON(STReponse{Status: "user has been created", Data: data}))

}

func GetAllUsers(c buffalo.Context) error {

	users := []models.User{}

	err := models.DB.All(&users)
	if err != nil {
		return c.Render(http.StatusBadRequest, r.JSON(map[string]string{"error": "server error"}))
	}

	return c.Render(http.StatusCreated, r.JSON(users))

}

func FindUserByEmail(c buffalo.Context) error {

	data := getDataFromBody(c.Request().Body)

	user := []models.User{}
	query := models.DB.Where("email = ?", data.Email)

	err := query.All(&user)
	if err != nil {
		fmt.Print(err)
		return c.Render(http.StatusBadRequest, r.JSON(map[string]string{"error": "server error"}))
	}

	return c.Render(http.StatusCreated, r.JSON(user))
}

func DeleteUserByEmail(c buffalo.Context) error {
	data := getDataFromBody(c.Request().Body)

	usersToDestroy := []models.User{}
	query := models.DB.Where("email = ?", data.Email)

	err := query.All(&usersToDestroy)
	if err != nil {
		fmt.Print(err)
		return c.Render(http.StatusBadRequest, r.JSON(map[string]string{"error": "server error"}))
	}

	err2 := models.DB.Destroy(&usersToDestroy)
	if err2 != nil {
		fmt.Print(err)
		return c.Render(http.StatusBadRequest, r.JSON(map[string]string{"error": "server error"}))
	}

	return c.Render(http.StatusCreated, r.JSON(usersToDestroy))
}

func UpdateUserByEmail(c buffalo.Context) error {
	data := getDataFromBody(c.Request().Body)

	usersToUpdate := []models.User{}
	query := models.DB.Where("email = ?", data.Email)

	err := query.All(&usersToUpdate)
	if err != nil {
		fmt.Print(err)
		return c.Render(http.StatusBadRequest, r.JSON(map[string]string{"error": "server error"}))
	}

	for i := 0; i < len(usersToUpdate); i++ {
		user := usersToUpdate[i]
		user.Name = data.Name
		models.DB.ValidateAndUpdate(&user)
	}

	usersToUpdate = []models.User{}

	err2 := query.All(&usersToUpdate)
	if err2 != nil {
		fmt.Print(err)
		return c.Render(http.StatusBadRequest, r.JSON(map[string]string{"error": "server error"}))
	}

	return c.Render(http.StatusCreated, r.JSON(usersToUpdate))
}
