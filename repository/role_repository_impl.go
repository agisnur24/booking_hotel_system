package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/agisnur24/booking_hotel_system.git/helper"
	"github.com/agisnur24/booking_hotel_system.git/model/domain"
)

type RoleRepositoryImpl struct {
}

func NewRoleRepository() RoleRepository {
	return &RoleRepositoryImpl{}
}

func (repository RoleRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, role domain.Role) domain.Role {
	SQL := "insert into role(rolename, description) values (?, ?)"
	result, err := tx.ExecContext(ctx, SQL, role.RoleName, role.RoleDescription)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	role.Id = int(id)
	return role
}

func (repository RoleRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, role domain.Role) domain.Role {
	SQL := "update role set rolename = ?, roledescription = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, role.RoleName, role.RoleDescription)
	helper.PanicIfError(err)

	return role
}

func (repository RoleRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, user domain.Role) {

}

func (repository RoleRepositoryImpl) FindByRoleName(ctx context.Context, tx *sql.Tx, roleName string) (domain.Role, error) {
	SQL := "select id, RoleName, RoleDescription from role where RoleName = ?"
	rows, err := tx.QueryContext(ctx, SQL, roleName)
	helper.PanicIfError(err)
	defer rows.Close()

	role := domain.Role{}
	if rows.Next() {
		err := rows.Scan(&role.Id, &role.RoleName, &role.RoleDescription)
		helper.PanicIfError(err)
		return role, nil
	} else {
		return role, errors.New("Role Name is not found")
	}
}

func (repository RoleRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Role {
	SQL := "select id, rolename, roledescription from role"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var roles []domain.Role
	for rows.Next() {
		role := domain.Role{}
		err := rows.Scan(&role.Id, &role.RoleName, &role.RoleDescription)
		helper.PanicIfError(err)
		roles = append(roles, role)
	}
	return roles
}
