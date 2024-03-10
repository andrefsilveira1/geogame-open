package models

import (
	"math/rand"
	"time"
)

type Tips struct {
	Id        uint   `json:"id,omitempty"`
	TipNumber uint   `json:"number,omitempty"`
	CountryId uint   `json:"country_id,omitempty"`
	Text      string `json:"text,omitempty"`
}

type Coordinates struct {
	Latitude  float64
	Longitude float64
}

type Country struct {
	Id       uint        `json:"id,omitempty"`
	Name     string      `json:"name,omitempty"`
	Tips     []Tips      `json:"tips,omitempty"`
	Level    uint        `json:"level,omitempty"`
	Score    uint        `json:"score,omitempty"`
	Location Coordinates `json:"location,omitempty"`
}

var countryFacts = map[string][]string{
	"USA": {
		"The United States is the third-largest country in the world by land area.",
		"The U.S. Declaration of Independence was adopted on July 4, 1776.",
		"The Statue of Liberty was a gift from France to the United States.",
		"The Grand Canyon is one of the most famous natural landmarks in the U.S.",
		"The United States has 50 states.",
	},

	"Canada": {
		"Canada is the second-largest country in the world by land area.",
		"Canada is known for its maple syrup production.",
		"The CN Tower in Toronto is one of the tallest free-standing structures in the world.",
		"Canada has two official languages: English and French.",
		"Banff National Park in Canada is famous for its stunning mountain scenery.",
	},

	"Brazil": {
		"Brazil is the largest country in South America.",
		"The Amazon Rainforest, located in Brazil, is the largest rainforest in the world.",
		"Carnival in Rio de Janeiro is one of the most famous festivals in Brazil.",
		"The Christ the Redeemer statue in Rio de Janeiro is an iconic symbol of Brazil.",
		"Samba music and dance are popular cultural expressions in Brazil.",
	},
}

func generateRandomCountry() Country {
	rand.Seed(time.Now().UnixNano()) // Seed for randomness

	// Example country names
	countryNames := []string{"USA", "Canada", "Brazil", "Germany", "Japan", "Australia"}

	// Generate a random country
	country := Country{
		Id:    uint(rand.Intn(100)), // Example: random ID between 0 and 99
		Name:  countryNames[rand.Intn(len(countryNames))],
		Tips:  generateRandomTips(countryNames),
		Level: uint(rand.Intn(10)),  // Example: random level between 0 and 9
		Score: uint(rand.Intn(100)), // Example: random score between 0 and 99
	}

	return country
}

func generateRandomTips(countryNames []string) []Tips {
	var tips []Tips
	for _, countryName := range countryNames {
		facts, exists := countryFacts[countryName]
		if !exists {
			continue
		}

		// Shuffle the facts randomly
		rand.Shuffle(len(facts), func(i, j int) {
			facts[i], facts[j] = facts[j], facts[i]
		})

		// Take the first 5 facts
		for i := 0; i < 5 && i < len(facts); i++ {
			tips = append(tips, Tips{
				Id:        uint(i),
				TipNumber: uint(i + 1),
				Text:      facts[i],
			})
		}
	}
	return tips
}

func GenerateRandomCountries() []Country {
	var countries []Country
	for i := 0; i < 3; i++ {
		countries = append(countries, generateRandomCountry())
	}
	return countries
}
