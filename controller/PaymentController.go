package controller

import (
	"encoding/json"
	"finalbackend/models"
	"finalbackend/rabitmq"
	"finalbackend/repository"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
)

func Rabit(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(writer.Header().Get("Content-Type"))
	rmq := rabitmq.RabbitMQ{
		ConnectionString: "amqp://tfs:tfs-ocg@174.138.40.239:5672/",
	}
	rmq.CreateConnection()
	defer rmq.Close()
	fmt.Println("Successfuly Connected To our RMQ Instance")
	ch := rmq.GetChannel()
	defer ch.Close()
	q, err := ch.QueueDeclare(
		"tung", // name
		false,  // durable
		false,  // delete when unused
		false,  // exclusive
		false,  // no-wait
		nil,    // arguments
	)
	rabitmq.FailOnError(err, "Failed to declare a queue")
	body := "Bui Duy Tung"
	rabitmq.Publish(ch, q.Name, body)
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
