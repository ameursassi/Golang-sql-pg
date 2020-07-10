package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestGetOffices(t *testing.T) {
	req, err := http.NewRequest("GET", "/offices", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetOffices)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler resturned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestGetOffice(t *testing.T) {
	req, err := http.NewRequest("GET", "/offices/1", nil)
	checkError(err, t)
	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/offices/{saas_company_id}", GetOffice).Methods("GET")
	r.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
	expected := string(`{
		"id": 1,
		"saas_company_id": 1,
		"name": "Marcel-paris",
		"allow_delivery": false
	}`)
	assert.JSONEq(t, expected, rr.Body.String(), "Response body differs")

}
