package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/agisnur24/booking_hotel_system.git/helper"
	"github.com/agisnur24/booking_hotel_system.git/model/domain"
)

type GuestRepoitoryImpl struct {
}

func NewGuestRepository() GuestRepository {
	return &GuestRepoitoryImpl{}
}

func (repository GuestRepoitoryImpl) Create(ctx context.Context, tx *sql.Tx, guest domain.Guest) domain.Guest {

	SQL := "insert into guests(name, address, phone_number, email) value(?,?,?,?)"
	result, err := tx.ExecContext(ctx, SQL, guest.Name, guest.Address, guest.Phone_Number, guest.Email)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	guest.Id = int(id)

	return guest
}

func (repository GuestRepoitoryImpl) Update(ctx context.Context, tx *sql.Tx, guest domain.Guest) domain.Guest {

	SQL := "update guests set name = ?, address = ?, phone_number=?, email =?  where id = ?"
	_, err := tx.ExecContext(ctx, SQL, guest.Name, guest.Address, guest.Phone_Number, guest.Email, guest.Id)
	helper.PanicIfError(err)

	return guest
}

func (repository GuestRepoitoryImpl) Delete(ctx context.Context, tx *sql.Tx, guest domain.Guest) {

	SQL := "delete from guests where id = ?"
	_, err := tx.ExecContext(ctx, SQL, guest.Id)
	helper.PanicIfError(err)

}

func (repository GuestRepoitoryImpl) FindById(ctx context.Context, tx *sql.Tx, guestId int) (domain.Guest, error) {

	SQL := "select id, name, address, phone_number,email from guests where id =?"
	rows, err := tx.QueryContext(ctx, SQL, guestId)
	helper.PanicIfError(err)
	defer rows.Close()

	guest := domain.Guest{}
	if rows.Next() {
		err := rows.Scan(&guest.Id, &guest.Name, &guest.Address, &guest.Phone_Number, &guest.Email)
		helper.PanicIfError(err)

		return guest, nil
	} else {

		return guest, errors.New("guest is not found")
	}
}

func (repository GuestRepoitoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Guest {

	SQL := "select id,name,address,phone_number,email from guests"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var guests []domain.Guest
	for rows.Next() {
		guest := domain.Guest{}
		err := rows.Scan(&guest.Id, &guest.Name, &guest.Address, &guest.Phone_Number, &guest.Email)
		helper.PanicIfError(err)
		guests = append(guests, guest)
	}

	return guests

}
