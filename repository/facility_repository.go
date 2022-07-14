package repository

import (
	"context"
	"database/sql"
	"github.com/agisnur24/booking_hotel_system.git/model/domain"
)

type FacilityRepository interface {
	Create(ctx context.Context, tx *sql.Tx, facility domain.Facility) domain.Facility
	Update(ctx context.Context, tx *sql.Tx, facility domain.Facility) domain.Facility
	Delete(ctx context.Context, tx *sql.Tx, facility domain.Facility)
	FindById(ctx context.Context, tx *sql.Tx, facilityId int) (domain.Facility, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Facility
}
