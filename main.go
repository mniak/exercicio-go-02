package main

import (
	"context"
	"fmt"
	"log"

	"github.com/mniak/desafio-curso-02/bpclient"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

const authBaseURL = "https://authsandbox.braspag.com.br"
const apiBaseURL = "https://apisandbox.cieloecommerce.cielo.com.br"

func getToken() (token *oauth2.Token, err error) {
	ctx := context.Background()
	auth := &clientcredentials.Config{
		ClientID:     "b99a463f-88db-442a-b5fa-982187b68f5c",
		ClientSecret: "VXT9EsUmN2JhszsEtRnb0bBXkUcyahsNtkkizGi+WfU=",
		Scopes:       []string{"CieloApi"},
		TokenURL:     fmt.Sprintf("%s/oauth2/token", authBaseURL),
	}
	token, err = auth.Token(ctx)
	return
}

func createPayment(api bpclient.OmniAPI, req bpclient.AuthorizationRequest) (resp bpclient.AuthorizationResponse, err error) {
	log.Println("Autorizando...")
	resp, err = api.Authorize(req)
	if err != nil {
		return
	}
	switch resp.Payment.Status {
	case bpclient.StatusNotFinished:
		err = fmt.Errorf("O pagamento não foi concluído: %s", resp.Payment.ReturnMessage)
	case bpclient.StatusDenied:
		err = fmt.Errorf("O pagamento foi negado: %s", resp.Payment.ReturnMessage)
	case bpclient.StatusAborted:
		err = fmt.Errorf("O pagamento foi abortado: %s", resp.Payment.ReturnMessage)
	case bpclient.StatusVoided:
		err = fmt.Errorf("O pagamento consta como cancelado: %s", resp.Payment.ReturnMessage)
	case bpclient.StatusPaid:
		log.Printf("Pagamento %s concluído\n", resp.Payment.PaymentID)
	case bpclient.StatusAuthorized:
		log.Printf("Pagamento %s autorizado mas não capturado\n", resp.Payment.PaymentID)
	default:
		err = fmt.Errorf("O pagamento consta com status desconhecido: %d %s", resp.Payment.Status, resp.Payment.ReturnMessage)
	}
	return
}

func confirmPayment(api bpclient.OmniAPI, pmt bpclient.Payment) (resp bpclient.ConfirmationResponse, err error) {
	log.Println("Confirmando autorização...")
	resp, err = api.ConfirmAuthorization(pmt.PaymentID)
	if err != nil {
		return
	}
	switch resp.ConfirmationStatus {
	case bpclient.Pending:
		err = fmt.Errorf("O pagamento não foi confirmado: %s", resp.ReturnMessage)
	case bpclient.Reversed:
		err = fmt.Errorf("O pagamento consta como desfeito: %s", resp.ReturnMessage)
	case bpclient.Confirmed:
		log.Printf("Pagamento %s confirmado\n", pmt.PaymentID)
	default:
		err = fmt.Errorf("A confirmação do pagamento consta com status desconhecido: %d %s", resp.ConfirmationStatus, resp.ReturnMessage)
	}
	return
}

func cancelPayment(api bpclient.OmniAPI, pmt bpclient.Payment) (resp bpclient.CancellationResponse, err error) {
	log.Println("Cancelando pagamento...")
	cancelRequest := pmt.TypedCancellationRequest()
	resp, err = api.Cancel(pmt.PaymentID, cancelRequest)
	if err != nil {
		log.Fatalln(err)
	}
	switch resp.Status {
	case bpclient.StatusNotFinished:
		err = fmt.Errorf("O pagamento consta como não concluído: %s", resp.ReturnMessage)
	case bpclient.StatusAuthorized:
		err = fmt.Errorf("O pagamento ainda consta como autorizado: %s", resp.ReturnMessage)
	case bpclient.StatusPaid:
		err = fmt.Errorf("O pagamento ainda consta como pago: %s", resp.ReturnMessage)
	case bpclient.StatusDenied:
		err = fmt.Errorf("O pagamento consta como negado: %s", resp.ReturnMessage)
	case bpclient.StatusAborted:
		err = fmt.Errorf("O pagamento consta como abortado: %s", resp.ReturnMessage)
	case bpclient.StatusVoided:
		log.Printf("Cancelamento %s efetuado\n", resp.VoidID)
	default:
		err = fmt.Errorf("O pagamento consta com status desconhecido: %d %s", resp.Status, resp.ReturnMessage)
	}
	return
}

func reverseCancellation(api bpclient.OmniAPI, pmt bpclient.Payment, cnc bpclient.CancellationResponse) (resp bpclient.CancellationReversalResponse, err error) {
	log.Println("Confirmando autorização...")
	resp, err = api.ReverseCancellation(pmt.PaymentID, cnc.VoidID)
	if err != nil {
		return
	}
	switch resp.CancellationStatus {
	case bpclient.CancellationNotFinished:
		err = fmt.Errorf("O cancelamento ainda consta como não finalizado: %s", resp.ReturnMessage)
	case bpclient.CancellationDenied:
		err = fmt.Errorf("O cancelamento consta como negado: %s", resp.ReturnMessage)
	case bpclient.CancellationAuthorized:
		err = fmt.Errorf("O cancelamento ainda consta como autorizado: %s", resp.ReturnMessage)
	case bpclient.CancellationReversed:
		log.Printf("O desfazimento do cancelamento %s foi realizado\n", cnc.VoidID)
	default:
		err = fmt.Errorf("O cancelamento consta com status desconhecido: %d %s", resp.CancellationStatus, resp.ReturnMessage)
	}
	return
}

func main() {
	token, err := getToken()
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
			Type:            bpclient.PhysicalCreditCard,
			SoftDescriptor:  "Desafio GO 2",
			PaymentDateTime: bpclient.CurrentDate(),
			Amount:          100,
			Installments:    1,
			Interest:        bpclient.InterestByMerchant,
			Capture:         true,
			ProductID:       1,
			CreditCard: bpclient.CreditCard{
				CardNumber:                     "5432123454321234",
				ExpirationDate:                 "12/2021",
				SecurityCodeStatus:             "Collected",
				SecurityCode:                   "123",
				BrandID:                        1,
				IssuerID:                       401,
				InputMode:                      bpclient.InputTyped,
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
	pmt, err := createPayment(api, pmtRequest)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = confirmPayment(api, pmt.Payment)
	if err != nil {
		log.Fatalln(err)
	}

	cnc, err := cancelPayment(api, pmt.Payment)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = reverseCancellation(api, pmt.Payment, cnc)
	if err != nil {
		log.Fatalln(err)
	}
}
