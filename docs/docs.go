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
        "/api/all_users": {
            "get": {
                "description": "Получает список всех пользователей",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Получить всех пользователей",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.UserResponse"
                            }
                        }
                    }
                }
            }
        },
        "/api/sort_users": {
            "get": {
                "description": "Получает список всех пользователей с сортировкой",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Получить всех пользователей с сортировкой",
                "parameters": [
                    {
                        "description": "Данные для сортировки пользователя",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.UserSort"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Имя пользователя",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Фамилия пользователя",
                        "name": "surname",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Отчество пользователя",
                        "name": "patronymic",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Возраст пользователя",
                        "name": "age",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Пол пользователя",
                        "name": "gender",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Страна пользователя",
                        "name": "nation",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.UserResponse"
                            }
                        }
                    }
                }
            }
        },
        "/api/user": {
            "post": {
                "description": "Создает нового пользователя",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Создать пользователя",
                "parameters": [
                    {
                        "description": "Данные для создания пользователя",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.CreateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/domain.UserResponse"
                        }
                    }
                }
            }
        },
        "/api/user/{user_id}": {
            "get": {
                "description": "Получает пользователя по заданному идентификатору",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Получить пользователя по идентификатору",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Идентификатор пользователя",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.UserResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Удаляет пользователя по заданному идентификатору",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Удалить пользователя",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Идентификатор пользователя",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "patch": {
                "description": "Обновляет существующего пользователя",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Обновить пользователя",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Идентификатор пользователя",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Данные для обновления пользователя",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.UpdateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.UserResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.CreateUserRequest": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string",
                    "example": "Dmitriy"
                },
                "patronymic": {
                    "type": "string",
                    "example": "Vasilevich"
                },
                "surname": {
                    "type": "string",
                    "example": "Ushakov"
                }
            }
        },
        "domain.UpdateUserRequest": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer",
                    "example": 28
                },
                "gender": {
                    "type": "string",
                    "example": "male"
                },
                "id": {
                    "type": "integer",
                    "example": 1567
                },
                "name": {
                    "type": "string",
                    "example": "Dmitriy"
                },
                "nation": {
                    "type": "string",
                    "example": "slav"
                },
                "patronymic": {
                    "type": "string",
                    "example": "Vasilevich"
                },
                "surname": {
                    "type": "string",
                    "example": "Ushakov"
                }
            }
        },
        "domain.UserResponse": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer",
                    "example": 28
                },
                "gender": {
                    "type": "string",
                    "example": "male"
                },
                "id": {
                    "type": "integer",
                    "example": 1567
                },
                "name": {
                    "type": "string",
                    "example": "Dmitriy"
                },
                "nation": {
                    "type": "string",
                    "example": "slav"
                },
                "patronymic": {
                    "type": "string",
                    "example": "Vasilevich"
                },
                "surname": {
                    "type": "string",
                    "example": "Ushakov"
                }
            }
        },
        "domain.UserSort": {
            "type": "object",
            "properties": {
                "gender": {
                    "type": "string",
                    "example": "male"
                },
                "name": {
                    "type": "string",
                    "example": "Dmitriy"
                },
                "nation": {
                    "type": "string",
                    "example": "slav"
                },
                "patronymic": {
                    "type": "string",
                    "example": "Vasilevich"
                },
                "surname": {
                    "type": "string",
                    "example": "Ushakov"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "localhost:3005",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "EffectiveMobileTask",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
