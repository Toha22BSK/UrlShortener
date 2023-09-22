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
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http",
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Convert long link to short link",
    "title": "UrlShortener",
    "version": "0.0.1"
  },
  "host": "localhost:8080",
  "basePath": "/api",
  "paths": {
    "/analytics/{short-url}": {
      "get": {
        "description": "Get analytics for short URL",
        "tags": [
          "analytics"
        ],
        "summary": "Get analytics for short URL",
        "operationId": "getAnalytics",
        "parameters": [
          {
            "type": "string",
            "description": "Short URL to get analytics for",
            "name": "short-url",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "object",
              "properties": {
                "redirects": {
                  "type": "number"
                }
              }
            }
          },
          "401": {
            "$ref": "#/responses/unauthorized"
          },
          "404": {
            "$ref": "#/responses/not-found"
          },
          "500": {
            "$ref": "#/responses/server-error"
          }
        }
      }
    },
    "/url": {
      "put": {
        "description": "Create short URL",
        "tags": [
          "ShortUrl"
        ],
        "summary": "Create short URL",
        "operationId": "createShortUrl",
        "parameters": [
          {
            "name": "data",
            "in": "body",
            "schema": {
              "type": "object",
              "properties": {
                "url": {
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "swagger/responses/short_url.yml"
            }
          },
          "400": {
            "$ref": "#/responses/invalid-request"
          },
          "500": {
            "$ref": "#/responses/server-error"
          }
        }
      }
    },
    "/{short_url}": {
      "get": {
        "description": "Get short URL with a redirect",
        "tags": [
          "ShortUrl"
        ],
        "summary": "Get short URL",
        "operationId": "getShortUrl",
        "parameters": [
          {
            "type": "string",
            "description": "URL to get",
            "name": "short_url",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "302": {
            "description": "Redirect to URL",
            "headers": {
              "Location": {
                "type": "string"
              }
            }
          },
          "404": {
            "$ref": "#/responses/not-found"
          },
          "500": {
            "$ref": "#/responses/server-error"
          }
        }
      }
    }
  },
  "responses": {
    "forbidden": {
      "description": "Insufficient privilege to execute action.",
      "schema": {
        "$ref": "swagger/responses/error.v1.yml"
      }
    },
    "invalid-request": {
      "description": "Invalid request",
      "schema": {
        "$ref": "swagger/responses/error.v1.yml"
      }
    },
    "not-found": {
      "description": "Not found.",
      "schema": {
        "type": "string"
      }
    },
    "server-error": {
      "description": "Internal Server Error",
      "schema": {
        "$ref": "swagger/responses/error.v1.yml"
      }
    },
    "unauthorized": {
      "description": "Unauthorized",
      "schema": {
        "$ref": "swagger/responses/error.v1.yml"
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http",
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Convert long link to short link",
    "title": "UrlShortener",
    "version": "0.0.1"
  },
  "host": "localhost:8080",
  "basePath": "/api",
  "paths": {
    "/analytics/{short-url}": {
      "get": {
        "description": "Get analytics for short URL",
        "tags": [
          "analytics"
        ],
        "summary": "Get analytics for short URL",
        "operationId": "getAnalytics",
        "parameters": [
          {
            "type": "string",
            "description": "Short URL to get analytics for",
            "name": "short-url",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "object",
              "properties": {
                "redirects": {
                  "type": "number"
                }
              }
            }
          },
          "401": {
            "description": "Unauthorized",
            "schema": {
              "$ref": "#/definitions/errorV1"
            }
          },
          "404": {
            "description": "Not found.",
            "schema": {
              "type": "string"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/errorV1"
            }
          }
        }
      }
    },
    "/url": {
      "put": {
        "description": "Create short URL",
        "tags": [
          "ShortUrl"
        ],
        "summary": "Create short URL",
        "operationId": "createShortUrl",
        "parameters": [
          {
            "name": "data",
            "in": "body",
            "schema": {
              "type": "object",
              "properties": {
                "url": {
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/shortUrl"
            }
          },
          "400": {
            "description": "Invalid request",
            "schema": {
              "$ref": "#/definitions/errorV1"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/errorV1"
            }
          }
        }
      }
    },
    "/{short_url}": {
      "get": {
        "description": "Get short URL with a redirect",
        "tags": [
          "ShortUrl"
        ],
        "summary": "Get short URL",
        "operationId": "getShortUrl",
        "parameters": [
          {
            "type": "string",
            "description": "URL to get",
            "name": "short_url",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "302": {
            "description": "Redirect to URL",
            "headers": {
              "Location": {
                "type": "string"
              }
            }
          },
          "404": {
            "description": "Not found.",
            "schema": {
              "type": "string"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/errorV1"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "errorV1": {
      "description": "Standard error format",
      "type": "object",
      "title": "Error Response",
      "properties": {
        "code": {
          "type": "string"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "shortUrl": {
      "description": "OK response with short URL",
      "type": "object",
      "title": "OK response with short URL",
      "properties": {
        "short": {
          "type": "string"
        }
      }
    }
  },
  "responses": {
    "forbidden": {
      "description": "Insufficient privilege to execute action.",
      "schema": {
        "$ref": "#/definitions/errorV1"
      }
    },
    "invalid-request": {
      "description": "Invalid request",
      "schema": {
        "$ref": "#/definitions/errorV1"
      }
    },
    "not-found": {
      "description": "Not found.",
      "schema": {
        "type": "string"
      }
    },
    "server-error": {
      "description": "Internal Server Error",
      "schema": {
        "$ref": "#/definitions/errorV1"
      }
    },
    "unauthorized": {
      "description": "Unauthorized",
      "schema": {
        "$ref": "#/definitions/errorV1"
      }
    }
  }
}`))
}
