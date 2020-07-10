package models

import (
	"github.com/lib/pq"
)

type RequestDetails struct {
	ID           uint `json:"id"`
	SaasOfficeID uint `json:"saas_office_id"`

	AddressPickUp        string         `json:"address_pick_up"`
	AddressDropOff       string         `json:"address_drop_off"`
	RideDate             string         `json:"ride_date"`
	State                string         `json:"state"`
	DriverPrice          uint           `json:"driver_price"`
	OptionsPrice         uint           `json:"options_price"`
	FlightNumber         string         `json:"flight_number"`
	Comment              string         `json:"comment"`
	PaymentType          string         `json:"payment_type"`
	Designation          string         `json:"designation"`
	RequestOptions       pq.StringArray `json:"request_options"`
	IntermediateSteps    pq.StringArray `json:"intermediate_steps"`
	PassengerFullName    string         `json:"passenger_full_name"`
	PassengerPhoneNumber string         `json:"passenger_phone_number"`
	DriverID             uint           `json:"driver_id"`
	DriverName           string         `json:"driver_name"`
	DriverPhone          string         `json:"driver_phone"`
	DriverStatus         string         `json:"driver_status"`
	Car                  string         `json:"car"`
	Type                 string         `json:"type"`
	EstimatePickUpDate   string         `json:"estimate_pick_up_date"`
	EstimateDropOffDate  string         `json:"estimate_drop_off_date"`
	RidePrice            float64        `json:"ride_price"`
	PackageTypes         string         `json:"package_types"`
	Recipient            string         `json:"recipient"`
	RecipientPhoneNumber string         `json:"recipient_phone_number"`
	Company              string         `json:"company"`
}

// type RequestOptions struct {
// 	RequestOptions string `json:"request_options"`
// }
