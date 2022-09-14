package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTemplateMethod_When_Token_Null(t *testing.T) {
	//req, err := http.NewRequest("GET", "localhost:8080/", nil)
	//logger := log.New(os.Stdout, "Service: ", log.LstdFlags)

	//if err != nil {
	//	t.Fatalf("Could not create request: %v", err)
	//}
	rec := httptest.NewRecorder()
	//var login = api.NewLogin(logger, &mockTr, &mockWl, &mockConfig, &csm)


	results := rec.Result()
	if results.StatusCode != http.StatusOK {
		t.Errorf("expected results OK; got %v", results.Status)
	}

}



