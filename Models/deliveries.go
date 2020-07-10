package models

type RequestD struct {
	// gorm.Model
	ID                     uint   `json:"id"`
	SaasOfficeID           uint   `json:"saas_office_id"`
	SaasCompanyID          uint   `json:"saas_company_id"`
	ReservationCode        string `json:"reservation_code"`
	PartnerID              string `json:"partner_id"`
	Status                 string `json:"status"`
	DriverName             string `json:"driver_name"`
	PackageTypes           string `json:"package_types"`
	EstimatePickUpDate       string `json:"estimate_pick_up_date"`
	EstimateDropOffDate      string `json:"estimate_drop_off_date"`
	AddressDropOff         string `json:"address_drop_off"`
	DropOffTimeWindowEnd   string `json:"drop_off_time_window_end"`
	PickUpTimeMargin       string `json:"pick_up_time_margin"`
	DropOffTimeMargin      string `json:"drop_off_time_margin"`
	AddressPickUp          string `json:"address_pick_up"`
	Recipient              string `json:"recipient"`
	RecipientPhoneNumber   string `json:"recipient_phone_number"`
	SenderPhoneNumber      string `json:"sender_phone_number"`
	Type                   string `json:"type"`
	AddressPickUpPostalCode  string `json:"address_pick_up_postal_code"`
	AddressDropOffPostalCode string `json:"address_drop_off_postal_code"`

}
