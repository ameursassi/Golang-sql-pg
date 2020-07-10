package controllers

import (
	"encoding/json"
	"fmt"
	helpers "github/map_dashboard/Helpers"
	models "github/map_dashboard/Models"

	// models "github/map_dashboard/Models"
	"strconv"

	"github.com/gorilla/mux"

	"net/http"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Request struct {
	// gorm.Model
	ID                       uint   `json:"id"`
	SaasOfficeID             uint   `json:"saas_office_id"`
	SaasCompanyID            uint   `json:"saas_company_id"`
	ReservationCode          string `json:"reservation_code"`
	PartnerID                string `json:"partner_id"`
	State                    string `json:"state"`
	DriverName               string `json:"driver_name"`
	DriverID                 string `json:"driver_id"`
	PackageTypes             string `json:"package_types"`
	EstimatePickUpDate       string `json:"estimate_pick_up_date"`
	EstimateDropOffDate      string `json:"estimate_drop_off_date"`
	AddressDropOff           string `json:"address_drop_off"`
	DropOffTimeWindowEnd     string `json:"drop_off_time_window_end"`
	PickUpTimeMargin         string `json:"pick_up_time_margin"`
	DropOffTimeMargin        string `json:"drop_off_time_margin"`
	AddressPickUp            string `json:"address_pick_up"`
	Recipient                string `json:"recipient"`
	RecipientPhoneNumber     string `json:"recipient_phone_number"`
	SenderPhoneNumber        string `json:"sender_phone_number"`
	Type                     string `json:"type"`
	AddressPickUpPostalCode  string `json:"address_pick_up_postal_code"`
	AddressDropOffPostalCode string `json:"address_drop_off_postal_code"`
}

const querydelivery = `SELECT
request.id                                                    "id" ,                                                          
request.saas_office_id                                       "saas_office_id" ,                                                          
request.driver_id                                             "driver_id",
request.saas_company_id                                                 "saas_company_id",
request.reservation_code                                                "reservation_code",
request.address_pick_up                                                 "address_pick_up",
request.address_drop_off                                                 "address_drop_off",
request.address_pick_up_postal_code,                                    
request.address_drop_off_postal_code,                                   
to_char(request.estimate_pick_up_date AT TIME ZONE 'UTC')   "estimate_pick_up_date",
	to_char(request.estimate_drop_off_date AT TIME ZONE 'UTC')  "estimate_drop_off_date",
to_char(pick_up_time_window.begin_date  AT TIME ZONE 'UTC') "pick_up_time_window_begin",
to_char(pick_up_time_window.end_date    AT TIME ZONE 'UTC') "pick_up_time_window_end",
to_char(drop_off_time_window.begin_date AT TIME ZONE 'UTC') "drop_off_time_window_begin",
to_char(drop_off_time_window.end_date   AT TIME ZONE 'UTC') "drop_off_time_window_end",
(CASE WHEN request.state IN ('packages_on_board', 'pick_up_done', 'way_to_drop_off', 'wait_drop_off', 'finish')
	then NULL
	else (EXTRACT(EPOCH FROM (pick_up_time_window.end_date - request.estimate_pick_up_date)::INTERVAL)::int/60 )
END) "pick_up_time_margin",
(EXTRACT(EPOCH FROM (drop_off_time_window.end_date - request.estimate_drop_off_date)::INTERVAL)::int/60 ) "drop_off_time_margin",
request.state                                                           "state",
CASE WHEN request.partner_id ISNULL then ''     END                         "partner_id",
colis.description                                                       "package_types",
sender.phone_number                                                     "sender_phone_number",
CONCAT(recipient.first_name || ' ' || recipient.last_name)              "recipient",
recipient.phone_number                                                  "recipient_phone_number",

CASE WHEN driver.id IS NULL THEN '-'
	ELSE CONCAT(back_end_user.first_name || ' ' || back_end_user.last_name)  END   "driver_name",

request.type                                                             "type"
FROM
	requests request
LEFT OUTER JOIN customers customer on customer.id = request.customer_id
LEFT OUTER JOIN passengers recipient on recipient.id = request.delivery_recipient_id
LEFT OUTER JOIN passengers sender    on sender.id = request.delivery_sender_id
LEFT OUTER JOIN front_end_users front_end_user on front_end_user.id = customer.front_end_user_id
LEFT OUTER JOIN companies company ON company.id = request.company_id
LEFT JOIN LATERAL (
	SELECT
	array_to_string(array_agg(package_types.name || ': ' || request_to_package_types.number_of_package), ', ') "description"
	FROM request_to_package_types
	INNER JOIN package_types ON package_types.id = request_to_package_types.package_type_id
	WHERE request.id = request_to_package_types.request_id
	AND   request_to_package_types.number_of_package <> 0
) colis ON TRUE
LEFT OUTER JOIN jobs pick_up_job ON (request.id = pick_up_job.request_id AND pick_up_job.type = 'PickUpJob')
LEFT OUTER JOIN jobs drop_off_job ON (request.id = drop_off_job.request_id AND drop_off_job.type = 'DropOffJob')
LEFT OUTER JOIN time_windows pick_up_time_window ON pick_up_time_window.id = pick_up_job.time_window_id
LEFT OUTER JOIN time_windows drop_off_time_window ON drop_off_time_window.id = drop_off_job.time_window_id
LEFT OUTER JOIN time_window_settings pick_up_time_window_setting ON pick_up_time_window.time_window_setting_id = pick_up_time_window_setting.id
LEFT OUTER JOIN time_window_settings drop_off_time_window_setting ON drop_off_time_window.time_window_setting_id = drop_off_time_window_setting.id
LEFT OUTER JOIN drivers driver on driver.id = request.driver_id
LEFT OUTER JOIN driver_companies driver_company on driver.driver_company_id = driver_company.id
LEFT OUTER JOIN back_end_users back_end_user on back_end_user.driver_id = driver.id
WHERE
request.type = 'DeliveryRequest' AND

request.ride_date>= date_trunc('month', now()) - interval '2 month'
ORDER BY request.ride_date ASC`

// @GetDeliveries godoc
// @Summary Get details for a given deliveriesId
// @Description Get details of deliveries corresponding to the input deliveriesId
// @Tags deliveries
// @Accept  json
// @Produce  json
// @Param saas_office_id path int true "ID of the deliveries"
// @Success 200 {object} models.RequestD
// @Router /deliveries/{saas_office_id} [get]
func GetDeliveries(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := helpers.DbConnect()
	var requests []models.Request

	err1 := db.Raw(querydelivery).Scan(&requests).Error
	if err1 != nil {
		fmt.Println(err1)
	}
	params := mux.Vars(r)
	fmt.Println(params)
	results := make([]models.Request, 0)
	for _, item := range requests {
		saas_office_id, err := strconv.ParseUint(params["saas_office_id"], 10, 32)
		fmt.Println(saas_office_id)
		if err != nil {
			fmt.Println(err)
		}

		if uint(saas_office_id) == item.SaasOfficeID {
			// json.NewEncoder(w).Encode(item)
			results = append(results, item)
		}
	}
	json.NewEncoder(w).Encode(results)
}

// ----------------------------------------------------------------------------

func DeleteDeliveries(w http.ResponseWriter, r *http.Request) {
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

func DeleteDriverBod(w http.ResponseWriter, r *http.Request) {
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
