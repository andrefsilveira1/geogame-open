package controllers

import (
	"fmt"
	db "geolocation/database"
	"geolocation/models"
	"geolocation/repositories"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/mitchellh/mapstructure"
)

func GetCountry(c *fiber.Ctx) error {
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

	repository := repositories.NewCountryRepo(db)
	result, err := repository.GetCountry(uint(id))

	if err != nil {
		return err
	}

	fmt.Println("LOGGING: Geting Country")

	return c.JSON(result)
}

func UpdateCountry(c *fiber.Ctx) error {
	paramId := c.Params("id")
	id, err := strconv.Atoi(paramId)

	if err != nil {
		return err
	}

	var body map[string]interface{}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(http.StatusBadRequest).SendString("Error decoding JSON")
	}

	var localData models.Country
	if err := mapstructure.Decode(body, &localData); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return c.Status(http.StatusBadRequest).SendString("Error decoding JSON")
	}

	db, erro := db.ConnectAzure()
	if erro != nil {
		return c.JSON(erro)
	}

	defer db.Close()

	repository := repositories.NewCountryRepo(db)
	result, err := repository.UpdateCountry(uint(id), localData)

	if err != nil {
		return err
	}

	fmt.Println("LOGGING: Updating Country")

	return c.JSON(result)
}

func DeleteCountry(c *fiber.Ctx) error {
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

	repository := repositories.NewCountryRepo(db)
	result, err := repository.DeleteCountry(uint(id))

	if err != nil {
		return err
	}

	fmt.Println("LOGGING: Deleting Country")

	return c.JSON(result)
}

func ListCountries(c *fiber.Ctx) error {
	db, erro := db.ConnectAzure()
	if erro != nil {
		return c.JSON(erro)
	}
	defer db.Close()

	repository := repositories.NewCountryRepo(db)
	result, err := repository.GetCountries()
	if err != nil {
		return err
	}

	fmt.Println("LOGGING: Listing countries")

	return c.JSON(result)
}

func PostCountry(c *fiber.Ctx) error {
	var body map[string]interface{}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(http.StatusBadRequest).SendString("Error decoding JSON")
	}

	var localData models.Country
	if err := mapstructure.Decode(body, &localData); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return c.Status(http.StatusBadRequest).SendString("Error decoding JSON")
	}

	db, erro := db.ConnectAzure()

	if erro != nil {
		return c.JSON(erro)
	}
	defer db.Close()

	repository := repositories.NewCountryRepo(db)
	result, err := repository.PostCountry(localData)
	if err != nil {
		return err
	}
	fmt.Println("LOGGING: Posting Country")

	return c.JSON(result)
}

func ResponseRandomCountry(c *fiber.Ctx) error {
	var body map[string]interface{}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(http.StatusBadRequest).SendString("Error decoding JSON")
	}

	var localData models.Excluded
	if err := mapstructure.Decode(body, &localData); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return c.Status(http.StatusBadRequest).SendString("Error decoding JSON")
	}

	db, erro := db.ConnectAzure()
	if erro != nil {
		return c.JSON(erro)
	}
	defer db.Close()

	var result []models.Country
	var err error
	repository := repositories.NewCountryRepo(db)

	if len(localData.Played.Ids) > 0 {
		result, err = repository.RandomCountry(localData.Played.Ids)
		if err != nil {
			return err
		}

	} else {
		result, err = repository.GetCountries()
		if err != nil {
			return err
		}
	}

	if len(result) > 0 {
		rand.Seed(time.Now().UnixNano())
		randomIndex := rand.Intn(len(result))
		pick := result[randomIndex]
		return c.JSON(pick)
	}

	fmt.Println("LOGGING: Sending random country")

	return c.JSON(nil)

}

func SendScore(c *fiber.Ctx) error {
	var body map[string]interface{}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(http.StatusBadRequest).SendString("Error decoding JSON")
	}

	var localData models.Score
	if err := mapstructure.Decode(body, &localData); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return c.Status(http.StatusBadRequest).SendString("Error decoding JSON")
	}

	db, erro := db.ConnectAzure()
	if erro != nil {
		return c.JSON(erro)
	}
	defer db.Close()

	repository := repositories.NewRepository(db)
	result, err := repository.SendScore(localData)

	if err != nil {
		return c.JSON(err)
	}

	fmt.Println("LOGGING: Sending score")

	return c.JSON(result)
}
