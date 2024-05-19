package handlers

import (
	"encoding/json"
	"net/http"

	"backend/models"
	"database/sql"
)

func MakeReservation(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var reservation models.Reservation
		if err := json.NewDecoder(r.Body).Decode(&reservation); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Insert reservation into the database
		if err := reservation.CreateReservation(db); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}
