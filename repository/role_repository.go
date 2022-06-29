package repository

import (
	"context"
	"database/sql"

	"github.com/agisnur24/booking_hotel_system.git/model/domain"
)

type RoleRepository interface {
	Create(ctx context.Context, tx *sql.Tx, user domain.Role) domain.Role
	Update(ctx context.Context, tx *sql.Tx, user domain.Role) domain.Role
	Delete(ctx context.Context, tx *sql.Tx, user domain.Role)
	FindByRoleName(ctx context.Context, tx *sql.Tx, roleName string) (domain.Role, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Role
}
