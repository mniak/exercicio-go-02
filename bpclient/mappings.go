package bpclient

import "time"

// TypedCancellationRequest creates a cancellation request from payment
func (pmt *Payment) TypedCancellationRequest() CancellationRequest {
	return CancellationRequest{
		Card: Card{
			InputMode:  InputTyped,
			CardNumber: pmt.CreditCard.CardNumber,
		},
		MerchantVoidID:   time.Now().Format("060102150405"),
		MerchantVoidDate: CurrentDate(),
	}
}
