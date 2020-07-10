package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func checkError(err error, t *testing.T) {
	if err != nil {
		t.Errorf("An error occurred. %v", err)
	}
}

func TestGetDeliveries(t *testing.T) {
	req, err := http.NewRequest("GET", "/deliveries/146", nil)
	checkError(err, t)
	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/deliveries/{saas_company_id}", GetDeliveries).Methods("GET")
	r.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
	expected := string(`[{ "id": 359283,"saas_company_id": 146,"reservation_code": "J44YKM","partner_id": "","status": "created","driver_name": "-","package_types": "Humain: 1","pick_up_time_window_begin": "01/02/2020 09:31:40","drop_off_time_window_begin": "01/02/2020 09:31:40","pick_up_time_window_end": "01/02/2020 13:18:20","address_drop_off": "21 Rue Saint-Fiacre, Paris, France","drop_off_time_window_end": "01/02/2020 13:18:58","pick_up_time_margin": "113","drop_off_time_margin": "113","address_pick_up": "15 Rue Saint-Fiacre, Paris, France","recipient": "Guillaume Deberdt","recipient_phone_number": "+33686077352","sender_phone_number": "+33686077352","type": "DeliveryRequest" },
    { "id": 359284,"saas_company_id": 146,"reservation_code": "833Z5J","partner_id": "","status": "wait_pick_up","driver_name": "Guigui Thebest","package_types": "Humain: 1, Nenfant: 1","pick_up_time_window_begin": "01/02/2020 11:25:00","drop_off_time_window_begin": "01/02/2020 10:45:22","pick_up_time_window_end": "01/02/2020 11:40:00","address_drop_off": "21 Rue Saint-Fiacre, Paris, France","drop_off_time_window_end": "01/02/2020 13:15:22","pick_up_time_margin": "14575","drop_off_time_margin": "105","address_pick_up": "15 Rue Saint-Fiacre, Paris, France","recipient": "Guillaume Deberdt","recipient_phone_number": "+33686077352","sender_phone_number": "+33686077352","type": "DeliveryRequest"}]`)
	assert.JSONEq(t, expected, rr.Body.String(), "Response body differs")

}
