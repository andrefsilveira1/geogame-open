package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"geolocation/models"
	"strconv"
	"strings"
)

type data struct {
	db *sql.DB
}

func NewCountryRepo(db *sql.DB) *data {
	return &data{db}
}

func (d data) GetCountries() ([]models.Country, error) {
	rows, err := d.db.Query("SELECT country.id, country.name, country.level, country.score, " +
		"country.latitude AS location_latitude, country.longitude AS location_longitude, " +
		"tips.id AS tip_id, tips.country_id AS tip_country_id, tips.tip_number, tips.text AS tip_text " +
		"FROM country " +
		"LEFT JOIN tips ON country.id = tips.country_id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var datas []models.Country
	currentCountryID := uint(0)
	var currentCountry *models.Country

	for rows.Next() {
		var countryID uint
		var tipID sql.NullInt64
		var data models.Country
		var tip models.Tips

		if err = rows.Scan(
			&countryID,
			&data.Name,
			&data.Level,
			&data.Score,
			&data.Location.Latitude,
			&data.Location.Longitude,
			&tipID,
			&tip.CountryId,
			&tip.TipNumber,
			&tip.Text,
		); err != nil {
			return nil, err
		}

		if countryID != currentCountryID {
			if currentCountry != nil {
				datas = append(datas, *currentCountry)
			}

			currentCountryID = countryID
			currentCountry = &models.Country{
				Id:       countryID,
				Name:     data.Name,
				Level:    data.Level,
				Score:    data.Score,
				Location: data.Location,
			}
		}

		if tipID.Valid {
			currentCountry.Tips = append(currentCountry.Tips, models.Tips{
				Id:        uint(tipID.Int64),
				CountryId: tip.CountryId,
				TipNumber: tip.TipNumber,
				Text:      tip.Text,
			})
		}
	}

	// Append the last country to the datas array
	if currentCountry != nil {
		datas = append(datas, *currentCountry)
	}

	return datas, nil

}

func (d data) PostCountry(country models.Country) (string, error) {
	statement, err := d.db.Prepare("INSERT INTO country (id, name, level, score, latitude, longitude) VALUES (NULL, ?, ?, ?, ?, ?)")

	if err != nil {
		return "FAIL", err
	}

	result, err := statement.Exec(country.Name, country.Level, country.Score, country.Location.Latitude, country.Location.Longitude)

	if err != nil {
		statement.Close()
		return "FAIL", err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return "FAIL ON ID", err
	}

	tipStatement, err := d.db.Prepare("INSERT INTO tips (country_id, tip_number, text) VALUES (?, ?, ?)")
	if err != nil {
		return "FAIL on TIPS", err
	}

	for _, tip := range country.Tips {
		_, err := tipStatement.Exec(id, tip.TipNumber, tip.Text)
		if err != nil {
			tipStatement.Close()
			return "FAIL on EXEC", err
		}
	}

	tipStatement.Close()

	statement.Close()

	tipStatement.Close()

	return "Country created", nil
}

func (d data) GetCountry(id uint) (models.Country, error) {
	row := d.db.QueryRow("SELECT id, name, level, score, latitude, longitude FROM country WHERE id = @p1", id)

	tipRows, err := d.db.Query("SELECT id, tip_number, text FROM tips WHERE country_id = @p1", id)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return models.Country{}, err
	}
	defer tipRows.Close()

	country := models.Country{}

	if err := row.Scan(&country.Id, &country.Name, &country.Level, &country.Score, &country.Location.Latitude, &country.Location.Longitude); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Country{}, fmt.Errorf("Country with ID %d not found", id)
		}
		return models.Country{}, err
	}

	tips := []models.Tips{}

	for tipRows.Next() {
		tip := models.Tips{}
		var tipText sql.NullString

		if err := tipRows.Scan(&tip.Id, &tip.TipNumber, &tipText); err != nil {
			return models.Country{}, err
		}

		if tipText.Valid {
			tip.Text = tipText.String
		}

		tips = append(tips, tip)
	}

	country.Tips = tips

	return country, nil
}

