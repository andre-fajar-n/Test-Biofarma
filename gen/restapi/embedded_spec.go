// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json",
    "multipart/form-data"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Test Biofarma",
    "version": "1.0.0"
  },
  "paths": {
    "/health": {
      "get": {
        "security": [],
        "description": "Check if the App is Running",
        "tags": [
          "health"
        ],
        "summary": "Health Check",
        "operationId": "health",
        "responses": {
          "200": {
            "description": "Health Check",
            "schema": {
              "$ref": "#/definitions/success"
            }
          },
          "default": {
            "description": "Server Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/v1/home": {
      "post": {
        "security": [],
        "description": "Create new home data",
        "tags": [
          "home"
        ],
        "summary": "Create",
        "operationId": "createHome",
        "parameters": [
          {
            "name": "data",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/createHomeParamsBody"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Success create",
            "schema": {
              "$ref": "#/definitions/successCreate"
            }
          },
          "default": {
            "description": "Server Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/v1/home/{home_id}": {
      "get": {
        "security": [],
        "description": "Find one home data",
        "tags": [
          "home"
        ],
        "summary": "Find One",
        "operationId": "FindOneHome",
        "parameters": [
          {
            "type": "integer",
            "format": "uint64",
            "name": "home_id",
            "in": "path",
            "required": true
          },
          {
            "enum": [
              "true",
              "false"
            ],
            "type": "string",
            "default": "false",
            "description": "flag for identify include deleted data or not",
            "name": "include_deleted_data",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "Success get data",
            "schema": {
              "$ref": "#/definitions/successFindOne"
            }
          },
          "default": {
            "description": "Server Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "put": {
        "security": [],
        "description": "Update home data",
        "tags": [
          "home"
        ],
        "summary": "Update",
        "operationId": "UpdateHome",
        "parameters": [
          {
            "type": "integer",
            "format": "uint64",
            "name": "home_id",
            "in": "path",
            "required": true
          },
          {
            "name": "data",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/updateHomeParamsBody"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Success update",
            "schema": {
              "$ref": "#/definitions/success"
            }
          },
          "default": {
            "description": "Server Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Home": {
      "allOf": [
        {
          "$ref": "#/definitions/modelIdentifier"
        },
        {
          "$ref": "#/definitions/homeData"
        },
        {
          "$ref": "#/definitions/modelTrackTime"
        }
      ]
    },
    "createHomeParamsBody": {
      "type": "object",
      "required": [
        "type",
        "address"
      ],
      "properties": {
        "address": {
          "type": "string"
        },
        "type": {
          "type": "number",
          "format": "int",
          "enum": [
            36,
            45,
            72
          ]
        }
      },
      "x-go-gen-location": "operations"
    },
    "error": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer"
        },
        "message": {
          "type": "string",
          "example": "error"
        }
      }
    },
    "homeData": {
      "type": "object",
      "properties": {
        "address": {
          "type": "string"
        },
        "address_line": {
          "type": "string"
        },
        "country": {
          "type": "string"
        },
        "formatted_address": {
          "type": "string"
        },
        "latitude": {
          "type": "number"
        },
        "longitude": {
          "type": "number"
        },
        "postal_code": {
          "type": "string"
        },
        "regency": {
          "type": "string"
        },
        "subdistrict": {
          "type": "string"
        },
        "type": {
          "type": "string"
        }
      }
    },
    "modelIdentifier": {
      "type": "object",
      "required": [
        "id"
      ],
      "properties": {
        "id": {
          "type": "integer",
          "format": "uint64"
        }
      }
    },
    "modelTrackTime": {
      "type": "object",
      "properties": {
        "created_at": {
          "type": "string",
          "format": "date-time",
          "x-go-custom-tag": "gorm:\"column:created_at\"",
          "x-nullable": true,
          "x-omitempty": false
        },
        "deleted_at": {
          "type": "string",
          "format": "date-time",
          "x-go-custom-tag": "gorm:\"column:deleted_at\"",
          "x-nullable": true,
          "x-omitempty": false
        },
        "updated_at": {
          "type": "string",
          "format": "date-time",
          "x-go-custom-tag": "gorm:\"column:updated_at\"",
          "x-nullable": true,
          "x-omitempty": false
        }
      }
    },
    "success": {
      "type": "object",
      "properties": {
        "message": {
          "type": "string"
        }
      }
    },
    "successCreate": {
      "allOf": [
        {
          "$ref": "#/definitions/successCreateAllOf0"
        },
        {
          "$ref": "#/definitions/successCreateAllOf1"
        }
      ]
    },
    "successCreateAllOf0": {
      "type": "object",
      "properties": {
        "message": {
          "type": "string"
        }
      },
      "x-go-gen-location": "models"
    },
    "successCreateAllOf1": {
      "type": "object",
      "properties": {
        "home_id": {
          "type": "number",
          "format": "uint64"
        }
      },
      "x-go-gen-location": "models"
    },
    "successFindOne": {
      "allOf": [
        {
          "$ref": "#/definitions/successCreateAllOf0"
        },
        {
          "$ref": "#/definitions/successFindOneAllOf1"
        }
      ]
    },
    "successFindOneAllOf1": {
      "type": "object",
      "properties": {
        "data": {
          "$ref": "#/definitions/successFindOneAllOf1Data"
        }
      },
      "x-go-gen-location": "models"
    },
    "successFindOneAllOf1Data": {
      "allOf": [
        {
          "$ref": "#/definitions/modelIdentifier"
        },
        {
          "$ref": "#/definitions/homeData"
        },
        {
          "$ref": "#/definitions/modelTrackTime"
        }
      ],
      "x-go-gen-location": "models"
    },
    "updateHomeParamsBody": {
      "type": "object",
      "required": [
        "type",
        "address"
      ],
      "properties": {
        "address": {
          "type": "string"
        },
        "type": {
          "type": "number",
          "format": "int",
          "enum": [
            36,
            45,
            72
          ]
        }
      },
      "x-go-gen-location": "operations"
    }
  },
  "securityDefinitions": {
    "authorization": {
      "type": "apiKey",
      "name": "Authorization",
      "in": "header"
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json",
    "multipart/form-data"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Test Biofarma",
    "version": "1.0.0"
  },
  "paths": {
    "/health": {
      "get": {
        "security": [],
        "description": "Check if the App is Running",
        "tags": [
          "health"
        ],
        "summary": "Health Check",
        "operationId": "health",
        "responses": {
          "200": {
            "description": "Health Check",
            "schema": {
              "$ref": "#/definitions/success"
            }
          },
          "default": {
            "description": "Server Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/v1/home": {
      "post": {
        "security": [],
        "description": "Create new home data",
        "tags": [
          "home"
        ],
        "summary": "Create",
        "operationId": "createHome",
        "parameters": [
          {
            "name": "data",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/createHomeParamsBody"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Success create",
            "schema": {
              "$ref": "#/definitions/successCreate"
            }
          },
          "default": {
            "description": "Server Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/v1/home/{home_id}": {
      "get": {
        "security": [],
        "description": "Find one home data",
        "tags": [
          "home"
        ],
        "summary": "Find One",
        "operationId": "FindOneHome",
        "parameters": [
          {
            "type": "integer",
            "format": "uint64",
            "name": "home_id",
            "in": "path",
            "required": true
          },
          {
            "enum": [
              "true",
              "false"
            ],
            "type": "string",
            "default": "false",
            "description": "flag for identify include deleted data or not",
            "name": "include_deleted_data",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "Success get data",
            "schema": {
              "$ref": "#/definitions/successFindOne"
            }
          },
          "default": {
            "description": "Server Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "put": {
        "security": [],
        "description": "Update home data",
        "tags": [
          "home"
        ],
        "summary": "Update",
        "operationId": "UpdateHome",
        "parameters": [
          {
            "type": "integer",
            "format": "uint64",
            "name": "home_id",
            "in": "path",
            "required": true
          },
          {
            "name": "data",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/updateHomeParamsBody"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Success update",
            "schema": {
              "$ref": "#/definitions/success"
            }
          },
          "default": {
            "description": "Server Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Home": {
      "allOf": [
        {
          "$ref": "#/definitions/modelIdentifier"
        },
        {
          "$ref": "#/definitions/homeData"
        },
        {
          "$ref": "#/definitions/modelTrackTime"
        }
      ]
    },
    "createHomeParamsBody": {
      "type": "object",
      "required": [
        "type",
        "address"
      ],
      "properties": {
        "address": {
          "type": "string"
        },
        "type": {
          "type": "number",
          "format": "int",
          "enum": [
            36,
            45,
            72
          ]
        }
      },
      "x-go-gen-location": "operations"
    },
    "error": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer"
        },
        "message": {
          "type": "string",
          "example": "error"
        }
      }
    },
    "homeData": {
      "type": "object",
      "properties": {
        "address": {
          "type": "string"
        },
        "address_line": {
          "type": "string"
        },
        "country": {
          "type": "string"
        },
        "formatted_address": {
          "type": "string"
        },
        "latitude": {
          "type": "number"
        },
        "longitude": {
          "type": "number"
        },
        "postal_code": {
          "type": "string"
        },
        "regency": {
          "type": "string"
        },
        "subdistrict": {
          "type": "string"
        },
        "type": {
          "type": "string"
        }
      }
    },
    "modelIdentifier": {
      "type": "object",
      "required": [
        "id"
      ],
      "properties": {
        "id": {
          "type": "integer",
          "format": "uint64"
        }
      }
    },
    "modelTrackTime": {
      "type": "object",
      "properties": {
        "created_at": {
          "type": "string",
          "format": "date-time",
          "x-go-custom-tag": "gorm:\"column:created_at\"",
          "x-nullable": true,
          "x-omitempty": false
        },
        "deleted_at": {
          "type": "string",
          "format": "date-time",
          "x-go-custom-tag": "gorm:\"column:deleted_at\"",
          "x-nullable": true,
          "x-omitempty": false
        },
        "updated_at": {
          "type": "string",
          "format": "date-time",
          "x-go-custom-tag": "gorm:\"column:updated_at\"",
          "x-nullable": true,
          "x-omitempty": false
        }
      }
    },
    "success": {
      "type": "object",
      "properties": {
        "message": {
          "type": "string"
        }
      }
    },
    "successCreate": {
      "allOf": [
        {
          "$ref": "#/definitions/successCreateAllOf0"
        },
        {
          "$ref": "#/definitions/successCreateAllOf1"
        }
      ]
    },
    "successCreateAllOf0": {
      "type": "object",
      "properties": {
        "message": {
          "type": "string"
        }
      },
      "x-go-gen-location": "models"
    },
    "successCreateAllOf1": {
      "type": "object",
      "properties": {
        "home_id": {
          "type": "number",
          "format": "uint64"
        }
      },
      "x-go-gen-location": "models"
    },
    "successFindOne": {
      "allOf": [
        {
          "$ref": "#/definitions/successCreateAllOf0"
        },
        {
          "$ref": "#/definitions/successFindOneAllOf1"
        }
      ]
    },
    "successFindOneAllOf1": {
      "type": "object",
      "properties": {
        "data": {
          "$ref": "#/definitions/successFindOneAllOf1Data"
        }
      },
      "x-go-gen-location": "models"
    },
    "successFindOneAllOf1Data": {
      "allOf": [
        {
          "$ref": "#/definitions/modelIdentifier"
        },
        {
          "$ref": "#/definitions/homeData"
        },
        {
          "$ref": "#/definitions/modelTrackTime"
        }
      ],
      "x-go-gen-location": "models"
    },
    "updateHomeParamsBody": {
      "type": "object",
      "required": [
        "type",
        "address"
      ],
      "properties": {
        "address": {
          "type": "string"
        },
        "type": {
          "type": "number",
          "format": "int",
          "enum": [
            36,
            45,
            72
          ]
        }
      },
      "x-go-gen-location": "operations"
    }
  },
  "securityDefinitions": {
    "authorization": {
      "type": "apiKey",
      "name": "Authorization",
      "in": "header"
    }
  }
}`))
}
