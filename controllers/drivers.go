package controllers

import (
	"encoding/json"
	"fmt"
	helpers "github/map_dashboard/Helpers"
	models "github/map_dashboard/Models"
	"strconv"

	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const querydrivers = `SELECT
saas_office_to_driver_relation.driver_id                                         "driver_id",
saas_office_to_driver_relation.commission_percentage                             "commission_percentage",                                                       
saas_office_to_driver_relation.saas_office_id                                    "saas_office_id" ,                                                          
saas_office_to_driver_relation.last_trace_lat                                    "last_trace_lat"  ,                          
saas_office_to_driver_relation.last_trace_long                                   "last_trace_long"  ,                          
saas_office_to_driver_relation.driver_status                                   "driver_status"  ,                          
saas_office_to_driver_relation.last_trace_date                                   "last_trace_date"  ,                          
date_trunc('min', now())-saas_office_to_driver_relation.last_trace_date           "last_trace"  ,                          
CONCAT(driver_car.car_marque || ' ' ||driver_car.car_model)                    "car_type",
saas_office.office_radius                                                        "office_radius"  ,
driver_action_to_ride.action_type   "action_type",
vehicle_type.marker_map                                                        "marker_map",
vehicle_type.designation                                                        "designation",
driver_note.note                                                                  "note",
driver.priority                                                       "priority",
CASE WHEN driver.id IS NULL THEN '-'
ELSE CONCAT(back_end_user.first_name || ' ' || back_end_user.last_name)  END   "driver_name",
back_end_user.phone_number                                                         "driver_phone"		            
FROM
saas_office_to_driver_relations 
LEFT OUTER JOIN drivers driver on driver.id = saas_office_to_driver_relations.driver_id 
LEFT OUTER JOIN driver_companies driver_company on driver.driver_company_id = driver_company.id
LEFT OUTER JOIN saas_offices saas_office on office_radius = saas_office.office_radius
LEFT OUTER JOIN back_end_users back_end_user on back_end_user.driver_id = saas_office_to_driver_relations.driver_id 
LEFT OUTER JOIN driver_cars driver_car on driver_car.driver_id=saas_office_to_driver_relations.driver_id  
LEFT OUTER JOIN driver_action_to_rides driver_action_to_ride on driver_action_to_ride.driver_id=saas_office_to_driver_relations.driver_id  
LEFT OUTER JOIN vehicle_types vehicle_type on vehicle_type.id=saas_office_to_driver_relations.driver_id  
LEFT OUTER JOIN driver_notes driver_note on driver_note.id=saas_office_to_driver_relations.driver_id  
LEFT OUTER JOIN saas_office_to_driver_relations saas_office_to_driver_relation on saas_office_to_driver_relation.driver_id = driver.id 
WHERE
saas_office_to_driver_relation.last_trace_date >= date_trunc('min', now()) - interval '15 min' AND
saas_office.max_time_to_consider_driver_available_in_min=10 AND
saas_office_to_driver_relation.last_trace_lat!= 0 AND
saas_office_to_driver_relation.last_trace_long!= 0 `

// @GetDrivers godoc
// @Summary Get details for a given driverId
// @Description Get details of driver corresponding to the input driverId
// @Tags drivers
// @Accept  json
// @Produce  json
// @Param saas_office_id path int true "ID of the drivers"
// @Success 200 {object} models.Drivers
// @Router /drivers/{saas_office_id} [get]
func GetDrivers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := helpers.DbConnect()
	var drivers []models.Drivers
	err1 := db.Raw(querydrivers).Scan(&drivers).Error
	if err1 != nil {
		fmt.Println(err1)
	}
	params := mux.Vars(r)
	result := make([]models.Drivers, 0)
	for _, item := range drivers {
		saas_office_id, err := strconv.ParseUint(params["saas_office_id"], 10, 32)
		if err != nil {
			fmt.Println(err)
		}
		if uint(saas_office_id) == item.SaasOfficeID {
			result = append(result, item)
		}
	}
	json.NewEncoder(w).Encode(result)
}
