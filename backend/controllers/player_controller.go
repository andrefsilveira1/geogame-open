package controllers

import (
	"fmt"
	db "geolocation/database"
	"geolocation/models"
	"geolocation/repositories"
	"geolocation/utils"
	"net/http"
	"strconv"

	"github.com/andrefsilveira1/go-haversine"
	"github.com/gofiber/fiber/v2"
	"github.com/mitchellh/mapstructure"
	"golang.org/x/crypto/bcrypt"
)

func Get(c *fiber.Ctx) error {
	paramId := c.Params("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		return err
	}

	db, erro := db.ConnectAzure()
	if erro != nil {
		return c.JSON(erro)
	}
	defer db.Close()

	repository := repositories.NewRepository(db)
	result, err := repository.GetUser(id)

	if err != nil {
		// Update error messages later
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(result)
}

func List(c *fiber.Ctx) error {
	db, erro := db.ConnectAzure()
	if erro != nil {
		return c.JSON(erro)
	}
	defer db.Close()

	repository := repositories.NewRepository(db)
	result, err := repository.ListUsers()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())

	}
	return c.JSON(result)
}

func Delete(c *fiber.Ctx) error {
	paramId := c.Params("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		return err
	}

	db, erro := db.ConnectAzure()
	if erro != nil {
		return c.JSON(erro)
	}
	defer db.Close()

	repository := repositories.NewRepository(db)
	result, err := repository.DeleteUser(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(result)
}

func Update(c *fiber.Ctx) error {
	paramId := c.Params("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		return err
	}

	player := &models.Player{}
	player.Id = uint(id)
	return c.JSON(player)
}

func Create(c *fiber.Ctx) error {
	var body map[string]string
	if err := c.BodyParser(&body); err != nil {
		return err
	}
	passwd, _ := bcrypt.GenerateFromPassword([]byte(body["password"]), bcrypt.DefaultCost)

	localData := models.Player{
		Name:     body["name"],
		Email:    body["email"],
		Password: passwd,
		Score:    0,
	}

	db, erro := db.ConnectAzure()
	if erro != nil {
		return c.JSON(erro)
	}
	defer db.Close()

	repositorie := repositories.NewRepository(db)
	result, err := repositorie.CreateUser(localData)
	if err != nil {
		return c.JSON(err)
	}

	return c.JSON(result)

}

func SendAnswer(c *fiber.Ctx) error {
	var body map[string]interface{}
	if err := c.BodyParser(&body); err != nil {
		return err
	}

	var localData models.Answer
	if err := mapstructure.Decode(body, &localData); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return c.Status(http.StatusBadRequest).SendString("Error decoding JSON")
	}

	db, erro := db.ConnectAzure()
	if erro != nil {
		return c.JSON(erro)
	}

	defer db.Close()

	repositorie := repositories.NewCountryRepo(db)
	result, err := repositorie.GetCountry(uint(localData.CountryId))

	if err != nil {
		return err
	}

	distance := haversine.Calculate(result.Location.Latitude, result.Location.Longitude, localData.Latitude, localData.Longitude)
	score, message := utils.CalculateScore(distance)

	resultData := models.Response{
		Id:          0,
		Amount:      0,
		Score:       score,
		Distance:    distance,
		Message:     message,
		CountryName: result.Name,
	}

	return c.JSON(resultData)

}

func GetScore(c *fiber.Ctx) error {
	db, erro := db.ConnectAzure()
	if erro != nil {
		return c.JSON(erro)
	}

	defer db.Close()

	repository := repositories.NewRepository(db)
	result, err := repository.GetScore()

	if err != nil {
		return err
	}

	return c.JSON(result)
}

func GetAzure(c *fiber.Ctx) error {
	db, err := db.ConnectAzure()
	if err != nil {
		return c.JSON(err)
	}

	defer db.Close()

	repository := repositories.NewRepository(db)
	res, err := repository.GetAzureTest()
	if err != nil {
		return c.JSON(err)
	}

	return c.JSON(res)
}
