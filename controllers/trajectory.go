package controllers

import (
	"encoding/json"
	"fmt"
	helpers "github/map_dashboard/Helpers"
	models "github/map_dashboard/Models"
	"io/ioutil"
	"strconv"

	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const querytrajectory = ` SELECT
request.id                                            "id" ,
request.address_pick_up_lat                           "address_pick_up_lat",
request.address_pick_up_long                          "address_pick_up_long",
request.address_drop_off_lat                          "address_drop_off_lat",
request.address_drop_off_long                         "address_drop_off_long"
FROM
requests request
WHERE
request.ride_date >= date_trunc('month', now()) - interval '3 month'
ORDER BY request.ride_date ASC,  request.id ASC`

type PathTrajectory struct {
	LatOrigin       float64
	LongOrigin      float64
	LatDestination  float64
	LongDestination float64
	Path            models.ReturnPath
}

// @GetTrajectory godoc
// @Summary Get details for a given trajectoryId
// @Description Get details of trajectory corresponding to the input trajectoryId
// @Tags trajectory
// @Accept  json
// @Produce  json
// @Param id path int true "ID of the trajectory"
// @Success 200 {object} models.Trajectory
// @Router /trajectory/{id} [get]
func GetTrajectory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := helpers.DbConnect()
	var requests []models.Trajectory
	err1 := db.Raw(querytrajectory).Scan(&requests).Error
	if err1 != nil {
		fmt.Println(err1)
	}
	params := mux.Vars(r)
	result := make(map[string]float64)
	for _, item := range requests {
		id, err := strconv.ParseUint(params["id"], 10, 32)
		if err != nil {
			fmt.Println(err)
		}
		if uint(id) == item.ID {
			fmt.Println(item.ID)
			// fmt.Println(uint(id))
			result["address_pick_up_lat"] = item.AddressPickUpLat
			result["address_pick_up_long"] = item.AddressPickUpLong
			result["address_drop_off_lat"] = item.AddressDropOffLat
			result["address_drop_off_long"] = item.AddressDropOffLong
		}
	}
	lat_pick := result["address_pick_up_lat"]
	long_pick := result["address_pick_up_long"]
	lat_drop := result["address_drop_off_lat"]
	long_drop := result["address_drop_off_long"]

	url := fmt.Sprintf("https://maps-manager.herokuapp.com/api/directions?provider=google&origin=%f,%f&destination=%f,%f&traffic=true&WithPath=true", lat_pick, long_pick, lat_drop, long_drop)
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE1OTQ4MDkyMTUsInVzZXIiOiJ5dXNvIn0.F4RaKLfGcdmiL8pvBZygJ6XaVcUCO_UL3E1RxA3ZLWY"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		// handle err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Token", token)
	resp, err := http.DefaultClient.Do(req) // Call the API
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()
	var pathreturn models.ReturnPath
	if resp.StatusCode == http.StatusOK {
		path, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("error to get result from api")
		}
		//Print the Result that is returned in the response of the API
		errpath := json.Unmarshal(path, &pathreturn)
		if errpath != nil {
			fmt.Println("error path", errpath)
		}
	} else {
		fmt.Println("StatusCode is not ok")
	}
	polyline := PathTrajectory{lat_pick, long_pick, lat_drop, long_drop, pathreturn}
	json.NewEncoder(w).Encode(polyline)

}
