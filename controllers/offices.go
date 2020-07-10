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

const queryoffices = `SELECT
request.id,                           "id" , 
request.saas_company_id               "saas_company_id",                                                
request.name                          "name",                          
request.allow_delivery                "allow_delivery",
request.office_center_lat             "office_center_lat",
request.office_center_long            "office_center_long",
request.office_radius                 "office_radius",
request.distance_unit                 "distance_unit"
FROM
saas_offices request
ORDER BY request.id ASC`

// @GetOffices list godoc
// @Summary Get offices for a given offices list
// @Description Get details of list offices
// @Tags offices
// @Accept  json
// @Produce  json
// @Success 200 {object} models.SaasOffices
// @Router /offices [get]
func GetOffices(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := helpers.DbConnect()
	var saas_offices []models.SaasOffices
	var err error
	if err != nil {
		fmt.Printf("The request failed with error %s\n", err)
	}
	err = db.Raw(queryoffices).Scan(&saas_offices).Error
	for _, element := range saas_offices {
		fmt.Println(element)
	}
	json.NewEncoder(w).Encode(saas_offices)

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
func GetOffice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := helpers.DbConnect()
	var saas_offices []models.SaasOffices
	err1 := db.Raw(queryoffices).Scan(&saas_offices).Error
	if err1 != nil {
		fmt.Println(err1)
	}
	params := mux.Vars(r)
	for _, item := range saas_offices {
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




