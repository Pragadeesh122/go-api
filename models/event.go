package models

import (
	"go_api/db"
	"time"
)

type Event struct {
	ID          int64
	Name        string `binding:"required"`
	Description string `binding:"required"`
	Location    string `binding:"required"`
	Date        time.Time
	UserID      int64
}

func (event *Event) Save() error {

	query := `
	INSERT INTO events(name,description,location,date,user_id)
	VALUES (?,?,?,?,?)
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()
	result, err := stmt.Exec(event.Name, event.Description, event.Location, event.Date, event.UserID)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	event.ID = id
	return err
}

func GetAllEvents() ([]Event, error) {
	query := `
	SELECT * FROM EVENTS
	`
	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.Date, &event.UserID)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}
	return events, nil
}

func GetEvent(id int64) (*Event, error) {

	query := `
	SELECT * FROM events WHERE id = ?
	`
	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.Date, &event.UserID)

	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (event *Event) UpdateEvent() error {

	query := `
		UPDATE events
		SET name = ?,
			description = ?,
			location = ?,
			date = CURRENT_TIMESTAMP,
			user_id = ?
		WHERE id = ?
	`
	_, err := db.DB.Exec(query, event.Name, event.Description, event.Location, event.UserID, event.ID)

	if err != nil {
		return err
	}

	return nil

}

func DeleteEvent(id int64) error {

	query := `
	DELETE FROM events WHERE id = ?
	`
	_, err := db.DB.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}
