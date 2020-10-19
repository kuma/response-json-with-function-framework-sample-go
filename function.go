package hello

import (
	"encoding/json"
	"net/http"
	"fmt"
)

type user struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Age int `json:"age"`
	DefaultPaymentMethod string `json:"default_payment_method"`
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	users := []user{
		{
			ID: 10297648,
			Name: "Kuma Arakawa",
			Age: 30,
			DefaultPaymentMethod: "google_pay",
		},
		{
			ID: 10234874,
			Name: "Nanashi Gonbe",
			Age: 45,
			DefaultPaymentMethod: "visa",
		},
		{
			ID: 10297498,
			Name: "Naoki Smith",
			Age: 41,
			DefaultPaymentMethod: "amex",
		},
	}

	res, _ := json.Marshal(users)

	fmt.Fprintf(w, string(res))
}
