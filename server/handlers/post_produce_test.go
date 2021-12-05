package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gannet/services"
	"gannet/services/produce"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPostProduce(t *testing.T) {
	url := fmt.Sprintf("%s:%d%s", "http://0.0.0.0", 4000, "/api/v1/produce")

	items := []produce.ProduceItem{
		{
			ProduceCode: "A12U-4GH7-QPL9-3N4M",
			Name:        "Tomato",
			UnitPrice:   4.46,
		},
		{
			ProduceCode: "E5T1-9UI3-TH15-QR88",
			Name:        "Peach Pits",
			UnitPrice:   7.99,
		},
		{

			ProduceCode: "YRT2-72AS-K736-L4AR",
			Name:        "Red Pepper",
			UnitPrice:   0.39,
		},
		{

			ProduceCode: "TQ4D-VV6T-75ZX-1RMR",
			Name:        "Green Apple",
			UnitPrice:   3.98,
		},
	}

	var payload bytes.Buffer
	err := json.NewEncoder(&payload).Encode(items)

	if err != nil {
		t.Fatalf("Failed to run test due to json.NewEncoder(bytes.Buffer).Endcode(items) failure. reason: %s\n", err)
	}

	req := httptest.NewRequest(http.MethodPost, url, &payload)
	w := httptest.NewRecorder()

	PostProduce(w, req)

	res := w.Result()

	if http.StatusOK != res.StatusCode {
		t.Fatalf("expected status 200, got %d\n", res.StatusCode)
	}

	store := services.FetchAll()

	for _, item := range items {
		found := false

		for _, dbItem := range store {
			if dbItem.Name == item.Name &&
				dbItem.ProduceCode == item.ProduceCode &&
				dbItem.UnitPrice == item.UnitPrice {
				found = true
				break
			}
		}

		if !found {
			t.Fatalf("recieved: %v, which is not in database\n", item)
		}
	}

}
