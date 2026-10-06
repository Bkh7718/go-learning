package main

import "fmt"

func main() {

	cards := []Card{
		{Balance: 5000, Currency: "USD", Active: true, Number: "5058 xxxx xxxx 8888"},
		{Balance: 2500, Currency: "USD", Active: true, Number: "5058 xxxx xxxx 8887"},
		{Balance: 3000, Currency: "USD", Active: false, Number: "5058 xxxx xxxx 8886"},
	}
	active := ActiveCards(cards)

	c := Card{
		Balance: 5000, Currency: "USD", Active: true,
	}

	fmt.Println(c.Balance)
	fmt.Println(active)

	Withdraw(&c, 1000)
	fmt.Println(c.Balance)

	Withdraw(&c, 500)
	fmt.Println(c.Balance)

	sources := PaymentSources(cards)
	fmt.Println(sources)

}

type PaymentSource struct {
	Number  string
	Balance int64
}

type Card struct {
	Balance  int64
	Currency string
	Active   bool
	Number   string
}

func ActiveCards(cards []Card) []Card {
	active := []Card{}

	for _, card := range cards {
		if card.Active {
			active = append(active, card)

		}

	}
	return active
}

func Withdraw(card *Card, amount int64) {
	card.Balance -= amount

}

func PaymentSources(cards []Card) []PaymentSource {
	sources := []PaymentSource{}

	for _, card := range cards {
		if card.Active {
			sources = append(sources, PaymentSource{Balance: card.Balance, Number: card.Number})

		}
	}
	return sources
}
