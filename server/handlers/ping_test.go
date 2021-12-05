package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestPing
// tests that the Ping handler, from ping.go,
// works as expected with a 200 status code and with "pong" as the response body
func TestPing(t *testing.T) {
	url := fmt.Sprintf("%s:%d%s", "http://0.0.0.0", 4000, "/api/v1/produce")

	req := httptest.NewRequest(http.MethodGet, url, nil)
	w := httptest.NewRecorder()

	Ping(w, req)

	res := w.Result()

	// Check status
	if http.StatusOK != res.StatusCode {
		t.Fatalf("expected a 200, instead got: %d\n", res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		t.Fatalf("unable to read response, reason: %s\n", err)
	}

	msg := string(body)

	if msg != "pong" {
		t.Fatalf("Expected \"pong\" as response, got: \"%s\"\n", msg)
	}
}
