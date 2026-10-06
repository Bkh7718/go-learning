package main

import "fmt"

func ExamplePaymentSources() {
	cards := []Card{
		{Balance: 5000, Currency: "USD", Active: true, Number: "5058 xxxx xxxx 8888"},
		{Balance: 2500, Currency: "USD", Active: true, Number: "5058 xxxx xxxx 8887"},
		{Balance: 3000, Currency: "USD", Active: false, Number: "5058 xxxx xxxx 8886"},
	}

	fmt.Println(PaymentSources(cards))

	// Output: [{5058 xxxx xxxx 8888 5000} {5058 xxxx xxxx 8887 2500}]
}
