// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/transaction": {
            "get": {
                "description": "Return transaction information by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Transaction"
                ],
                "summary": "Get transaction by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Transaction ID",
                        "name": "tid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Transaction"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Saves the money transaction for borrowing book",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Transaction"
                ],
                "summary": "Register transaction",
                "parameters": [
                    {
                        "description": "Details of borrowing contains bookID, userID, and the price",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.TransactionCreateReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.TransactionCreateResp"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Transaction": {
            "type": "object",
            "properties": {
                "bookID": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "price": {
                    "type": "number"
                },
                "userID": {
                    "type": "integer"
                }
            }
        },
        "model.TransactionCreateReq": {
            "type": "object",
            "properties": {
                "bid": {
                    "type": "integer"
                },
                "price": {
                    "type": "number"
                },
                "uid": {
                    "type": "integer"
                }
            }
        },
        "model.TransactionCreateResp": {
            "type": "object",
            "properties": {
                "tid": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Transaction API",
	Description:      "This is a microservice for book borrowing transactions",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
