package jsonpzngo

import (
	"encoding/json"
	"os"
	"testing"
)

func TestEncoder(t *testing.T) {
	writer, _ := os.Create("CustomerOut.json")
	encoder := json.NewEncoder(writer)

	customer := Customer{
		FirstName:  "Luthfi",
		MiddleName: "Izzuddin",
		LastName:   "Hanif",
	}

	encoder.Encode(customer)
}
