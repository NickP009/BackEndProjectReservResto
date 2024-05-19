package models

import (
	"database/sql"
)

type Reservation struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Gender          string `json:"gender"`
	Email           string `json:"email"`
	PhoneNumber     string `json:"phoneNumber"`
	ReservationDay  string `json:"reservationDay"`
	ReservationTime string `json:"reservationTime"`
	NumberOfPeople  int    `json:"numberOfPeople"`
}

func (r *Reservation) CreateReservation(db *sql.DB) error {
	query := `INSERT INTO reservations (name, gender, email, phone_number, reservation_day, reservation_time, number_of_people) VALUES (?, ?, ?, ?, ?, ?, ?)`
	_, err := db.Exec(query, r.Name, r.Gender, r.Email, r.PhoneNumber, r.ReservationDay, r.ReservationTime, r.NumberOfPeople)
	return err
}
