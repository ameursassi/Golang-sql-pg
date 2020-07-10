package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestGetDrivers(t *testing.T) {
	req, err := http.NewRequest("GET", "/drivers/2", nil)
	checkError(err, t)
	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/drivers/{saas_office_id}", GetDrivers).Methods("GET")
	r.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
	expected := string(`[ ]`)
	assert.JSONEq(t, expected, rr.Body.String(), "Response body differs")

}
