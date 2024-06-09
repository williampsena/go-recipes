package web

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/brianvoe/gofakeit/v7"
)

// Struct responsible for holding all card data, such as the holder's name and card number
type Card struct {
	HolderName string `json:"holder_name"` // card holder name
	Type       string `json:"type"`        // card type (master, visa, amex)
	Number     string `json:"number"`      // card number
	Cvv        string `json:"cvv"`         // card verification code
	Expiration string `json:"exp"`         // the expiration year + month
}

// Create a Fake Card Struct
func BuildCard() (*Card, error) {
	creditCard := gofakeit.CreditCard()

	card := Card{
		HolderName: gofakeit.Name(),
		Type:       creditCard.Type,
		Number:     creditCard.Number,
		Cvv:        creditCard.Cvv,
		Expiration: creditCard.Exp,
	}

	return &card, nil
}

// Endpoint is responsible for responding to a false card generation
func CardGeneratorEndpoint(w http.ResponseWriter, r *http.Request) {
	card, err := BuildCard()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Sorry, something wrong happened!")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(card)
}
