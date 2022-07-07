package repository

import (
	"context"
	"database/sql"
	"github.com/agisnur24/booking_hotel_system.git/model/domain"
)

type FloorRepository interface {
	Create(ctx context.Context, tx *sql.Tx, floor domain.Floor) domain.Floor
	Update(ctx context.Context, tx *sql.Tx, floor domain.Floor) domain.Floor
	Delete(ctx context.Context, tx *sql.Tx, floor domain.Floor)
	FindById(ctx context.Context, tx *sql.Tx, floorId int) (domain.Floor, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Floor
}
