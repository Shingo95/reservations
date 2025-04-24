package repository

import "github.com/elorenzotti/bookings/internal/models"

type DatabaseRepo interface {
	InsertReservation(res models.Reservation) (int, error)
	InsertRestrictions(r *models.RoomRestriction) error
}
