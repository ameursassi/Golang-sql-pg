package main

import (
	router "github/map_dashboard/Routers"
	_ "github/map_dashboard/docs"

	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/subosito/gotenv"
	httpSwagger "github.com/swaggo/http-swagger"
)

func Init() {
	gotenv.Load()
}

// @title bookings API
// @version 1.0
// @description This is a sample service for managing orders
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email soberkoder@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost/8080
// @BasePath /
func main() {
	Init()
	router := router.InitializeRouter()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	err := http.ListenAndServe(":"+port, handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT","DELETE", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"https://intense-hamlet-72331.herokuapp.com", "http://localhost:3000"}))(router))
	fmt.Println(err)
}
