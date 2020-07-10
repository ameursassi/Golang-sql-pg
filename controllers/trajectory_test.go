package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestGetTrajectory(t *testing.T) {
	req, err := http.NewRequest("GET", "/trajectory/356084", nil)
	checkError(err, t)
	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/trajectory/{id}", GetTrajectory).Methods("GET")
	r.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
	expected := string(`{
		"LatOrigin": 48.8630776,
		"LongOrigin": 2.3590373,
		"LatDestination": 48.8795376,
		"LongDestination": 2.35549820000006,
		"Path": {
			"distance": 2347,
			"duration": 766,
			" duration_in_traffic": 0,
			"path": [
				{"lat": 48.8630345,	"lng": 2.3589802},{"lat": 48.8633709,"lng": 2.3582523},{"lat": 48.8672778,"lng": 2.3633596},{"lat": 48.8679289,"lng": 2.3625209},{"lat": 48.8684648,"lng": 2.3630858},{"lat": 48.87863730000001,"lng": 2.3540497},{"lat": 48.8798685,"lng": 2.3547182}]
		} 
	}`)
	assert.JSONEq(t, expected, rr.Body.String(), "Response body differs")

}
