package models

type Request struct {
	// gorm.Model
	// tableName                struct{} `pg:"requests"`
	ID uint `json:"id"`

	SaasOfficeID             uint    `json:"saas_office_id"`
	SaasCompanyID            uint    `json:"saas_company_id"`
	EstimateDistanceM        uint    `json:"estimate_distance_m"`
	ReservationCode          string  `json:"reservation_code"`
	EstimatePickUpDate       string  `json:"estimate_pick_up_date"`
	EstimateDropOffDate      string  `json:"estimate_drop_off_date"`
	FireTime                 string  `json:"fire_time"`
	RideDate                 string  `json:"ride_date"`
	State                    string  `json:"state"`
	AddressPickUpPostalCode  string  `json:"address_pick_up_postal_code"`
	AddressPickUp            string  `json:"address_pick_up"`
	AddressDropOffPostalCode string  `json:"address_drop_off_postal_code"`
	AddressDropOff           string  `json:"address_drop_off"`
	AddressPickUpLat         float64 `json:"address_pick_up_lat"`
	AddressPickUpLong        float64 `json:"address_pick_up_long"`
	AddressDropOffLat        float64 `json:"address_drop_off_lat"`
	AddressDropOffLong       float64 `json:"address_drop_off_long"`
	PassengerFullName        string  `json:"passenger_full_name"`
	PassengerPhoneNumber     string  `json:"passenger_phone_number"`
	FlightNumber             string  `json:"flight_number"`
	DriverID                 string  `json:"driver_id"`
	DriverName               string  `json:"driver_name"`
	DriverPhone              string  `json:"driver_phone"`
	Company                  string  `json:"company"`
	Type                     string  `json:"type"`
}

type SaasOffices struct {
	ID               uint    `json:"id"`
	SaasCompanyID    uint    `json:"saas_company_id"`
	Name             string  `json:"name"`
	AllowDelivery    bool    `json:"allow_delivery"`
	OfficeCenterLat  float64 `json:"office_center_lat"`
	OfficeCenterLong float64 `json:"office_center_long"`
	OfficeRadius     int     `json:"office_radius"`
	DistanceUnit     string  `json:"distance_unit"`
}
