package repository

import (
	"context"
	"database/sql"
	"github.com/agisnur24/booking_hotel_system.git/model/domain"
)

type MeetingRoomPricingRepository interface {
	Create(ctx context.Context, tx *sql.Tx, meetingRoomPricing domain.MeetingRoomPricing) domain.MeetingRoomPricing
	Update(ctx context.Context, tx *sql.Tx, meetingRoomPricing domain.MeetingRoomPricing) domain.MeetingRoomPricing
	Delete(ctx context.Context, tx *sql.Tx, meetingRoomPricing domain.MeetingRoomPricing)
	FindById(ctx context.Context, tx *sql.Tx, meetingRoomPricingId int) (domain.MeetingRoomPricing, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.MeetingRoomPricing
}
