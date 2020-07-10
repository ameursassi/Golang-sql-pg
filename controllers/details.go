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

const querydetails = `SELECT
request.id                                                           "id" ,                                                          
request.saas_office_id                                                           "saas_office_id" ,                                                          
request.address_pick_up                                       "address_pick_up" ,                                                          
request.address_drop_off                                                 "address_drop_off",
request.ride_date                                                "ride_date",
request.state                                                    "state",
request.driver_price                                                "driver_price",
request.options_price                                                "options_price",
request.ride_price                                                "ride_price",
request.flight_number                                                "flight_number",
request.comment                                                  "comment",
request.payment_type                                                  "payment_type",
vehicle_type.designation                                                        "designation",
request_static_information.request_options                             "request_options",
request_static_information.intermediate_steps                             "intermediate_steps",
CASE WHEN real_passenger.id IS NULL
THEN passenger.first_name || ' ' || passenger.last_name ELSE
real_passenger.first_name || ' ' || real_passenger.last_name END  "passenger_full_name",
CASE WHEN real_passenger.id IS NULL
THEN passenger.phone_number ELSE
real_passenger.phone_number END                                   "passenger_phone_number",
request.driver_id                                                            "driver_id",
CASE WHEN request.driver_id IS NULL THEN '-'
ELSE CONCAT(back_end_user.first_name || ' ' || back_end_user.last_name)  END   "driver_name",
back_end_user.phone_number                                                         "driver_phone",
driver.driver_status                                   "driver_status" ,                           
CONCAT(driver_car.car_marque || ' ' ||driver_car.car_model)                    "car",
saas_office_to_driver_relation.last_trace_date                                   "last_trace_date"  ,                          
colis.description                                                       "package_types",
sender.phone_number                                                     "sender_phone_number",
CONCAT(recipient.first_name || ' ' || recipient.last_name)              "recipient",
recipient.phone_number                                                  "recipient_phone_number",
saas_companie.name                                                       "company",
request.type                                                                 "type"
FROM
requests request
LEFT OUTER JOIN customers customer on customer.id = request.customer_id
LEFT OUTER JOIN passengers recipient on recipient.id = request.delivery_recipient_id
LEFT OUTER JOIN passengers sender    on sender.id = request.delivery_sender_id
LEFT OUTER JOIN passengers passenger on passenger.id = request.passenger_id	
LEFT OUTER JOIN vehicle_types vehicle_type on vehicle_type.id=request.driver_id 
LEFT OUTER JOIN request_static_informations request_static_information on request_static_information.request_id=request.id  
LEFT OUTER JOIN requests outsourced_origin_request on outsourced_origin_request.id = request.outsourced_origin_request_id
LEFT OUTER JOIN passengers real_passenger on real_passenger.id = outsourced_origin_request.passenger_id
LEFT OUTER JOIN requests driver_id on driver_id.id = request.driver_id 
LEFT OUTER JOIN back_end_users back_end_user on back_end_user.driver_id = request.driver_id 
LEFT OUTER JOIN drivers driver on driver.id = request.driver_id 
LEFT OUTER JOIN driver_cars driver_car on driver_car.driver_id=request.driver_id 
LEFT OUTER JOIN saas_office_to_driver_relations saas_office_to_driver_relation on saas_office_to_driver_relation.driver_id=request.driver_id  
LEFT OUTER JOIN saas_companies saas_companie on saas_companie.id=request.saas_company_id  
LEFT JOIN LATERAL (
	SELECT
	array_to_string(array_agg(package_types.name || ': ' || request_to_package_types.number_of_package), ', ') "description"
	FROM request_to_package_types
	INNER JOIN package_types ON package_types.id = request_to_package_types.package_type_id
	WHERE request.id = request_to_package_types.request_id
	AND   request_to_package_types.number_of_package <> 0
) colis ON TRUE
WHERE
 request.ride_date >= date_trunc('month', now()) - interval '3 month' AND
 request.type!='null'
   ORDER BY request.estimate_pick_up_date ASC,  request.id ASC`

// @GetDetails godoc
// @Summary Get details for a given Id
// @Description Get details of ride corresponding to the input Id
// @Tags details
// @Accept  json
// @Produce  json
// @Param id path int true "ID of the rides"
// @Success 200 {object} models.RequestDetails
// @Router /details/{id} [get]
func GetDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := helpers.DbConnect()
	var requestdetails []models.RequestDetails
	err1 := db.Raw(querydetails).Scan(&requestdetails).Error
	if err1 != nil {
		fmt.Println(err1)
	}
	params := mux.Vars(r)
	for _, item := range requestdetails {
		id, err := strconv.ParseUint(params["id"], 10, 32)
		if err != nil {
			fmt.Println(err)
		}
		if uint(id) == item.ID {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

}
