// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2018-03-05 14:37:53.672989679 +0900 JST m=+0.029691360

package docs

import (
	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "description": "This is a sample server.",
        "title": "Go Echo SQLBoiler sample project",
        "contact": {
            "name": "API Support",
            "url": "dummy",
            "email": "hoge"
        },
        "license": {
            "name": "ken-aio",
            "url": "dummmy"
        },
        "version": "1.0"
    },
    "host": "localhost:1314",
    "basePath": "/",
    "paths": {
        "/api/v1/users": {
            "get": {
                "description": "create new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "User create API",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Some ID",
                        "name": "some_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "create ok",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.Test"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Test": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        }
    }
}`

type s struct{}

func (s *s) ReadDoc() string {
	return doc
}
func init() {
	swag.Register(swag.Name, &s{})
}
