package handlers

import (
	"encoding/json"
	"fmt"
	"gannet/services"
	"gannet/services/produce"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetProduce(t *testing.T) {
	url := fmt.Sprintf("%s:%d%s", "http://0.0.0.0", 4000, "/api/v1/produce")

	req := httptest.NewRequest(http.MethodGet, url, nil)
	w := httptest.NewRecorder()

	GetProduce(w, req)

	res := w.Result()

	// Check status
	if http.StatusOK != res.StatusCode {
		t.Fatalf("expected a 200, instead got: %d\n", res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		t.Fatalf("unable to read response, reason: %s\n", err)
	}

	var data []produce.ProduceItem
	err = json.Unmarshal(body, &data)

	if err != nil {
		t.Fatalf("unable to unmarshal response, reason: %s\n", err)
	}

	store := services.FetchAll()

	for _, item := range data {
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
