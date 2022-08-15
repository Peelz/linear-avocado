// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
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
        "/projects": {
            "get": {
                "description": "list entity.Project",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "projects"
                ],
                "summary": "list projects",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entity.Project"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "get entity.Project by uuid",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "projects"
                ],
                "summary": "create project",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/projects/:id": {
            "get": {
                "description": "get entity.Project by uuid",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "projects"
                ],
                "summary": "get project by uuid",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "uuid of project",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Project"
                        }
                    }
                }
            },
            "put": {
                "description": "get entity.Project by uuid",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "projects"
                ],
                "summary": "update project",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "uuid of project",
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
            "delete": {
                "description": "get entity.Project by uuid",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "projects"
                ],
                "summary": "delete project from uuid",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "uuid of project",
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
            }
        },
        "/projects/:id/scan": {
            "post": {
                "description": "scan project and initial job",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "projects"
                ],
                "summary": "scan project from uuid",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "uuid of project",
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
            }
        }
    },
    "definitions": {
        "entity.Job": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "description": "CreatedAt: timestamp",
                    "type": "string"
                },
                "detail": {
                    "type": "object",
                    "additionalProperties": true
                },
                "finishedAt": {
                    "description": "FinishedAt: timestamp",
                    "type": "string"
                },
                "project": {
                    "$ref": "#/definitions/entity.Project"
                },
                "projectID": {
                    "description": "ProjectID: Project ID use for RDBMS",
                    "type": "integer"
                },
                "projectUUID": {
                    "description": "ProjectUUID:  Global ID use for expose public",
                    "type": "string"
                },
                "startedAt": {
                    "description": "StartedAt: timestamp",
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "uuid": {
                    "description": "UUID: Global ID",
                    "type": "string"
                }
            }
        },
        "entity.Project": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "ID: any type of unique id",
                    "type": "integer"
                },
                "jobs": {
                    "description": "Jobs in",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Job"
                    }
                },
                "name": {
                    "description": "RepositoryName: string",
                    "type": "string"
                },
                "url": {
                    "description": "RepositoryUrl: string",
                    "type": "string"
                },
                "uuid": {
                    "description": "UUID: any type of unique id",
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
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
