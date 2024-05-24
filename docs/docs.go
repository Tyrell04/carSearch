// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/car": {
            "post": {
                "description": "Create a new car",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "car"
                ],
                "parameters": [
                    {
                        "description": "Car object",
                        "name": "Car",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http.Car"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/api/car/csv": {
            "post": {
                "description": "Create cars from CSV",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "car"
                ],
                "summary": "Create cars from CSV",
                "parameters": [
                    {
                        "type": "file",
                        "description": "CSV file",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/api/car/{hsn}": {
            "get": {
                "description": "Get car by hsn",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "manufacturer"
                ],
                "summary": "Get car by hsn",
                "parameters": [
                    {
                        "type": "string",
                        "description": "hsn",
                        "name": "hsn",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/api/car/{hsn}/{tsn}": {
            "get": {
                "description": "Get car by hsn and tsn",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "car"
                ],
                "summary": "Get car by hsn and tsn",
                "parameters": [
                    {
                        "type": "string",
                        "description": "hsn",
                        "name": "hsn",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "tsn",
                        "name": "tsn",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        }
    },
    "definitions": {
        "http.Car": {
            "type": "object",
            "properties": {
                "hsn": {
                    "type": "string"
                },
                "manufacturer": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "tsn": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
