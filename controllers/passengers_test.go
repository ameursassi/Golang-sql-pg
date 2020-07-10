package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestGetBooking(t *testing.T) {
	req, err := http.NewRequest("GET", "/bookings/2", nil)
	checkError(err, t)
	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/bookings/{saas_company_id}", GetBooking).Methods("GET")
	r.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
	expected := string(`[ ]`)
	assert.JSONEq(t, expected, rr.Body.String(), "Response body differs")

}
