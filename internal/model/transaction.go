package model

type CreateTransactionRequest struct {
	SenderID   string `json:"sender_id" validate:"uuid,required"`
	ReceiverID string `json:"receiver_id" validate:"uuid,required"`
	Amount     int    `json:"amount" validate:"gte=0,required"`
}
