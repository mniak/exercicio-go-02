package main

import (
	"context"
	"fmt"
	"log"

	"github.com/mniak/desafio-curso-02/bpclient"

	"golang.org/x/oauth2/clientcredentials"
)

const authBaseURL = "https://authsandbox.braspag.com.br"
const apiBaseURL = "https://apisandbox.cieloecommerce.cielo.com.br"

func main() {
	ctx := context.Background()
	auth := &clientcredentials.Config{
		ClientID:     "b99a463f-88db-442a-b5fa-982187b68f5c",
		ClientSecret: "VXT9EsUmN2JhszsEtRnb0bBXkUcyahsNtkkizGi+WfU=",
		Scopes:       []string{"CieloApi"},
		TokenURL:     fmt.Sprintf("%s/oauth2/token", authBaseURL),
	}
	token, err := auth.Token(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	api := bpclient.OmniAPI{
		BaseURL: apiBaseURL,
		Token:   *token,
	}
	pmtRequest := bpclient.AuthorizationRequest{
		MerchantOrderID: "1587997030607",
		Payment: bpclient.Payment{
			Type:            "PhysicalCreditCard",
			SoftDescriptor:  "Desafio GO 2",
			PaymentDateTime: "2020-01-08T11:00:00",
			Amount:          100,
			Installments:    1,
			Interest:        "ByMerchant",
			Capture:         true,
			ProductID:       1,
			CreditCard: bpclient.CreditCard{
				CardNumber:                     "5432123454321234",
				ExpirationDate:                 "12/2021",
				SecurityCodeStatus:             "Collected",
				SecurityCode:                   "123",
				BrandID:                        1,
				IssuerID:                       401,
				InputMode:                      "Typed",
				AuthenticationMethod:           "NoPassword",
				TruncateCardNumberWhenPrinting: true,
			},
			PinPadInformation: bpclient.PinPadInformation{
				PhysicalCharacteristics: "PinPadWithChipReaderWithoutSamAndContactless",
				ReturnDataInfo:          "00",
				SerialNumber:            "0820471929",
				TerminalID:              "42004558",
			},
		},
	}
	log.Println("Autorizando...")
	pmt, err := api.Authorize(pmtRequest)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Pagamento %s feito\n", pmt.Payment.PaymentID)

	log.Println("Confirmando autorização...")
	err = api.ConfirmAuthorization(pmt.Payment.PaymentID)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Confirmação do pagamento %s feito\n", pmt.Payment.PaymentID)
}
