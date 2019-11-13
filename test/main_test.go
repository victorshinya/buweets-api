package test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/victorshinya/buweets-api/handler"
)

var entry = "Eu gostaria de expressa a minha raiva sobre esse serviço que não funciona direito!"

func TestGetEmotion(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/get-emotion?text="+entry, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handler.GetEmotion)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned an unexpected status code. Expected %d. Got: %d.", http.StatusOK, status)
	}

	expected := "{\"language\":\"en\",\"usage\":{\"features\":2,\"text_characters\":76,\"text_units\":1},\"emotion\":{\"document\":{\"emotion\":{\"anger\":0.161509,\"disgust\":0.046091,\"fear\":0.039121,\"joy\":0.043777,\"sadness\":0.390525}}},\"sentiment\":{\"document\":{\"label\":\"negative\",\"score\":-0.836939}}}"
	if strings.EqualFold(expected, rr.Body.String()) {
		t.Errorf("Handler returned an unexpected body. Expected: %s. Got: %s.", expected, rr.Body.String())
	}
}
