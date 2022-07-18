package web

type InvoiceUpdateRequest struct {
	Id                   int    `validate:"required" json:"id"`
	Number               int    `validate:"required" json:"number"`
	EmpeloyeeId          int    `validate:"required" json:"empeloyee_id"`
	MeetingRoomPricingId int    `validate:"required" json:"meeting_room_pricing_id"`
	DiscountId           int    `validate:"required" json:"discount_id"`
	Pic                  string `validate:"required,min=1,max=100" json:"pic"`
}
