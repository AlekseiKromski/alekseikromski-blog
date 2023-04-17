// Code generated by swaggo/swag. DO NOT EDIT
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
        "/v1/create-post": {
            "post": {
                "description": "Create a post",
                "summary": "Create post",
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/v1/get-last-posts": {
            "get": {
                "description": "Get last posts from storage",
                "produces": [
                    "application/json"
                ],
                "summary": "List of last posts",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Post"
                            }
                        }
                    },
                    "400": {
                        "description": "if we cannot decode or encode payload",
                        "schema": {
                            "$ref": "#/definitions/v1.JsonError"
                        }
                    },
                    "500": {
                        "description": "if we have bad payload",
                        "schema": {
                            "$ref": "#/definitions/v1.InputError"
                        }
                    }
                }
            }
        },
        "/v1/post/get-last-posts-by-category/{category_id}/{size}/{offset}": {
            "get": {
                "description": "Get last posts from storage filtered by category",
                "produces": [
                    "application/json"
                ],
                "summary": "List of last posts filtered by category",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Post"
                            }
                        }
                    },
                    "400": {
                        "description": "if we cannot decode or encode payload",
                        "schema": {
                            "$ref": "#/definitions/v1.JsonError"
                        }
                    },
                    "500": {
                        "description": "if we have bad payload",
                        "schema": {
                            "$ref": "#/definitions/v1.InputError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Category": {
            "type": "object",
            "properties": {
                "ID": {
                    "type": "integer"
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "models.Post": {
            "type": "object",
            "properties": {
                "category": {
                    "$ref": "#/definitions/models.Category"
                },
                "category_id": {
                    "type": "integer"
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "v1.InputError": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "v1.JsonError": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:3001",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Swagger Aleksei Kromski blog API",
	Description:      "This is a simple api for aleksei kromski blog",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
