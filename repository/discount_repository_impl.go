package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/agisnur24/booking_hotel_system.git/helper"
	"github.com/agisnur24/booking_hotel_system.git/model/domain"
)

type DiscountRepositoryImpl struct {
}

func NewDiscountRepository() DiscountRepository {
	return &DiscountRepositoryImpl{}
}

// id
//employee_id
//rate
//status
//request_date

func (repository DiscountRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, discount domain.Discount) domain.Discount {
	SQL := "insert into discounts(employee_id, rate, status, request_date,hotel_id,room_id) values (?,?, ?, ?,?,?)"
	result, err := tx.ExecContext(ctx, SQL, discount.EmployeeId, discount.Rate, discount.Status, discount.RequestDate, discount.HotelId, discount.RoomId)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	discount.Id = int(id)
	return discount
}

func (repository DiscountRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, discount domain.Discount) domain.Discount {
	SQL := "update discounts set employee_id =?, rate = ?, status = ?, request_date = ?, hotel_id=?, room_id=? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, discount.EmployeeId, discount.Rate, discount.Status, discount.RequestDate, discount.HotelId, discount.RoomId, discount.Id)
	helper.PanicIfError(err)

	return discount
}

func (repository DiscountRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, discount domain.Discount) {
	SQL := "delete from discounts where id = ?"
	_, err := tx.ExecContext(ctx, SQL, discount.Id)
	helper.PanicIfError(err)
}

func (repository DiscountRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, discountId int) (domain.Discount, error) {
	SQL := "select d.id, d.employee_id, d.hotel_id, d.room_id, d.rate, d.status, d.request_date, e.name as employee_name, h.name as hotel_name, r.name as room_name from discounts d inner join employees e on d.employee_id=e.id inner join hotels h on d.hotel_id=h.id inner join meeting_rooms r on d.room_id=r.id where d.id = ?"
	rows, err := tx.QueryContext(ctx, SQL, discountId)
	helper.PanicIfError(err)
	defer rows.Close()

	discount := domain.Discount{}
	if rows.Next() {
		err := rows.Scan(&discount.Id, &discount.EmployeeId, &discount.HotelId, &discount.RoomId, &discount.Rate, &discount.Status, &discount.RequestDate, &discount.EmployeeName, &discount.HotelName, &discount.RoomName)
		helper.PanicIfError(err)
		return discount, nil
	} else {
		return discount, errors.New("id not found")
	}
}

func (repository DiscountRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Discount {
	//SQL := "select u.id, u.name, u.email, u.role_id, r.role_name from users u inner join roles r on u.role_id=r.id"
	SQL := "select d.id, d.employee_id, d.hotel_id, d.room_id, d.rate, d.status, d.request_date, e.name as employee_name, h.name as hotel_name, r.name as room_name from discounts d inner join employees e on d.employee_id=e.id inner join hotels h on d.hotel_id=h.id inner join meeting_rooms r on d.room_id=r.id"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var discounts []domain.Discount
	for rows.Next() {
		discount := domain.Discount{}
		err := rows.Scan(&discount.Id, &discount.EmployeeId, &discount.HotelId, &discount.RoomId, &discount.Rate, &discount.Status, &discount.RequestDate, &discount.EmployeeName, &discount.HotelName, &discount.RoomName)
		helper.PanicIfError(err)
		discounts = append(discounts, discount)
	}
	return discounts
}
