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

type DiscountServiceImpl struct {
	DiscountRepository repository.DiscountRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewDiscountService(discountRepository repository.DiscountRepository, DB *sql.DB, validate *validator.Validate) DiscountService {
	return &DiscountServiceImpl{
		DiscountRepository: discountRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

func (service *DiscountServiceImpl) Create(ctx context.Context, request web.DiscountCreateRequest) web.DiscountResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	discount := domain.Discount{
		Rate:        request.Rate,
		Status:      request.Status,
		RequestDate: request.Request_Date,
	}

	discount = service.DiscountRepository.Create(ctx, tx, discount)
	return helper.ToDiscountResponse(discount)
}

func (service *DiscountServiceImpl) Update(ctx context.Context, request web.DiscountUpdateRequest) web.DiscountResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	discount, err := service.DiscountRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	discount.Rate = request.Rate
	discount.Status = request.Status
	discount.RequestDate = request.RequestDate

	discount = service.DiscountRepository.Update(ctx, tx, discount)

	return helper.ToDiscountResponse(discount)

}

func (service *DiscountServiceImpl) Delete(ctx context.Context, discountId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	discount, err := service.DiscountRepository.FindById(ctx, tx, discountId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.DiscountRepository.Delete(ctx, tx, discount)
}

func (service *DiscountServiceImpl) FindById(ctx context.Context, discountId int) web.DiscountResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	discount, err := service.DiscountRepository.FindById(ctx, tx, discountId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return helper.ToDiscountResponse(discount)
}

func (service *DiscountServiceImpl) FindAll(ctx context.Context) []web.DiscountResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	discounts := service.DiscountRepository.FindAll(ctx, tx)

	return helper.ToDiscountResponses(discounts)
}