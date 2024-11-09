package main

import (
	payment "gomall/kitex_gen/payment/paymentservice"
	"log"
)

func main() {
	svr := payment.NewServer(new(PaymentServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
