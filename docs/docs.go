// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2020-06-14 23:00:36.8482597 +0200 CEST m=+0.132007501

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "email": "soberkoder@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/bookings/{saas_office_id}": {
            "get": {
                "description": "Get details of passenger corresponding to the input passengerId",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "bookings"
                ],
                "summary": "Get details for a given passengerId",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the passengers",
                        "name": "saas_office_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.RequestP"
                        }
                    }
                }
            }
        },
        "/deliveries/{saas_office_id}": {
            "get": {
                "description": "Get details of deliveries corresponding to the input deliveriesId",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "deliveries"
                ],
                "summary": "Get details for a given deliveriesId",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the deliveries",
                        "name": "saas_office_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.RequestD"
                        }
                    }
                }
            }
        },
        "/drivers/{saas_office_id}": {
            "get": {
                "description": "Get details of driver corresponding to the input driverId",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "drivers"
                ],
                "summary": "Get details for a given driverId",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the drivers",
                        "name": "saas_office_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Drivers"
                        }
                    }
                }
            }
        },
        "/offices": {
            "get": {
                "description": "Get details of list offices",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "offices"
                ],
                "summary": "Get offices for a given offices list",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.SaasOffices"
                        }
                    }
                }
            }
        },
        "/offices/{saas_office_id}": {
            "get": {
                "description": "Get details of office corresponding to the input officeId",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "offices"
                ],
                "summary": "Get details for a given officeId",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the offices",
                        "name": "saas_office_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.SaasOffices"
                        }
                    }
                }
            }
        },
        "/trajectory/{id}": {
            "get": {
                "description": "Get details of trajectory corresponding to the input trajectoryId",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "trajectory"
                ],
                "summary": "Get details for a given trajectoryId",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the trajectory",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Trajectory"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Drivers": {
            "type": "object",
            "properties": {
                "action_type": {
                    "type": "integer"
                },
                "car_type": {
                    "type": "string"
                },
                "commission_percentage": {
                    "type": "number"
                },
                "designation": {
                    "type": "string"
                },
                "driver_id": {
                    "type": "integer"
                },
                "driver_name": {
                    "type": "string"
                },
                "driver_phone": {
                    "type": "string"
                },
                "driver_status": {
                    "type": "string"
                },
                "last_trace": {
                    "type": "string"
                },
                "last_trace_date": {
                    "type": "string"
                },
                "last_trace_lat": {
                    "type": "number"
                },
                "last_trace_long": {
                    "type": "number"
                },
                "marker_map": {
                    "type": "string"
                },
                "note": {
                    "type": "integer"
                },
                "office_radius": {
                    "type": "integer"
                },
                "priority": {
                    "type": "integer"
                },
                "saas_office_id": {
                    "type": "integer"
                }
            }
        },
        "models.RequestD": {
            "type": "object",
            "properties": {
                "address_drop_off": {
                    "type": "string"
                },
                "address_drop_off_postal_code": {
                    "type": "string"
                },
                "address_pick_up": {
                    "type": "string"
                },
                "address_pick_up_postal_code": {
                    "type": "string"
                },
                "driver_name": {
                    "type": "string"
                },
                "drop_off_time_margin": {
                    "type": "string"
                },
                "drop_off_time_window_end": {
                    "type": "string"
                },
                "estimate_drop_off_date": {
                    "type": "string"
                },
                "estimate_pick_up_date": {
                    "type": "string"
                },
                "id": {
                    "description": "gorm.Model",
                    "type": "integer"
                },
                "package_types": {
                    "type": "string"
                },
                "partner_id": {
                    "type": "string"
                },
                "pick_up_time_margin": {
                    "type": "string"
                },
                "recipient": {
                    "type": "string"
                },
                "recipient_phone_number": {
                    "type": "string"
                },
                "reservation_code": {
                    "type": "string"
                },
                "saas_company_id": {
                    "type": "integer"
                },
                "saas_office_id": {
                    "type": "integer"
                },
                "sender_phone_number": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "models.RequestP": {
            "type": "object",
            "properties": {
                "address_drop_off": {
                    "type": "string"
                },
                "address_drop_off_lat": {
                    "type": "number"
                },
                "address_drop_off_long": {
                    "type": "number"
                },
                "address_drop_off_postal_code": {
                    "type": "string"
                },
                "address_pick_up": {
                    "type": "string"
                },
                "address_pick_up_lat": {
                    "type": "number"
                },
                "address_pick_up_long": {
                    "type": "number"
                },
                "address_pick_up_postal_code": {
                    "type": "string"
                },
                "company": {
                    "type": "string"
                },
                "driver_id": {
                    "type": "string"
                },
                "driver_name": {
                    "type": "string"
                },
                "driver_phone": {
                    "type": "string"
                },
                "estimate_distance_m": {
                    "type": "integer"
                },
                "estimate_drop_off_date": {
                    "type": "string"
                },
                "estimate_pick_up_date": {
                    "type": "string"
                },
                "fire_time": {
                    "type": "string"
                },
                "flight_number": {
                    "type": "string"
                },
                "id": {
                    "description": "gorm.Model\ntableName                struct{} ` + "`" + `pg:\"requests\"` + "`" + `",
                    "type": "integer"
                },
                "passenger_full_name": {
                    "type": "string"
                },
                "passenger_phone_number": {
                    "type": "string"
                },
                "reservation_code": {
                    "type": "string"
                },
                "ride_date": {
                    "type": "string"
                },
                "saas_company_id": {
                    "type": "integer"
                },
                "saas_office_id": {
                    "type": "integer"
                },
                "state": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "models.SaasOffices": {
            "type": "object",
            "properties": {
                "allow_delivery": {
                    "type": "boolean"
                },
                "distance_unit": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "office_center_lat": {
                    "type": "number"
                },
                "office_center_long": {
                    "type": "number"
                },
                "office_radius": {
                    "type": "integer"
                },
                "saas_company_id": {
                    "type": "integer"
                }
            }
        },
        "models.Trajectory": {
            "type": "object",
            "properties": {
                "address_drop_off_lat": {
                    "type": "number"
                },
                "address_drop_off_long": {
                    "type": "number"
                },
                "address_pick_up_lat": {
                    "type": "number"
                },
                "address_pick_up_long": {
                    "type": "number"
                },
                "id": {
                    "type": "integer"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "localhost/8080",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "bookings API",
	Description: "This is a sample service for managing orders",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
