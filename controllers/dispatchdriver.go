package controllers

import (
	"encoding/json"
	"fmt"
	helpers "github/map_dashboard/Helpers"
	models "github/map_dashboard/Models"
	"strconv"

	"github.com/gorilla/mux"

	"net/http"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const querydispatchdriver = `SELECT
saas_office_to_driver_relation.driver_id                                         "driver_id",
saas_office_to_driver_relation.saas_office_id                                    "saas_office_id" ,                                                          
saas_office_to_driver_relation.driver_status                                   "driver_status"  ,                          
vehicle_type.designation                                                        "designation",
CASE WHEN driver.id IS NULL THEN '-'
ELSE CONCAT(back_end_user.first_name || ' ' || back_end_user.last_name)  END   "driver_name",
back_end_user.phone_number                                                         "driver_phone"
FROM
saas_office_to_driver_relations 
LEFT OUTER JOIN drivers driver on driver.id = saas_office_to_driver_relations.driver_id 
LEFT OUTER JOIN back_end_users back_end_user on back_end_user.driver_id = saas_office_to_driver_relations.driver_id 
LEFT OUTER JOIN saas_office_to_driver_relations saas_office_to_driver_relation on saas_office_to_driver_relation.driver_id = driver.id 
LEFT OUTER JOIN vehicle_types vehicle_type on vehicle_type.id=saas_office_to_driver_relations.driver_id  
WHERE
saas_office_to_driver_relation.last_state_date >= date_trunc('month', now()) - interval '25 month'`

// @Getdispatchdriver list godoc
// @Summary Get dispatchdriver for a given dispatchdriver list
// @Description Get details of list dispatchdriver
// @Tags dispatchdriver
// @Accept  json
// @Produce  json
// @Success 200 {object} models.SaasOffices
// @Router /offices [get]
func GetDispatchDriver(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := helpers.DbConnect()
	var driverdispatch []models.DriverDispatch
	err1 := db.Raw(querydispatchdriver).Scan(&driverdispatch).Error
	if err1 != nil {
		fmt.Println(err1)
	}
	params := mux.Vars(r)
	result := make([]models.DriverDispatch, 0)

	for _, item := range driverdispatch {
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


// @GetOffice godoc
// @Summary Get details for a given officeId
// @Description Get details of office corresponding to the input officeId
// @Tags offices
// @Accept  json
// @Produce  json
// @Param saas_office_id path int true "ID of the offices"
// @Success 200 {object} models.SaasOffices
// @Router /offices/{saas_office_id} [get]
func GetDriverId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := helpers.DbConnect()
	var driverdispatch []models.DriverDispatch
	err1 := db.Raw(querydispatchdriver).Scan(&driverdispatch).Error
	if err1 != nil {
		fmt.Println(err1)
	}
	params := mux.Vars(r)
	for _, item := range driverdispatch {
		driver_id, err := strconv.ParseUint(params["driver_id"], 10, 32)
		if err != nil {
			fmt.Println(err)
		}
		if uint(driver_id) == item.DriverID {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

}