package main

import "fmt"

func ExampleSum() {
	fmt.Println(Sum([]int{1, 2, 3}))
	// Output: 6
}

func ExampleWithdraw() {
	card := Card{Balance: 20_000, Currency: "USD", Active: true}
	Withdraw(&card, 10_000)
	fmt.Println(card.Balance)

	//Output:10000
}

func ExampleActiveCards() {
	cards := []Card{
		{Balance: 1000, Currency: "USD", Active: true},
		{Balance: 1000, Currency: "USD", Active: false},
		{Balance: 1000, Currency: "USD", Active: false},
	}

	activeCards := ActiveCards(cards)
	result := len(activeCards)
	fmt.Println(result)
	//Output:1

}
