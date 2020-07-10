package router

import (
	controllers "github/map_dashboard/controllers"

	"github.com/gorilla/mux"
)

func InitializeRouter() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/login", controllers.Auth).Methods("POST")
	router.HandleFunc("/bookings/{saas_office_id}", controllers.GetBooking).Methods("GET")
	router.HandleFunc("/deliveries/{saas_office_id}", controllers.GetDeliveries).Methods("GET")
	router.HandleFunc("/offices", controllers.GetOffices).Methods("GET")
	router.HandleFunc("/offices/{id}", controllers.GetOffice).Methods("GET")
	router.HandleFunc("/drivers/{saas_office_id}", controllers.GetDrivers).Methods("GET")
	router.HandleFunc("/dispatchdrivers/{saas_office_id}", controllers.GetDispatchDriver).Methods("GET")
	router.HandleFunc("/statedrivers/{saas_office_id}/{driver_id}", controllers.GetDriverId).Methods("GET")
	router.HandleFunc("/drivers/{driver_id}/{saas_office_id}/{id}", controllers.AffectedDriver).Methods("PUT")
	router.HandleFunc("/trajectory/{id}", controllers.GetTrajectory).Methods("GET")
	router.HandleFunc("/details/{id}", controllers.GetDetails).Methods("GET")
	router.HandleFunc("/bookings/{saas_office_id}/{id}", controllers.DeleteBooking).Methods("DELETE")
	router.HandleFunc("/deliveries/{saas_office_id}/{id}", controllers.DeleteDeliveries).Methods("DELETE")
	router.HandleFunc("/drivers/{saas_office_id}/{id}", controllers.DeleteDriver).Methods("PUT")
	router.HandleFunc("/driversbod/{saas_office_id}/{id}", controllers.DeleteDriverBod).Methods("PUT")

	return router
}
