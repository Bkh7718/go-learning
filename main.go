package main

import "fmt"

type Card struct {
	Balance  int64
	Currency string
	Active   bool
}

func main() {
	fmt.Println("hello comeback")

	cards := []Card{
		{Balance: 5000, Currency: "USD", Active: true},
		{Balance: 2500, Currency: "USD", Active: true},
		{Balance: 8000, Currency: "USD", Active: false},
		{Balance: 1000, Currency: "USD", Active: false},
	}

	c := Card{Balance: 5000, Currency: "USD", Active: true}

	Withdraw(&c, 1000)
	fmt.Println(c.Balance)

	Withdraw(&c, 500)
	fmt.Println(c.Balance)

	active := ActiveCards(cards)

	fmt.Println(cards)
	fmt.Println(active)
	fmt.Println(len(active))

}

func Withdraw(card *Card, amount int64) {
	card.Balance -= amount
}

func ActiveCards(cards []Card) []Card {
	activeCards := []Card{}

	for _, card := range cards {
		if card.Active {
			activeCards = append(activeCards, card)

		}
	}
	return activeCards
}

func Even(nums []int) []int {

	result := []int{}

	for _, value := range nums {
		if value%2 == 0 {
			result = append(result, value)

		}
	}

	return result
}

func Sum(nums []int) int {
	sum := 0

	for _, value := range nums {
		sum = value + sum
	}
	return sum
}

func Max(nums []int) int {

	result := nums[0]

	for _, value := range nums {
		if value > result {
			result = value
		}
	}
	return result
}