func (d data) DeleteCountry(id uint) (string, error) {
	tx, err := d.db.Begin()
	if err != nil {
		return "Error: begin DB", err
	}

	_, err = tx.Exec("DELETE FROM tips WHERE country_id = ?", id)
	if err != nil {
		tx.Rollback()
		return "Error: Exec Tips", err
	}

	_, err = tx.Exec("DELETE FROM country WHERE id = ?", id)
	if err != nil {
		tx.Rollback()
		return "Error: Exec Country", err
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return "Error: Commit", err
	}

	return "Delete Successful", nil
}

func (d data) UpdateCountry(id uint, country models.Country) (string, error) {
	statement, err := d.db.Prepare("UPDATE country SET name=?, level=?, score=?, latitude=?, longitude=? WHERE id=?")
	if err != nil {
		return "FAIL: Error at preparing stage", err
	}

	_, err = statement.Exec(country.Name, country.Level, country.Score, country.Location.Latitude, country.Location.Longitude, id)
	if err != nil {
		statement.Close()
		return "FAIL: Error at Exec stage", err
	}

	tipStatement, err := d.db.Prepare("UPDATE tips SET country_id=?, tip_number=?, text=? WHERE id=? AND country_id=?")
	if err != nil {
		return "FAIL on TIPS prepare stage", err
	}

	for _, tip := range country.Tips {
		_, err := tipStatement.Exec(id, tip.TipNumber, tip.Text, tip.Id, id)
		if err != nil {
			tipStatement.Close()
			return "FAIL on EXEC", err
		}
	}

	statement.Close()

	tipStatement.Close()

	return "UPDATE SUCCESSFUL", nil
}

func (d data) RandomCountry(ids []int) ([]models.Country, error) {
	placeholders := make([]string, len(ids))
	params := make([]interface{}, len(ids))

	for i, id := range ids {
		placeholders[i] = "@p" + strconv.Itoa(i+1)
		params[i] = id
	}

	query := "SELECT country.id, country.name, country.level, country.score, " +
		"country.latitude AS location_latitude, country.longitude AS location_longitude, " +
		"tips.id AS tip_id, tips.country_id AS tip_country_id, tips.tip_number, tips.text AS tip_text " +
		"FROM country " +
		"LEFT JOIN tips ON country.id = tips.country_id " +
		"WHERE country.id NOT IN (" + strings.Join(placeholders, ",") + ")"

	rows, err := d.db.Query(query, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var datas []models.Country
	currentCountryID := uint(0)
	var currentCountry *models.Country

	for rows.Next() {
		var countryID uint
		var tipID sql.NullInt64
		var data models.Country
		var tip models.Tips

		if err = rows.Scan(
			&countryID,
			&data.Name,
			&data.Level,
			&data.Score,
			&data.Location.Latitude,
			&data.Location.Longitude,
			&tipID,
			&tip.CountryId,
			&tip.TipNumber,
			&tip.Text,
		); err != nil {
			return nil, err
		}

		if countryID != currentCountryID {
			if currentCountry != nil {
				datas = append(datas, *currentCountry)
			}

			currentCountryID = countryID
			currentCountry = &models.Country{
				Id:       countryID,
				Name:     data.Name,
				Level:    data.Level,
				Score:    data.Score,
				Location: data.Location,
			}
		}

		if tipID.Valid {
			currentCountry.Tips = append(currentCountry.Tips, models.Tips{
				Id:        uint(tipID.Int64),
				CountryId: tip.CountryId,
				TipNumber: tip.TipNumber,
				Text:      tip.Text,
			})
		}
	}

	if currentCountry != nil {
		datas = append(datas, *currentCountry)
	}

	return datas, nil
}
