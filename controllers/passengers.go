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

const querypassenger = ` SELECT
    request.id                                                  "id" ,                                                       
    request.saas_office_id                                      "saas_office_id" ,                                                          
	request.saas_company_id                                     "saas_company_id" ,                                                          
	request.estimate_distance_m                                 "estimate_distance_m"  ,                          
	request.reservation_code                                    "reservation_code"  ,                          
	request.customer_id                                         "customer_id",
	to_char(request.estimate_pick_up_date AT TIME ZONE 'UTC')   "estimate_pick_up_date",
	to_char(request.estimate_drop_off_date AT TIME ZONE 'UTC')  "estimate_drop_off_date",
	to_char(request.fire_time AT TIME ZONE 'UTC')               "fire_time",
	request.ride_date                                           "ride_date",
	request.state,                                                           
	request.address_pick_up_postal_code,                                    
	request.address_pick_up,                                                 
	request.address_drop_off_postal_code,                                   
	request.address_drop_off,                                                
	request.address_pick_up_lat,
	request.address_pick_up_long,
	request.address_drop_off_lat,
	request.address_drop_off_long,
	CASE WHEN real_passenger.id IS NULL
		THEN passenger.first_name || ' ' || passenger.last_name ELSE
		  real_passenger.first_name || ' ' || real_passenger.last_name END  "passenger_full_name",
	CASE WHEN real_passenger.id IS NULL
		THEN passenger.phone_number ELSE
		  real_passenger.phone_number END                                   "passenger_phone_number",
		  COALESCE(request.flight_number, '')                                     "flight_number",
		  request.driver_id                                                       "driver_id",
	CASE WHEN driver.id IS NULL THEN '-'
	  ELSE CONCAT(back_end_user.first_name || ' ' || back_end_user.last_name)  END   "driver_name",
  back_end_user.phone_number                                                    "driver_phone",			            
	companies.name                                                          "company",
	 request.type                                                             "type"
	FROM
	requests request
	LEFT OUTER JOIN passengers passenger on passenger.id = request.passenger_id
	LEFT OUTER JOIN customers customer on customer.id = request.customer_id
    LEFT OUTER JOIN customer_statistics stats on stats.customer_id = customer.id
	LEFT OUTER JOIN drivers driver on driver.id = request.driver_id
	LEFT OUTER JOIN back_end_users back_end_user on back_end_user.driver_id = driver.id
	LEFT OUTER JOIN flights flight on flight.id = request.flight_id
	LEFT OUTER JOIN requests outsourced_origin_request on outsourced_origin_request.id = request.outsourced_origin_request_id
	LEFT OUTER JOIN passengers real_passenger on real_passenger.id = outsourced_origin_request.passenger_id
	 LEFT OUTER JOIN saas_office_to_driver_relations saas_office_to_driver_relation on saas_office_to_driver_relation.driver_id = driver.id AND saas_office_to_driver_relation.saas_office_id = request.saas_company_id
	LEFT OUTER JOIN companies on companies.id = request.company_id
	WHERE
   customer.saas_company_id = request.saas_company_id AND
  request.type = 'PassengerRequest' AND
  request.ride_date >= date_trunc('month', now()) - interval '2 month' 
  ORDER BY request.ride_date ASC`

// @GetBooking godoc
// @Summary Get details for a given passengerId
// @Description Get details of passenger corresponding to the input passengerId
// @Tags bookings
// @Accept  json
// @Produce  json
// @Param saas_office_id path int true "ID of the passengers"
// @Success 200 {object} models.RequestP
// @Router /bookings/{saas_office_id} [get]
func GetBooking(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := helpers.DbConnect()
	var requests []models.Request
	err1 := db.Raw(querypassenger).Scan(&requests).Error
	if err1 != nil {
		fmt.Println(err1)
	}
	params := mux.Vars(r)
	result := make([]models.Request, 0)
	for _, item := range requests {
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

// ----------------------------------------------------------------------------

func DeleteBooking(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, _ := strconv.ParseUint(params["id"], 10, 32)
	db := helpers.DbConnect()
	var request models.Request

	db = db.Debug().Model(&request).Where("id = ?", id).Take(&request).Delete(&request)
	if db.Error != nil {
		fmt.Println(db.Error)
	}
	//
	json.NewEncoder(w).Encode(db.RowsAffected)

}

// ----------------------------------------------------------------------------
// 					   delete driver
// ----------------------------------------------------------------------------

func DeleteDriver(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, _ := strconv.ParseUint(params["id"], 10, 32)
	db := helpers.DbConnect()
	var request []models.Request

	error1 := db.Model(&request).Where("id = ?", id).Update("driver_id", nil).Error
	if error1 != nil {
		json.NewEncoder(w).Encode(error1)
	}
	error2 := db.Model(&request).Where("id = ?", id).Update("state", "Not dispatched").Error
	if error2 != nil {
		json.NewEncoder(w).Encode(error2)
	}
	json.NewEncoder(w).Encode("element deleted")

}

// ----------------------------------------------------------------------------
// 					  affected driver
// ----------------------------------------------------------------------------

func AffectedDriver(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, _ := strconv.ParseUint(params["id"], 10, 32)
	driver_id, _ := strconv.ParseUint(params["driver_id"], 10, 32)

	db := helpers.DbConnect()
	var request []models.Request

	error1 := db.Model(&request).Where("id = ?", id).Update("driver_id", driver_id).Error
	if error1 != nil {
		json.NewEncoder(w).Encode(error1)
	}
	error2 := db.Model(&request).Where("id = ?", id).Update("state", "affected").Error
	if error2 != nil {
		json.NewEncoder(w).Encode(error2)
	}
	json.NewEncoder(w).Encode("element deleted")

}
