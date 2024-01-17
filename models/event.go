package models

import (
	"Rest_Api_Go/db"
	"errors"
	"time"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int64
}

//var events := []Event{} //empty array of type Event

func (e *Event) Save() error {
	query := `
INSERT INTO events
    (name,description,location,dateTime,user_id)
Values (?,?,?,?,?)` //????? prevents sql injection attacks
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	e.ID = id
	return err
	//events = append(events, e)
}

func GetALlEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close() //free up memory by closing fetched rows
	var events []Event
	//read all rows
	for rows.Next() {
		var event Event //reference to data at current row returned
		err = rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func GetEventByID(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"
	row := db.DB.QueryRow(query, id)
	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

	if err != nil {
		return nil, err
	}
	return &event, nil
}

func (event Event) Update() error {
	query := `
UPDATE events 
SET name = ?, description = ?, location = ?, dateTime = ?
WHERE id = ?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID)
	return err
}

func DeleteEvent(eventID int64) error {
	query := "DELETE FROM events WHERE id = ?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(eventID)
	return err
}

func (e *Event) Register(userID int64) error {
	query := "INSERT INTO registrations(eventID, userID) VALUES(?,?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.ID, userID)
	//if err != nil {
	return err
	//}
	//id, err := res.LastInsertId()
	//if err != nil {
	//	return err
	//}
}

func (e *Event) Cancel(userID int64) error {
	query := "DELETE FROM registrations WHERE userID = ? AND eventID = ?"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(userID, e.ID)
	numRows, err := result.RowsAffected()
	if numRows == 0 || err != nil {
		return errors.New("Unable to delete registration")
	}
	return nil
}
