package tests

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"vehicle-tax/api"
)

func TestHeartbeat(t *testing.T) {
	req, err := http.NewRequest("GET", "localhost:8080/heartbeat", nil)
	logger := log.New(os.Stdout, "Service: ", log.LstdFlags)

	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}
	rec := httptest.NewRecorder()
	api.NewHeartbeat(logger).Heartbeat(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status is OK, 200 but got %v", res.Status)
	}
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("Could not read response: %v", err)
	}
	if string(b) != "I'm Alive!!!!!" {
		t.Errorf("Expected \"I'm Alive!!!!!\" but got %v", string(b))
	}
}
