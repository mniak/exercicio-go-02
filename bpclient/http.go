package bpclient

import (
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
	"golang.org/x/oauth2"
)

// OmniAPI api
type OmniAPI struct {
	BaseURL string
	Token   oauth2.Token
}

// Authorize creates a new payment
func (api *OmniAPI) Authorize(req AuthorizationRequest) (result AuthorizationResponse, err error) {
	client := resty.New().
		SetHostURL(api.BaseURL).
		SetAuthToken(api.Token.AccessToken)
	resp, err := client.R().
		SetBody(req).
		SetError(&[]Error{}).
		SetResult(&AuthorizationResponse{}).
		Post("/1/physicalSales")
	if err != nil {
		return
	}
	if s := resp.StatusCode(); s != 201 {
		errors := resp.Error().([]Error)
		log.Printf("Authorization errors: %v\n", errors)
		err = fmt.Errorf("Invalid status code: %d", s)
		return
	}
	if resp.IsSuccess() {
		result = *resp.Result().(*AuthorizationResponse)
	}
	return
}

// ConfirmAuthorization confirms an authorization (or payment)
func (api *OmniAPI) ConfirmAuthorization(paymentID string) (result ConfirmationResponse, err error) {
	client := resty.New().
		SetHostURL(api.BaseURL).
		SetAuthToken(api.Token.AccessToken)
	resp, err := client.R().
		SetBody(map[string]string{}).
		SetError(&[]Error{}).
		SetResult(&ConfirmationResponse{}).
		Put(fmt.Sprintf("/1/physicalSales/%s/confirmation", paymentID))

	if err != nil {
		return
	}
	if resp.IsError() {
		if e := resp.Error(); e != nil {
			errors := e.([]Error)
			log.Printf("Confirmation errors: %v\n", errors)
		}
		err = fmt.Errorf("Error response with status %d", resp.StatusCode())
	}
	if resp.IsSuccess() {
		result = *resp.Result().(*ConfirmationResponse)
	}
	return
}

// Cancel cancels a payment
func (api *OmniAPI) Cancel(paymentID string, req CancellationRequest) (result CancellationResponse, err error) {
	client := resty.New().
		SetHostURL(api.BaseURL).
		SetAuthToken(api.Token.AccessToken)
	resp, err := client.R().
		SetBody(req).
		SetError(&[]Error{}).
		SetResult(&CancellationResponse{}).
		Post(fmt.Sprintf("/1/physicalSales/%s/voids", paymentID))

	if err != nil {
		return
	}
	if resp.IsError() {
		b := resp.Body()
		log.Print("Body")
		log.Print(string(b))
		if e := resp.Error(); e != nil {
			errors := e.([]Error)
			log.Printf("Cancellation errors: %v\n", errors)
		}
		err = fmt.Errorf("Error response with status %d", resp.StatusCode())
	}
	if resp.IsSuccess() {
		result = *resp.Result().(*CancellationResponse)
	}
	return
}

func (api *OmniAPI) ReverseCancellation(paymentID, voidID string) (result CancellationReversalResponse, err error) {
	client := resty.New().
		SetHostURL(api.BaseURL).
		SetAuthToken(api.Token.AccessToken)
	resp, err := client.R().
		SetBody(map[string]string{}).
		SetError(&[]Error{}).
		SetResult(&CancellationReversalResponse{}).
		Delete(fmt.Sprintf("/1/physicalSales/%s/voids/%s", paymentID, voidID))

	if err != nil {
		return
	}
	if resp.IsError() {
		if e := resp.Error(); e != nil {
			errors := e.([]Error)
			log.Printf("Cancellation reveral errors: %v\n", errors)
		}
		err = fmt.Errorf("Error response with status %d", resp.StatusCode())
	}
	if resp.IsSuccess() {
		result = *resp.Result().(*CancellationReversalResponse)
	}
	return
}
