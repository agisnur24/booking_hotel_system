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

	SQL := "select i.id, i.invoice_number, i.employee_id, i.meeting_room_pricings_id, i.discount_id, i.pic, e.name, m.price, m.price_type, d.rate, d.status from (((invoices i inner join employees e on i.employee_id=e.id) inner join discounts d on i.discount_id=d.id) inner join meeting_room_pricings m  on i.meeting_room_pricings_id=m.id) where i.id = ?"

	rows, err := tx.QueryContext(ctx, SQL, invoiceId)
	helper.PanicIfError(err)
	defer rows.Close()

	invoice := domain.Invoice{}
	if rows.Next() {
		err := rows.Scan(&invoice.Id, &invoice.Invoice_Number, &invoice.Employee_Id, &invoice.Meeting_Room_Pricings_Id, &invoice.Discount_Id, &invoice.PIC, &invoice.Employee_Name, &invoice.Price, &invoice.Price_Type, &invoice.Discount_Rate, &invoice.Discount_Status)
		helper.PanicIfError(err)

		return invoice, nil
	} else {

		return invoice, errors.New("invoice number is not found")
	}
}

func (repository InvoiceRepoitoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Invoice {

	SQL := "select i.id, i.invoice_number, i.employee_id, i.meeting_room_pricings_id, i.discount_id, i.pic, e.name, m.price, m.price_type, d.rate, d.status from (((invoices i inner join employees e on i.employee_id=e.id) inner join discounts d on i.discount_id=d.id) inner join meeting_room_pricings m  on i.meeting_room_pricings_id=m.id)"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var invoices []domain.Invoice
	for rows.Next() {
		invoice := domain.Invoice{}
		err := rows.Scan(&invoice.Id, &invoice.Invoice_Number, &invoice.Employee_Id, &invoice.Meeting_Room_Pricings_Id, &invoice.Discount_Id, &invoice.PIC, &invoice.Employee_Name, &invoice.Price, &invoice.Price_Type, &invoice.Discount_Rate, &invoice.Discount_Status)
		helper.PanicIfError(err)
		invoices = append(invoices, invoice)
	}

	return invoices

}
