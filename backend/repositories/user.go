package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"geolocation/models"
)

type Data struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Data {
	return &Data{db}
}

func (d Data) CreateUser(user models.Player) (uint64, error) {
	statement, err := d.db.Prepare("INSERT INTO Users (name, email, password, score) VALUES (?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}

	result, err := statement.Exec(user.Name, user.Email, user.Password, user.Score)
	if err != nil {
		statement.Close()
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	statement.Close()

	return uint64(id), nil
}

func (d Data) GetUser(id int) (models.Player, error) {
	row := d.db.QueryRow("SELECT id, name, email, score FROM Users WHERE id = ?", id)
	player := models.Player{}

	if err := row.Scan(&player.Id, &player.Name, &player.Email, &player.Score); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Player{}, fmt.Errorf("user with ID %d not found", id)
		}
		return models.Player{}, err
	}

	return player, nil
}
func (d Data) ListUsers() ([]models.Player, error) {
	rows, err := d.db.Query("SELECT id, name, email, score FROM Users")
	if err != nil {
		return nil, err
	}

	var datas []models.Player
	for rows.Next() {
		var data models.Player
		if err = rows.Scan(
			&data.Id,
			&data.Name,
			&data.Email,
			&data.Score,
		); err != nil {
			return nil, err
		}
		datas = append(datas, data)
	}
	return datas, nil
}

func (d Data) DeleteUser(id int) (string, error) {
	result, err := d.db.Exec("DELETE FROM Users WHERE id = ?", id)
	if err != nil {
		return "ERROR", err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return "ERROR", err
	}

	if rowsAffected == 0 {
		return "ERROR: ID NOT FOUND", nil
	}

	return "SUCCESS", nil
}

func (d Data) SendScore(sent models.Score) (string, error) {
	statement, err := d.db.Prepare("INSERT INTO players (name, score) VALUES (@p1, @p2)")
	if err != nil {
		return "ERROR", err
	}

	result, err := statement.Exec(sent.Name, sent.Score)
	if err != nil {
		statement.Close()
		return "ERROR", err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return "ERROR", err
	}

	fmt.Println("ID:", id)
	statement.Close()
	return "SUCCESS", nil
}

func (d Data) GetScore() ([]models.Score, error) {
	rows, err := d.db.Query("SELECT name, score FROM players")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var players []models.Score
	for rows.Next() {
		var data models.Score
		if err = rows.Scan(
			&data.Name,
			&data.Score,
		); err != nil {
			return nil, err
		}
		players = append(players, data)
	}
	return players, nil
}

func (d Data) GetAzureTest() (string, error) {
	row, err := d.db.Query("SELECT 1")
	if err != nil {
		return "ERROR", err
	}

	fmt.Println("ROW:", row)
	return "SUCCESS", nil
}
