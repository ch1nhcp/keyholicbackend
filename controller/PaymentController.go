package controller

import (
	"encoding/json"
	"finalbackend/models"
	"finalbackend/repository"
	"io/ioutil"
	"net/http"

	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
)

func Rabit(writer http.ResponseWriter, request *http.Request) {

}
func Payment(writer http.ResponseWriter, request *http.Request) {
	var payment models.Charge
	requestBody, _ := ioutil.ReadAll(request.Body)
	json.Unmarshal(requestBody, &payment)
	json.NewEncoder(writer).Encode(payment)
	apiKey := "sk_test_51JFCbdJA4XI9xLpDaECRznvhXIXRhJ517Q9HAY9TfLGyvoIveTiD2zK15OD77D8JtL8bNR8jr5gq5FGV095IJ7jJ00QJTwvbQE"
	// return
	stripe.Key = apiKey
	pi, err := charge.New(&stripe.ChargeParams{
		Amount:       stripe.Int64(payment.Amount),
		Currency:     stripe.String(string(stripe.CurrencyUSD)),
		Description:  stripe.String(payment.ProductName),
		Source:       &stripe.SourceParams{Token: stripe.String("tok_visa")},
		ReceiptEmail: stripe.String(payment.ReceiptEmail),
	})
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode("Payment Unsuccessfull")
		return
	}
	err1 := repository.SavePayment(&payment)
	if err1 != nil {
		json.NewEncoder(writer).Encode("error occured")
	}
	json.NewEncoder(writer).Encode(pi)

}
