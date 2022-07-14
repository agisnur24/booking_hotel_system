package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/agisnur24/booking_hotel_system.git/helper"
	"github.com/agisnur24/booking_hotel_system.git/model/domain"
)

type InvoiceRepoitoryImpl struct {
}

func NewInvoiceRepository() InvoiceRepository {
	return &InvoiceRepoitoryImpl{}
}

func (repository InvoiceRepoitoryImpl) Create(ctx context.Context, tx *sql.Tx, invoice domain.Invoice) domain.Invoice {

	SQL := "insert into invoices(invoice_number, employee_id, meeting_room_pricings_id, discount_id, pic) value(?,?,?,?,?)"
	result, err := tx.ExecContext(ctx, SQL, invoice.Invoice_Number, invoice.Employee_Id, invoice.Meeting_Room_Pricings_Id, invoice.Discount_Id, invoice.PIC)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	invoice.Id = int(id)

	return invoice
}

func (repository InvoiceRepoitoryImpl) Update(ctx context.Context, tx *sql.Tx, invoice domain.Invoice) domain.Invoice {

	SQL := "update invoices set invoice_number = ?, employee_id = ?, meeting_room_pricings_id=?, discount_id=?, pic=? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, invoice.Invoice_Number, invoice.Employee_Id, invoice.Meeting_Room_Pricings_Id, invoice.Discount_Id, invoice.PIC)
	helper.PanicIfError(err)

	return invoice
}

func (repository InvoiceRepoitoryImpl) Delete(ctx context.Context, tx *sql.Tx, invoice domain.Invoice) {

	SQL := "delete from invoices where id = ?"
	_, err := tx.ExecContext(ctx, SQL, invoice.Id)
	helper.PanicIfError(err)

}

func (repository InvoiceRepoitoryImpl) FindById(ctx context.Context, tx *sql.Tx, invoiceId int) (domain.Invoice, error) {

	SQL := "select id, name, address, phone_number,email from invoices where id =?"
	rows, err := tx.QueryContext(ctx, SQL, invoiceId)
	helper.PanicIfError(err)
	defer rows.Close()

	invoice := domain.Invoice{}
	if rows.Next() {
		err := rows.Scan(&invoice.Id, &invoice.Name, &invoice.Address, &invoice.Phone_Number, &invoice.Email)
		helper.PanicIfError(err)

		return invoice, nil
	} else {

		return invoice, errors.New("invoice number is not found")
	}
}

func (repository InvoiceRepoitoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Invoice {

	SQL := "select id,name,address,phone_number,email from invoices"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var invoices []domain.Invoice
	for rows.Next() {
		invoice := domain.Invoice{}
		err := rows.Scan(&invoice.Id, &invoice.Name, &invoice.Address, &invoice.Phone_Number, &invoice.Email)
		helper.PanicIfError(err)
		invoices = append(invoices, invoice)
	}

	return invoices

}

/*type GuestRepoitoryImpl struct {
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

}*/
