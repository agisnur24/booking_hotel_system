package repository

import (
	"context"
	"database/sql"
	"github.com/agisnur24/booking_hotel_system.git/model/domain"
)

type GuestRepository interface {
	Create(ctx context.Context, tx *sql.Tx, guest domain.Guest) domain.Guest
	Update(ctx context.Context, tx *sql.Tx, guest domain.Guest) domain.Guest
	Delete(ctx context.Context, tx *sql.Tx, guest domain.Guest)
	FindById(ctx context.Context, tx *sql.Tx, guestId int) (domain.Guest, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Guest
}
