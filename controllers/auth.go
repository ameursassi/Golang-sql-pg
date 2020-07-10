package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func Auth(w http.ResponseWriter, request *http.Request) {
	fmt.Println("Server running on port 8080...")
	response, err := http.Post(os.Getenv("URL"), "application/json", request.Body)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
		w.WriteHeader(response.StatusCode)
		w.Write(data)
		return
	}
}
