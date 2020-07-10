package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuth(t *testing.T) {
	var jsonStr = []byte(`{"email": "mohamed.ali.sfari@classnco.com", "password": "Azerty!29"}`)
	mux := http.NewServeMux()
	mux.HandleFunc("/authentification", Auth)
	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/authentification", bytes.NewBuffer(jsonStr))
	mux.ServeHTTP(writer, request)
	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
	var result map[string]interface{}
	json.Unmarshal(writer.Body.Bytes(), &result)
	fmt.Println(result["success"])
}
