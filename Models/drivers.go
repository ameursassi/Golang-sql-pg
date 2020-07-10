package models

import (
	"time"
)

type Drivers struct {
	SaasOfficeID         uint      `json:"saas_office_id"`
	DriverID             uint      `json:"driver_id"`
	LastTraceLat         float64   `json:"last_trace_lat"`
	LastTraceLong        float64   `json:"last_trace_long"`
	DriverStatus         string    `json:"driver_status"`
	LastTraceDate        time.Time `json:"last_trace_date"`
	LastTrace            string    `json:"last_trace"`
	CarType              string    `json:"car_type"`
	OfficeRadius         int       `json:"office_radius"`
	MarkerMap            string    `json:"marker_map"`
	Designation          string    `json:"designation"`
	Note                 int       `json:"note"`
	DriverName           string    `json:"driver_name"`
	DriverPhone          string    `json:"driver_phone"`
	CommissionPercentage float64   `json:"commission_percentage"`
	ActionType           int       `json:"action_type"`
	Priority             int       `json:"priority"`
}
