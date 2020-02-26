// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

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
        "contact": {},
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth/device": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "로그인한 유저 고객사의 모델리스트  Authorization Bearer token 을 넣으세요",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "device"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Device"
                            }
                        }
                    }
                }
            }
        },
        "/auth/device/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "모델 단일조회  Authorization Bearer token 을 넣으세요",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "device"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "장치아이디",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Device"
                        }
                    }
                }
            }
        },
        "/auth/model": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "로그인한 유저 고객사의 모델리스트  Authorization Bearer token 을 넣으세요",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "model"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Model"
                            }
                        }
                    }
                }
            }
        },
        "/auth/model/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "모델 단일조회  Authorization Bearer token 을 넣으세요",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "model"
                ],
                "operationId": "get-string-by-int",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "모델아이디",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Model"
                        }
                    }
                }
            }
        },
        "/auth/site": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "로그인한 유저 고객사의 사이트리스트  Authorization Bearer token 을 넣으세요",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "site"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Site"
                            }
                        }
                    }
                }
            }
        },
        "/auth/site/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "site"
                ],
                "operationId": "get-string-by-int",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "사이트아이디",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Site"
                        }
                    }
                }
            }
        },
        "/command": {
            "post": {
                "description": "redis publish",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "redis"
                ],
                "parameters": [
                    {
                        "description": "command",
                        "name": "command",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.Command"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.CommandResponse"
                        }
                    }
                }
            }
        },
        "/customs": {
            "get": {
                "description": "전체고객 목록조회",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Custom"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Custom"
                            }
                        }
                    }
                }
            }
        },
        "/customs/{id}": {
            "get": {
                "description": "고객 id로 고객 단일조회",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Custom"
                ],
                "summary": "{id}고객아이디",
                "operationId": "get-string-by-int",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "고객아이디",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Custom"
                        }
                    }
                }
            }
        },
        "/equip": {
            "get": {
                "description": "iotdata5 조회 limit offset 을 넣지않으면 최근 10건",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "equip"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "장치아이디",
                        "name": "id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "측정키",
                        "name": "key",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "offset",
                        "name": "offset",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.EquipResponse"
                        }
                    }
                }
            }
        },
        "/event": {
            "get": {
                "description": "장치의 최근이벤트 히스토리 조회",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "equip"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "장치아이디",
                        "name": "id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "limit",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseEventData"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "로그인후 jwt토큰발행",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "login"
                ],
                "parameters": [
                    {
                        "description": "아이디 및 패스워드",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/mid.login"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/mid.loginSuccess"
                        }
                    }
                }
            }
        },
        "/site": {
            "get": {
                "description": "해당사이트에 속한 장치들의 최근측정데이터",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "equip"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "사이트아이디",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.SiteLastDataResponse"
                        }
                    }
                }
            }
        },
        "/users": {
            "get": {
                "description": "전체사용자 목록조회",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Users"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.FailedMessage"
                        }
                    }
                }
            }
        },
        "/users/{id}": {
            "get": {
                "description": "사용자 id로 사용자 단일조회",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "{id} 사용자아이디",
                "operationId": "get-string-by-int",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "사용자아이디",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Users"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.FailedMessage"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.Command": {
            "type": "object",
            "required": [
                "actionkey",
                "actionval",
                "dev-id"
            ],
            "properties": {
                "actionkey": {
                    "type": "string",
                    "example": "acky"
                },
                "actionval": {
                    "type": "integer",
                    "example": 200
                },
                "dev-id": {
                    "type": "string",
                    "example": "Wxldi39DJecl2dUdlJWL34"
                }
            }
        },
        "controller.CommandResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "success"
                },
                "subscribeCount": {
                    "type": "integer",
                    "example": 5
                }
            }
        },
        "controller.EquipResponse": {
            "type": "object",
            "properties": {
                "i": {
                    "type": "string",
                    "example": "KyvrjRACQJGd3Q7Q9udzo4"
                },
                "k": {
                    "type": "string",
                    "example": "co2"
                },
                "value": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/controller.TimeValue"
                    }
                }
            }
        },
        "controller.EquipSeries": {
            "type": "object",
            "properties": {
                "i": {
                    "type": "string",
                    "example": "KyvrjRACQJGd3Q7Q9udzo4"
                },
                "k": {
                    "type": "string",
                    "example": "co2"
                },
                "m": {
                    "type": "integer",
                    "example": 302
                },
                "t": {
                    "type": "string",
                    "example": "2019-02-02"
                }
            }
        },
        "controller.EventData": {
            "type": "object",
            "properties": {
                "actkey": {
                    "type": "string",
                    "example": "acky"
                },
                "actval": {
                    "type": "string",
                    "example": "20"
                },
                "curval": {
                    "type": "integer",
                    "example": 812
                },
                "etype": {
                    "type": "string",
                    "example": "err"
                },
                "eventkey": {
                    "type": "string",
                    "example": "co2"
                },
                "limit": {
                    "type": "string",
                    "example": "50"
                },
                "model": {
                    "type": "string",
                    "example": "14"
                },
                "notikey": {
                    "type": "string",
                    "example": "1"
                },
                "site": {
                    "type": "string",
                    "example": "site004"
                },
                "t": {
                    "type": "string",
                    "example": "2019-02-02"
                }
            }
        },
        "controller.ResponseEventData": {
            "type": "object",
            "properties": {
                "eventList": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/controller.EventData"
                    }
                },
                "i": {
                    "type": "string",
                    "example": "KyvrjRACQJGd3Q7Q9udzo4"
                }
            }
        },
        "controller.SiteLastDataResponse": {
            "type": "object",
            "properties": {
                "SID": {
                    "type": "string",
                    "example": "site004"
                },
                "lastData": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/controller.EquipSeries"
                    }
                }
            }
        },
        "controller.TimeValue": {
            "type": "object",
            "properties": {
                "m": {
                    "type": "integer",
                    "example": 302
                },
                "t": {
                    "type": "string",
                    "example": "2019-02-02"
                }
            }
        },
        "mid.login": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "mid.loginSuccess": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string",
                    "example": "200"
                },
                "expire": {
                    "type": "string",
                    "example": "2020-02-28T09:15:29+09:00"
                },
                "token": {
                    "type": "string",
                    "example": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJDSUQiOjIsIklzQWRtaW4iOnRydWUsIlVJRCI6NDgsIlVzZXJOYW1lIjoi6rmA7ZiV6re8IiwiZXhwIjoxNTgyODQ4OTI5LCJvcmlnX2lhdCI6MTU4MjI0NDEyOX0.ErN3ajVe7Sj5s6jaJQFGfMMaeP8jietm9uP0feacfxA"
                }
            }
        },
        "models.Custom": {
            "type": "object",
            "properties": {
                "appryn": {
                    "type": "string",
                    "example": "Y"
                },
                "basadr": {
                    "type": "string",
                    "example": "서울특별시 강서구 화곡로 12"
                },
                "compno": {
                    "type": "string",
                    "example": "17282718127"
                },
                "createdat": {
                    "type": "string",
                    "example": "2020-01-29 13:10:39"
                },
                "dtladr": {
                    "type": "string",
                    "example": "심당빌딩 3층 302호"
                },
                "eid": {
                    "type": "string",
                    "example": "SAMSUNG"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "name": {
                    "type": "string",
                    "example": "삼성"
                },
                "orgid": {
                    "type": "integer",
                    "example": 2
                },
                "postno": {
                    "type": "string",
                    "example": "10414"
                },
                "rec": {
                    "type": "string",
                    "example": "비고"
                },
                "reptnm": {
                    "type": "string",
                    "example": "고영호"
                },
                "type": {
                    "type": "integer",
                    "example": 1
                },
                "updatedat": {
                    "type": "string",
                    "example": "2020-01-29 13:10:39"
                },
                "useyn": {
                    "type": "string",
                    "example": "Y"
                }
            }
        },
        "models.Device": {
            "type": "object",
            "properties": {
                "authid": {
                    "type": "string",
                    "example": "?"
                },
                "authkey": {
                    "type": "string",
                    "example": "??"
                },
                "createdat": {
                    "type": "string",
                    "example": "2020-01-29 13:10:39"
                },
                "custid": {
                    "type": "integer",
                    "example": 4
                },
                "equpid": {
                    "type": "integer",
                    "example": 1
                },
                "firmware": {
                    "type": "string",
                    "example": "firmware version"
                },
                "id": {
                    "type": "string",
                    "example": "TD2BDhvjKEaZo37c4DLAq6"
                },
                "modelid": {
                    "type": "integer",
                    "example": 4
                },
                "name": {
                    "type": "string",
                    "example": "device-37"
                },
                "prtctype": {
                    "type": "string",
                    "example": "prtctype"
                },
                "rec": {
                    "type": "string",
                    "example": "비고"
                },
                "serialno": {
                    "type": "string",
                    "example": "102312-12312"
                },
                "siteid": {
                    "type": "integer",
                    "example": 1
                },
                "status": {
                    "type": "string",
                    "example": "1"
                },
                "tag": {
                    "type": "string",
                    "example": "tag"
                },
                "type": {
                    "type": "string",
                    "example": "??"
                },
                "updatedat": {
                    "type": "string",
                    "example": "2020-01-29 13:10:39"
                },
                "useyn": {
                    "type": "string",
                    "example": "Y"
                }
            }
        },
        "models.FailedMessage": {
            "type": "object",
            "properties": {
                "err": {
                    "type": "object"
                },
                "message": {
                    "type": "string",
                    "example": "failed"
                }
            }
        },
        "models.Model": {
            "type": "object",
            "properties": {
                "createdat": {
                    "type": "string",
                    "example": "2020-01-29 13:10:39"
                },
                "custid": {
                    "type": "integer",
                    "example": 4
                },
                "firmware": {
                    "type": "string",
                    "example": "firmware version"
                },
                "id": {
                    "type": "integer",
                    "example": 4
                },
                "name": {
                    "type": "string",
                    "example": "viux-38"
                },
                "prtcio": {
                    "type": "string",
                    "example": "1"
                },
                "prtctype": {
                    "type": "string",
                    "example": "prtctype"
                },
                "rec": {
                    "type": "string",
                    "example": "비고"
                },
                "tag": {
                    "type": "string",
                    "example": "tag"
                },
                "type": {
                    "type": "string",
                    "example": "??"
                },
                "updatedat": {
                    "type": "string",
                    "example": "2020-01-29 13:10:39"
                },
                "useyn": {
                    "type": "string",
                    "example": "Y"
                }
            }
        },
        "models.Site": {
            "type": "object",
            "properties": {
                "basadr": {
                    "type": "string",
                    "example": "서울특별시 강서구 화곡로 132"
                },
                "createdat": {
                    "type": "string",
                    "example": "2020-01-29 13:10:39"
                },
                "custid": {
                    "type": "integer",
                    "example": 1
                },
                "dtladr": {
                    "type": "string",
                    "example": "심당빌딩 3층 301호"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "lat": {
                    "type": "string",
                    "example": "32°"
                },
                "lon": {
                    "type": "string",
                    "example": "24°"
                },
                "name": {
                    "type": "string",
                    "example": "제주1공장"
                },
                "postno": {
                    "type": "string",
                    "example": "10232"
                },
                "rec": {
                    "type": "string",
                    "example": "비고"
                },
                "type": {
                    "type": "string",
                    "example": "??"
                },
                "updatedat": {
                    "type": "string",
                    "example": "2020-01-29 13:10:39"
                },
                "useyn": {
                    "type": "string",
                    "example": "Y"
                }
            }
        },
        "models.Users": {
            "type": "object",
            "properties": {
                "admin": {
                    "type": "integer",
                    "example": 1
                },
                "createdat": {
                    "type": "string",
                    "example": "TIMESTAMP"
                },
                "custid": {
                    "type": "integer",
                    "example": 23
                },
                "dept": {
                    "type": "string",
                    "example": "경영지원팀"
                },
                "email": {
                    "type": "string",
                    "example": "example@example.com"
                },
                "gfid": {
                    "type": "integer",
                    "example": 12
                },
                "id": {
                    "type": "integer",
                    "example": 4
                },
                "name": {
                    "type": "string",
                    "example": "강준구"
                },
                "phone": {
                    "type": "string",
                    "example": "01012341234"
                },
                "position": {
                    "type": "string",
                    "example": "사원"
                },
                "pwd": {
                    "type": "string",
                    "example": "abcd1234"
                },
                "rec": {
                    "type": "string",
                    "example": "비고"
                },
                "updatedat": {
                    "type": "string",
                    "example": "SAMSUNG"
                },
                "useyn": {
                    "type": "string",
                    "example": "Y"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
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
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
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
