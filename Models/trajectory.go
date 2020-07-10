package models

type Trajectory struct {
	ID                 uint    `json:"id"`
	AddressPickUpLat   float64 `json:"address_pick_up_lat"`
	AddressPickUpLong  float64 `json:"address_pick_up_long"`
	AddressDropOffLat  float64 `json:"address_drop_off_lat"`
	AddressDropOffLong float64 `json:"address_drop_off_long"`
	// Path               []Path  `json:"path"`
}

type Path struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type ReturnPath struct {
	Distance          float64 `json:"distance"`
	Duration          float64 `json:"duration"`
	DurationInTraffic float64 `json:" duration_in_traffic"`
	Path              []Path  `json:"path"`
}
