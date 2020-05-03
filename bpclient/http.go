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
func (api *OmniAPI) Authorize(a AuthorizationRequest) (result AuthorizationResponse, err error) {
	client := resty.New().
		SetHostURL(api.BaseURL).
		SetAuthToken(api.Token.AccessToken)
	resp, err := client.R().
		SetBody(a).
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
	result = *resp.Result().(*AuthorizationResponse)
	return
}

// ConfirmAuthorization confirms an authorization (or payment)
func (api *OmniAPI) ConfirmAuthorization(paymentID string) (err error) {
	client := resty.New().
		SetHostURL(api.BaseURL).
		SetAuthToken(api.Token.AccessToken)
	resp, err := client.R().
		SetBody(map[string]string{}).
		SetResult(&AuthorizationResponse{}).
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
	return
}
