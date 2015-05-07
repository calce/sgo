package main

import (
	"github.com/calce/sgo"
	"github.com/calce/sgo/payments"
	"fmt"
)

func formatMoney(m sgo.Money) string {
	return fmt.Sprint("$", m.Amount / 100)
}

func main() {
	sgo.Token = "Secret"
	sgo.MerchantId = "me"
	
	params := sgo.PaymentListParams{}
	payments, err := payments.List(&params)
	
	if err != nil { 
		if err.ApiError != nil { fmt.Println(*err.ApiError) }
		if err.ParseError != nil { fmt.Println(err.ParseError) }
		return
	}
	
	index := 0
	for payments.HasNext() {
		payment, _ := payments.Next()
		fmt.Println("Payment", index, payment.Id, formatMoney(payment.NetTotalMoney))
		index++
	}
	
	fmt.Println(index, "payments listed")
}