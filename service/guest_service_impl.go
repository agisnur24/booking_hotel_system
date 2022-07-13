package service

import (
	"context"
	"database/sql"
	"github.com/agisnur24/booking_hotel_system.git/exception"

	"github.com/agisnur24/booking_hotel_system.git/helper"
	"github.com/agisnur24/booking_hotel_system.git/model/domain"
	"github.com/agisnur24/booking_hotel_system.git/model/web"
	"github.com/agisnur24/booking_hotel_system.git/repository"
	"github.com/go-playground/validator/v10"
)

type InvoiceServiceImpl struct {
	InvoiceRepository repository.InvoiceRepository
	DB                *sql.DB
	validate          *validator.Validate
}

func NewInvoiceService(invoiceRepository repository.InvoiceRepository, DB *sql.DB, validate *validator.Validate) InvoiceService {
	return &InvoiceServiceImpl{
		InvoiceRepository: invoiceRepository,
		DB:                DB,
		validate:          validate,
	}
}

func (service *InvoiceServiceImpl) Create(ctx context.Context, request web.InvoiceCreateRequest) web.InvoiceResponse {
	err := service.validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	invoice := domain.Invoice{
		Invoice_Number:           request.Invoice_Number,
		Employee_Id:              request.Employee_Id,
		Meeting_Room_Pricings_Id: request.Meeting_Room_Pricings_Id,
		Discount_Id:              request.Discount_Id,
		PIC:                      request.PIC,
	}

	invoice = service.InvoiceRepository.Create(ctx, tx, invoice)
	return helper.ToInvoiceResponse(invoice)
}

func (service *InvoiceServiceImpl) Update(ctx context.Context, request web.InvoiceUpdateRequest) web.InvoiceResponse {
	err := service.validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	invoice, err := service.InvoiceRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	invoice.Invoice_Number = request.Invoice_Number
	invoice.Employee_Id = request.Employee_Id
	invoice.Meeting_Room_Pricings_Id = request.Meeting_Room_Pricings_Id
	invoice.Discount_Id = request.Discount_Id
	invoice.PIC = request.PIC

	invoice = service.InvoiceRepository.Update(ctx, tx, invoice)

	return helper.ToInvoiceResponse(invoice)
}

func (service *InvoiceServiceImpl) Delete(ctx context.Context, invoiceId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	invoice, err := service.InvoiceRepository.FindById(ctx, tx, invoiceId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.InvoiceRepository.Delete(ctx, tx, invoice)
}

func (service *InvoiceServiceImpl) FindById(ctx context.Context, invoiceId int) web.InvoiceResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	invoice, err := service.InvoiceRepository.FindById(ctx, tx, invoiceId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToInvoiceResponse(invoice)
}

func (service *InvoiceServiceImpl) FindAll(ctx context.Context) []web.InvoiceResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	invoices := service.InvoiceRepository.FindAll(ctx, tx)

	return helper.ToInvoiceResponses(invoices)
}
