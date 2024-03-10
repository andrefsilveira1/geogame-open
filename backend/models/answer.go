package models

type Answer struct {
	Latitude  float64 `json:"latitude,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
	CountryId uint    `json:"countryId,omitempty"`
	Time      float64 `json:"time,omitempty"`
}

type Response struct {
	Id          uint
	Score       int
	Amount      int
	Distance    float64
	Message     string
	CountryName string
}

type Excluded struct {
	Played struct {
		Ids []int `json:"Ids"`
	} `json:"played"`
}

type Score struct {
	Name  string  `json:"name"`
	Score float64 `json:"score"`
}
