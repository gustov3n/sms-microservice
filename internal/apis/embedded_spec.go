// Code generated by go-swagger; DO NOT EDIT.

package apis

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
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Sms Microservive",
    "title": "Sms Microservive",
    "version": "1.0.0"
  },
  "host": "localhost:8080",
  "paths": {
    "/provider": {
      "get": {
        "description": "Health check endpoint",
        "produces": [
          "application/json"
        ],
        "tags": [
          "Provider"
        ],
        "summary": "Health check",
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "type": "object",
              "properties": {
                "response_code": {
                  "type": "string"
                },
                "response_data": {
                  "type": "array",
                  "items": {
                    "properties": {
                      "id": {
                        "description": "success retrieve data",
                        "type": "string",
                        "example": "app running well"
                      },
                      "name": {
                        "description": "success retrieve data",
                        "type": "string",
                        "example": "app running well"
                      }
                    }
                  }
                },
                "response_messege": {
                  "type": "string"
                }
              }
            }
          },
          "400": {
            "$ref": "#/responses/BadRequest"
          },
          "401": {
            "$ref": "#/responses/BadRequest"
          },
          "404": {
            "$ref": "#/responses/NotFound"
          }
        }
      },
      "put": {
        "description": "Health check endpoint",
        "produces": [
          "application/json"
        ],
        "tags": [
          "Provider"
        ],
        "summary": "Health check",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "properties": {
                "id": {
                  "type": "string",
                  "example": ""
                },
                "password": {
                  "type": "string",
                  "example": ""
                },
                "username": {
                  "type": "string",
                  "example": ""
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "type": "object",
              "properties": {
                "response_code": {
                  "type": "string"
                },
                "response_data": {
                  "type": "string"
                },
                "response_messege": {
                  "type": "string"
                }
              }
            }
          },
          "400": {
            "$ref": "#/responses/BadRequest"
          },
          "401": {
            "$ref": "#/responses/BadRequest"
          },
          "404": {
            "$ref": "#/responses/NotFound"
          }
        }
      }
    },
    "/sms": {
      "post": {
        "description": "Health check endpoint",
        "produces": [
          "application/json"
        ],
        "tags": [
          "Sms"
        ],
        "summary": "Health check",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/CreateSmsRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "type": "object",
              "properties": {
                "response_code": {
                  "type": "string"
                },
                "response_data": {
                  "type": "object",
                  "$ref": "#/definitions/SmsResponse"
                },
                "response_messege": {
                  "type": "string"
                }
              }
            }
          },
          "400": {
            "$ref": "#/responses/BadRequest"
          },
          "401": {
            "$ref": "#/responses/BadRequest"
          },
          "404": {
            "$ref": "#/responses/NotFound"
          }
        }
      }
    },
    "/sms/history": {
      "get": {
        "description": "Health check endpoint",
        "produces": [
          "application/json"
        ],
        "tags": [
          "Sms"
        ],
        "summary": "Health check",
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "type": "object",
              "properties": {
                "metadata": {
                  "type": "object",
                  "$ref": "#/definitions/Metadata"
                },
                "response_code": {
                  "type": "string"
                },
                "response_data": {
                  "type": "object",
                  "$ref": "#/definitions/SmsHistory"
                },
                "response_messege": {
                  "type": "string"
                }
              }
            }
          },
          "400": {
            "$ref": "#/responses/BadRequest"
          },
          "401": {
            "$ref": "#/responses/BadRequest"
          },
          "404": {
            "$ref": "#/responses/NotFound"
          }
        }
      }
    },
    "/template": {
      "get": {
        "description": "Health check endpoint",
        "produces": [
          "application/json"
        ],
        "tags": [
          "Template"
        ],
        "summary": "Health check",
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "type": "object",
              "properties": {
                "response_code": {
                  "type": "string"
                },
                "response_data": {
                  "type": "object",
                  "$ref": "#/definitions/GetTemplate"
                },
                "response_messege": {
                  "type": "string"
                }
              }
            }
          },
          "400": {
            "$ref": "#/responses/BadRequest"
          },
          "401": {
            "$ref": "#/responses/BadRequest"
          },
          "404": {
            "$ref": "#/responses/NotFound"
          }
        }
      },
      "post": {
        "description": "Health check endpoint",
        "produces": [
          "application/json"
        ],
        "tags": [
          "Template"
        ],
        "summary": "Health check",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/CreateTemplate"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "type": "object",
              "properties": {
                "response_code": {
                  "type": "string"
                },
                "response_data": {
                  "type": "string"
                },
                "response_messege": {
                  "type": "string"
                }
              }
            }
          },
          "400": {
            "$ref": "#/responses/BadRequest"
          },
          "401": {
            "$ref": "#/responses/BadRequest"
          },
          "404": {
            "$ref": "#/responses/NotFound"
          }
        }
      }
    },
    "/template/{template_id}": {
      "put": {
        "description": "Health check endpoint",
        "produces": [
          "application/json"
        ],
        "tags": [
          "Template"
        ],
        "summary": "Health check",
        "parameters": [
          {
            "type": "string",
            "name": "template_id",
            "in": "path",
            "required": true
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/CreateTemplate"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "type": "object",
              "properties": {
                "response_code": {
                  "type": "string"
                },
                "response_data": {
                  "type": "string"
                },
                "response_messege": {
                  "type": "string"
                }
              }
            }
          },
          "400": {
            "$ref": "#/responses/BadRequest"
          },
          "401": {
            "$ref": "#/responses/BadRequest"
          },
          "404": {
            "$ref": "#/responses/NotFound"
          }
        }
      },
      "delete": {
        "description": "Health check endpoint",
        "produces": [
          "application/json"
        ],
        "tags": [
          "Template"
        ],
        "summary": "Health check",
        "parameters": [
          {
            "type": "string",
            "name": "template_id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "type": "object",
              "properties": {
                "response_code": {
                  "type": "string"
                },
                "response_data": {
                  "type": "null"
                },
                "response_messege": {
                  "type": "string"
                }
              }
            }
          },
          "400": {
            "$ref": "#/responses/BadRequest"
          },
          "401": {
            "$ref": "#/responses/BadRequest"
          },
          "404": {
            "$ref": "#/responses/NotFound"
          }
        }
      }
    }
  },
  "definitions": {
    "CreateSmsRequest": {
      "type": "object",
      "properties": {
        "apps_name": {
          "type": "string",
          "example": ""
        },
        "phone_number": {
          "type": "string",
          "example": ""
        },
        "token": {
          "type": "string"
        },
        "type": {
          "type": "string",
          "example": ""
        }
      }
    },
    "CreateTemplate": {
      "type": "object",
      "properties": {
        "apps_name": {
          "type": "string",
          "example": ""
        },
        "name": {
          "type": "string",
          "example": ""
        },
        "text": {
          "type": "string",
          "example": ""
        },
        "type": {
          "type": "string",
          "example": ""
        }
      }
    },
    "Error": {
      "type": "object",
      "properties": {
        "code": {
          "type": "string"
        },
        "message": {
          "type": "string",
          "example": "error"
        }
      }
    },
    "GetTemplate": {
      "type": "array",
      "items": {
        "properties": {
          "apps_name": {
            "type": "string",
            "example": ""
          },
          "id": {
            "type": "string",
            "example": ""
          },
          "text": {
            "type": "string",
            "example": ""
          },
          "type": {
            "type": "string",
            "example": ""
          }
        }
      }
    },
    "Health": {
      "type": "object",
      "properties": {
        "status": {
          "type": "string",
          "example": "Good"
        }
      }
    },
    "MakingPayment": {
      "type": "array",
      "items": {
        "type": "object",
        "properties": {
          "description": {
            "type": "string",
            "example": ""
          },
          "header": {
            "type": "string",
            "example": ""
          }
        }
      }
    },
    "Metadata": {
      "type": "object",
      "properties": {
        "limit": {
          "type": "integer",
          "example": 1
        },
        "page": {
          "type": "integer",
          "example": 1
        },
        "total_row": {
          "type": "integer",
          "example": 1
        }
      }
    },
    "SmsHistory": {
      "type": "array",
      "items": {
        "properties": {
          "id": {
            "type": "string"
          },
          "phone_number": {
            "type": "string"
          },
          "request": {
            "type": "object",
            "properties": {
              "sender": {
                "type": "string",
                "example": ""
              },
              "type_messege": {
                "type": "string",
                "example": ""
              }
            }
          },
          "responses": {
            "type": "object",
            "properties": {
              "status": {
                "type": "string",
                "example": "Success"
              }
            }
          }
        }
      }
    },
    "SmsResponse": {
      "type": "object",
      "properties": {
        "apps_name": {
          "type": "string",
          "example": ""
        },
        "phone_number": {
          "type": "string",
          "example": ""
        },
        "status": {
          "type": "string",
          "example": ""
        },
        "type": {
          "type": "string",
          "example": ""
        }
      }
    }
  },
  "responses": {
    "BadRequest": {
      "description": "Unauthorized",
      "schema": {
        "$ref": "#/definitions/Error"
      }
    },
    "InternalServerError": {
      "description": "Internal server error",
      "schema": {
        "$ref": "#/definitions/Error"
      }
    },
    "NotFound": {
      "description": "The specified resource was not found",
      "schema": {
        "$ref": "#/definitions/Error"
      }
    },
    "Unauthorized": {
      "description": "Unauthorized",
      "schema": {
        "$ref": "#/definitions/Error"
      }
    }
  },
  "securityDefinitions": {
    "Bearer": {
      "type": "apiKey",
      "name": "Authorization",
      "in": "header"
    }
  },
  "tags": [
    {
      "description": "Every request and response about Provider",
      "name": "Provider"
    },
    {
      "description": "Every request and response about Sms",
      "name": "Sms"
    },
    {
      "description": "Every request and response about Template",
      "name": "Template"
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Sms Microservive",
    "title": "Sms Microservive",
    "version": "1.0.0"
  },
  "host": "localhost:8080",
  "paths": {
    "/provider": {
      "get": {
        "description": "Health check endpoint",
        "produces": [
          "application/json"
        ],
        "tags": [
          "Provider"
        ],
        "summary": "Health check",
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "type": "object",
              "properties": {
                "response_code": {
                  "type": "string"
                },
                "response_data": {
                  "type": "array",
                  "items": {
                    "$ref": "#/definitions/ResponseDataItems0"
                  }
                },
                "response_messege": {
                  "type": "string"
                }
              }
            }
          },
          "400": {
            "description": "Unauthorized",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "401": {
            "description": "Unauthorized",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "404": {
            "description": "The specified resource was not found",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "put": {
        "description": "Health check endpoint",
        "produces": [
          "application/json"
        ],
        "tags": [
          "Provider"
        ],
        "summary": "Health check",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "properties": {
                "id": {
                  "type": "string",
                  "example": ""
                },
                "password": {
                  "type": "string",
                  "example": ""
                },
                "username": {
                  "type": "string",
                  "example": ""
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "type": "object",
              "properties": {
                "response_code": {
                  "type": "string"
                },
                "response_data": {
                  "type": "string"
                },
                "response_messege": {
                  "type": "string"
                }
              }
            }
          },
          "400": {
            "description": "Unauthorized",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "401": {
            "description": "Unauthorized",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "404": {
            "description": "The specified resource was not found",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/sms": {
      "post": {
        "description": "Health check endpoint",
        "produces": [
          "application/json"
        ],
        "tags": [
          "Sms"
        ],
        "summary": "Health check",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/CreateSmsRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "type": "object",
              "properties": {
                "response_code": {
                  "type": "string"
                },
                "response_data": {
                  "type": "object",
                  "$ref": "#/definitions/SmsResponse"
                },
                "response_messege": {
                  "type": "string"
                }
              }
            }
          },
          "400": {
            "description": "Unauthorized",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "401": {
            "description": "Unauthorized",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "404": {
            "description": "The specified resource was not found",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/sms/history": {
      "get": {
        "description": "Health check endpoint",
        "produces": [
          "application/json"
        ],
        "tags": [
          "Sms"
        ],
        "summary": "Health check",
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "type": "object",
              "properties": {
                "metadata": {
                  "type": "object",
                  "$ref": "#/definitions/Metadata"
                },
                "response_code": {
                  "type": "string"
                },
                "response_data": {
                  "type": "object",
                  "$ref": "#/definitions/SmsHistory"
                },
                "response_messege": {
                  "type": "string"
                }
              }
            }
          },
          "400": {
            "description": "Unauthorized",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "401": {
            "description": "Unauthorized",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "404": {
            "description": "The specified resource was not found",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/template": {
      "get": {
        "description": "Health check endpoint",
        "produces": [
          "application/json"
        ],
        "tags": [
          "Template"
        ],
        "summary": "Health check",
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "type": "object",
              "properties": {
                "response_code": {
                  "type": "string"
                },
                "response_data": {
                  "type": "object",
                  "$ref": "#/definitions/GetTemplate"
                },
                "response_messege": {
                  "type": "string"
                }
              }
            }
          },
          "400": {
            "description": "Unauthorized",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "401": {
            "description": "Unauthorized",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "404": {
            "description": "The specified resource was not found",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "post": {
        "description": "Health check endpoint",
        "produces": [
          "application/json"
        ],
        "tags": [
          "Template"
        ],
        "summary": "Health check",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/CreateTemplate"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "type": "object",
              "properties": {
                "response_code": {
                  "type": "string"
                },
                "response_data": {
                  "type": "string"
                },
                "response_messege": {
                  "type": "string"
                }
              }
            }
          },
          "400": {
            "description": "Unauthorized",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "401": {
            "description": "Unauthorized",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "404": {
            "description": "The specified resource was not found",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/template/{template_id}": {
      "put": {
        "description": "Health check endpoint",
        "produces": [
          "application/json"
        ],
        "tags": [
          "Template"
        ],
        "summary": "Health check",
        "parameters": [
          {
            "type": "string",
            "name": "template_id",
            "in": "path",
            "required": true
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/CreateTemplate"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "type": "object",
              "properties": {
                "response_code": {
                  "type": "string"
                },
                "response_data": {
                  "type": "string"
                },
                "response_messege": {
                  "type": "string"
                }
              }
            }
          },
          "400": {
            "description": "Unauthorized",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "401": {
            "description": "Unauthorized",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "404": {
            "description": "The specified resource was not found",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "delete": {
        "description": "Health check endpoint",
        "produces": [
          "application/json"
        ],
        "tags": [
          "Template"
        ],
        "summary": "Health check",
        "parameters": [
          {
            "type": "string",
            "name": "template_id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "type": "object",
              "properties": {
                "response_code": {
                  "type": "string"
                },
                "response_data": {
                  "type": "null"
                },
                "response_messege": {
                  "type": "string"
                }
              }
            }
          },
          "400": {
            "description": "Unauthorized",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "401": {
            "description": "Unauthorized",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "404": {
            "description": "The specified resource was not found",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "CreateSmsRequest": {
      "type": "object",
      "properties": {
        "apps_name": {
          "type": "string",
          "example": ""
        },
        "phone_number": {
          "type": "string",
          "example": ""
        },
        "token": {
          "type": "string"
        },
        "type": {
          "type": "string",
          "example": ""
        }
      }
    },
    "CreateTemplate": {
      "type": "object",
      "properties": {
        "apps_name": {
          "type": "string",
          "example": ""
        },
        "name": {
          "type": "string",
          "example": ""
        },
        "text": {
          "type": "string",
          "example": ""
        },
        "type": {
          "type": "string",
          "example": ""
        }
      }
    },
    "Error": {
      "type": "object",
      "properties": {
        "code": {
          "type": "string"
        },
        "message": {
          "type": "string",
          "example": "error"
        }
      }
    },
    "GetTemplate": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/GetTemplateItems0"
      }
    },
    "GetTemplateItems0": {
      "properties": {
        "apps_name": {
          "type": "string",
          "example": ""
        },
        "id": {
          "type": "string",
          "example": ""
        },
        "text": {
          "type": "string",
          "example": ""
        },
        "type": {
          "type": "string",
          "example": ""
        }
      }
    },
    "Health": {
      "type": "object",
      "properties": {
        "status": {
          "type": "string",
          "example": "Good"
        }
      }
    },
    "MakingPayment": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/MakingPaymentItems0"
      }
    },
    "MakingPaymentItems0": {
      "type": "object",
      "properties": {
        "description": {
          "type": "string",
          "example": ""
        },
        "header": {
          "type": "string",
          "example": ""
        }
      }
    },
    "Metadata": {
      "type": "object",
      "properties": {
        "limit": {
          "type": "integer",
          "example": 1
        },
        "page": {
          "type": "integer",
          "example": 1
        },
        "total_row": {
          "type": "integer",
          "example": 1
        }
      }
    },
    "ResponseDataItems0": {
      "properties": {
        "id": {
          "description": "success retrieve data",
          "type": "string",
          "example": "app running well"
        },
        "name": {
          "description": "success retrieve data",
          "type": "string",
          "example": "app running well"
        }
      }
    },
    "SmsHistory": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/SmsHistoryItems0"
      }
    },
    "SmsHistoryItems0": {
      "properties": {
        "id": {
          "type": "string"
        },
        "phone_number": {
          "type": "string"
        },
        "request": {
          "type": "object",
          "properties": {
            "sender": {
              "type": "string",
              "example": ""
            },
            "type_messege": {
              "type": "string",
              "example": ""
            }
          }
        },
        "responses": {
          "type": "object",
          "properties": {
            "status": {
              "type": "string",
              "example": "Success"
            }
          }
        }
      }
    },
    "SmsHistoryItems0Request": {
      "type": "object",
      "properties": {
        "sender": {
          "type": "string",
          "example": ""
        },
        "type_messege": {
          "type": "string",
          "example": ""
        }
      }
    },
    "SmsHistoryItems0Responses": {
      "type": "object",
      "properties": {
        "status": {
          "type": "string",
          "example": "Success"
        }
      }
    },
    "SmsResponse": {
      "type": "object",
      "properties": {
        "apps_name": {
          "type": "string",
          "example": ""
        },
        "phone_number": {
          "type": "string",
          "example": ""
        },
        "status": {
          "type": "string",
          "example": ""
        },
        "type": {
          "type": "string",
          "example": ""
        }
      }
    }
  },
  "responses": {
    "BadRequest": {
      "description": "Unauthorized",
      "schema": {
        "$ref": "#/definitions/Error"
      }
    },
    "InternalServerError": {
      "description": "Internal server error",
      "schema": {
        "$ref": "#/definitions/Error"
      }
    },
    "NotFound": {
      "description": "The specified resource was not found",
      "schema": {
        "$ref": "#/definitions/Error"
      }
    },
    "Unauthorized": {
      "description": "Unauthorized",
      "schema": {
        "$ref": "#/definitions/Error"
      }
    }
  },
  "securityDefinitions": {
    "Bearer": {
      "type": "apiKey",
      "name": "Authorization",
      "in": "header"
    }
  },
  "tags": [
    {
      "description": "Every request and response about Provider",
      "name": "Provider"
    },
    {
      "description": "Every request and response about Sms",
      "name": "Sms"
    },
    {
      "description": "Every request and response about Template",
      "name": "Template"
    }
  ]
}`))
}