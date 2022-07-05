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

type GuestServiceImpl struct {
	GuestRepository repository.GuestRepository
	DB              *sql.DB
	validate        *validator.Validate
}

func NewGuestService(guestRepository repository.GuestRepository, DB *sql.DB, validate *validator.Validate) GuestService {
	return &GuestServiceImpl{
		GuestRepository: guestRepository,
		DB:              DB,
		validate:        validate,
	}
}

func (service *GuestServiceImpl) Create(ctx context.Context, request web.GuestCreateRequest) web.GuestResponse {
	err := service.validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	guest := domain.Guest{
		Name:         request.Name,
		Address:      request.Address,
		Phone_Number: request.Phone_Number,
		Email:        request.Email,
	}

	guest = service.GuestRepository.Create(ctx, tx, guest)
	return helper.ToGuestResponse(guest)
}

func (service *GuestServiceImpl) Update(ctx context.Context, request web.GuestUpdateRequest) web.GuestResponse {
	err := service.validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	guest, err := service.GuestRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	guest.Name = request.Name
	guest.Address = request.Address
	guest.Phone_Number = request.Phone_Number
	guest.Email = request.Email

	guest = service.GuestRepository.Update(ctx, tx, guest)

	return helper.ToGuestResponse(guest)
}

func (service *GuestServiceImpl) Delete(ctx context.Context, guestId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	guest, err := service.GuestRepository.FindById(ctx, tx, guestId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.GuestRepository.Delete(ctx, tx, guest)
}

func (service *GuestServiceImpl) FindById(ctx context.Context, guestId int) web.GuestResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	guest, err := service.GuestRepository.FindById(ctx, tx, guestId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToGuestResponse(guest)
}

func (service *GuestServiceImpl) FindAll(ctx context.Context) []web.GuestResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	guests := service.GuestRepository.FindAll(ctx, tx)

	return helper.ToGuestResponses(guests)
}
