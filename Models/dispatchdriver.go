package models

type DriverDispatch struct {
	DriverID     uint    `json:"driver_id"`
	SaasOfficeID uint    `json:"saas_office_id"`
	DriverName   string  `json:"driver_name"`
	DriverPhone  string  `json:"driver_phone"`
	DriverStatus string  `json:"driver_status"`
	Designation  string  `json:"designation"`
}
