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
        "/tasks": {
            "post": {
                "description": "Create a new task for a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "Create a new task",
                "parameters": [
                    {
                        "description": "Task object",
                        "name": "task",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Task"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Task created successfully",
                        "schema": {
                            "$ref": "#/definitions/handlers.Message"
                        }
                    },
                    "400": {
                        "description": "Invalid body data",
                        "schema": {
                            "$ref": "#/definitions/handlers.Message"
                        }
                    },
                    "500": {
                        "description": "Failed to create task",
                        "schema": {
                            "$ref": "#/definitions/handlers.Message"
                        }
                    }
                }
            }
        },
        "/tasks/{id}": {
            "delete": {
                "description": "Delete a task by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "Delete a task",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Task ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Task deleted",
                        "schema": {
                            "$ref": "#/definitions/handlers.Message"
                        }
                    },
                    "400": {
                        "description": "ID should be an integer",
                        "schema": {
                            "$ref": "#/definitions/handlers.Message"
                        }
                    },
                    "500": {
                        "description": "Failed to delete task",
                        "schema": {
                            "$ref": "#/definitions/handlers.Message"
                        }
                    }
                }
            }
        },
        "/tasks/{id}/end": {
            "post": {
                "description": "End a period for a task by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "End a period for a task",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Task ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Period ended",
                        "schema": {
                            "$ref": "#/definitions/handlers.Message"
                        }
                    },
                    "400": {
                        "description": "Failed to end period. Period not started",
                        "schema": {
                            "$ref": "#/definitions/handlers.Message"
                        }
                    },
                    "500": {
                        "description": "Failed to end",
                        "schema": {
                            "$ref": "#/definitions/handlers.Message"
                        }
                    }
                }
            }
        },
        "/tasks/{id}/finish": {
            "post": {
                "description": "Mark a task as finished by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "Finish a task",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Task ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Task ended",
                        "schema": {
                            "$ref": "#/definitions/handlers.Message"
                        }
                    },
                    "400": {
                        "description": "ID should be an integer",
                        "schema": {
                            "$ref": "#/definitions/handlers.Message"
                        }
                    },
                    "500": {
                        "description": "Failed to finish task",
                        "schema": {
                            "$ref": "#/definitions/handlers.Message"
                        }
                    }
                }
            }
        },
        "/tasks/{id}/start": {
            "post": {
                "description": "Start a period for a task by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "Start a period for a task",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Task ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Period started",
                        "schema": {
                            "$ref": "#/definitions/handlers.Message"
                        }
                    },
                    "400": {
                        "description": "Failed to start period. Period not finished",
                        "schema": {
                            "$ref": "#/definitions/handlers.Message"
                        }
                    },
                    "500": {
                        "description": "Failed to start",
                        "schema": {
                            "$ref": "#/definitions/handlers.Message"
                        }
                    }
                }
            }
        },
        "/users": {
            "get": {
                "description": "Retrieve users based on filters",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get users",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Passport Number",
                        "name": "passport_number",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Surname",
                        "name": "surname",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Name",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Patronymic",
                        "name": "patronymic",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Address",
                        "name": "address",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "List of users",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.User"
                            }
                        }
                    },
                    "500": {
                        "description": "Failed to get users",
                        "schema": {
                            "$ref": "#/definitions/handlers.Message"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new user with the provided passport number",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Create a new user",
                "parameters": [
                    {
                        "description": "User",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.createUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User ID",
                        "schema": {
                            "$ref": "#/definitions/handlers.Message"
                        }
                    },
                    "400": {
                        "description": "Invalid body data or invalid passport number",
                        "schema": {
                            "$ref": "#/definitions/handlers.Message"
                        }
                    },
                    "500": {
                        "description": "Failed to create user",
                        "schema": {
                            "$ref": "#/definitions/handlers.Message"
                        }
                    }
                }
            }
        },
        "/users/{id}": {
            "put": {
                "description": "Update a user with the provided data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Update a user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Passport Number",
                        "name": "passport_number",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Surname",
                        "name": "surname",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Name",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Patronymic",
                        "name": "patronymic",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Address",
                        "name": "address",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User updated",
                        "schema": {
                            "$ref": "#/definitions/handlers.Message"
                        }
                    },
                    "400": {
                        "description": "Incorrect ID or invalid input data",
                        "schema": {
                            "$ref": "#/definitions/handlers.Message"
                        }
                    },
                    "500": {
                        "description": "Failed to update user",
                        "schema": {
                            "$ref": "#/definitions/handlers.Message"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a user by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Delete a user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User deleted",
                        "schema": {
                            "$ref": "#/definitions/handlers.Message"
                        }
                    },
                    "400": {
                        "description": "ID should be an integer",
                        "schema": {
                            "$ref": "#/definitions/handlers.Message"
                        }
                    },
                    "500": {
                        "description": "Failed to delete user",
                        "schema": {
                            "$ref": "#/definitions/handlers.Message"
                        }
                    }
                }
            }
        },
        "/users/{id}/tasks": {
            "get": {
                "description": "Get tasks for a user within a specified date range and with optional sorting",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get user tasks",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Start date in RFC3339 format",
                        "name": "start_date",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "End date in RFC3339 format",
                        "name": "end_date",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Sort order, can be 'asc' or 'desc'",
                        "name": "sort",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Tasks found",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Task"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid input data",
                        "schema": {
                            "$ref": "#/definitions/handlers.Message"
                        }
                    },
                    "500": {
                        "description": "Failed to get tasks for user",
                        "schema": {
                            "$ref": "#/definitions/handlers.Message"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handlers.Message": {
            "type": "object",
            "properties": {
                "msg": {
                    "type": "string"
                }
            }
        },
        "handlers.createUser": {
            "type": "object",
            "required": [
                "passportNumber"
            ],
            "properties": {
                "passportNumber": {
                    "type": "string"
                }
            }
        },
        "models.Task": {
            "type": "object",
            "required": [
                "task_name",
                "user_id"
            ],
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "is_finished": {
                    "type": "boolean"
                },
                "task_name": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "passport_number": {
                    "type": "string"
                },
                "patronymic": {
                    "type": "string"
                },
                "surname": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.1",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "time-tracker application",
	Description:      "This is a simple backend for time-tracker application without authorization",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
