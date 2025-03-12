export const swaggerSpec = {
  openapi: "3.0.0",
  info: {
    title: "User Time Tracking API",
    version: "1.0.0",
    description: "API for tracking user time entries"
  },
  servers: [
    {
      url: "http://localhost:3000",
      description: "Development server"
    }
  ],
  paths: {
    "/user-times": {
      get: {
        summary: "Get all time records",
        description: "Retrieves all user time records from the database",
        responses: {
          "200": {
            description: "A list of time records",
            content: {
              "application/json": {
                schema: {
                  type: "object",
                  properties: {
                    success: {
                      type: "boolean",
                      example: true
                    },
                    count: {
                      type: "integer",
                      example: 2
                    },
                    data: {
                      type: "array",
                      items: {
                        $ref: "#/components/schemas/TimeRecord"
                      }
                    }
                  }
                }
              }
            }
          },
          "500": {
            description: "Server error",
            content: {
              "application/json": {
                schema: {
                  $ref: "#/components/schemas/Error"
                }
              }
            }
          }
        }
      },
      post: {
        summary: "Create a new time record",
        description: "Creates a new time record with the provided user and time information",
        requestBody: {
          required: true,
          content: {
            "application/json": {
              schema: {
                $ref: "#/components/schemas/TimeRecordInput"
              }
            }
          }
        },
        responses: {
          "200": {
            description: "Time record created successfully",
            content: {
              "application/json": {
                schema: {
                  type: "object",
                  properties: {
                    success: {
                      type: "boolean",
                      example: true
                    },
                    message: {
                      type: "string",
                      example: "Data received successfully"
                    },
                    data: {
                      $ref: "#/components/schemas/TimeRecordInput"
                    },
                    timeSpent: {
                      type: "object",
                      properties: {
                        minutes: {
                          type: "integer",
                          example: 360
                        },
                        seconds: {
                          type: "integer",
                          example: 45
                        },
                        hours: {
                          type: "integer",
                          example: 6
                        },
                        formatted: {
                          type: "string",
                          example: "6h 0m 45s"
                        }
                      }
                    },
                    savedToDatabase: {
                      type: "boolean",
                      example: true
                    }
                  }
                }
              }
            }
          },
          "400": {
            description: "Invalid input",
            content: {
              "application/json": {
                schema: {
                  $ref: "#/components/schemas/Error"
                }
              }
            }
          }
        }
      },
      delete: {
        summary: "Delete a time record",
        description: "Deletes a time record by its ID",
        parameters: [
          {
            name: "id",
            in: "query",
            required: true,
            description: "ID of the time record to delete",
            schema: {
              type: "integer",
              example: 1
            }
          }
        ],
        responses: {
          "200": {
            description: "Record deleted successfully",
            content: {
              "application/json": {
                schema: {
                  type: "object",
                  properties: {
                    success: {
                      type: "boolean",
                      example: true
                    },
                    message: {
                      type: "string",
                      example: "Record with ID 1 deleted successfully"
                    }
                  }
                }
              }
            }
          },
          "400": {
            description: "Invalid input",
            content: {
              "application/json": {
                schema: {
                  $ref: "#/components/schemas/Error"
                }
              }
            }
          },
          "404": {
            description: "Record not found",
            content: {
              "application/json": {
                schema: {
                  type: "object",
                  properties: {
                    success: {
                      type: "boolean",
                      example: false
                    },
                    error: {
                      type: "string",
                      example: "Record not found"
                    }
                  }
                }
              }
            }
          },
          "500": {
            description: "Server error",
            content: {
              "application/json": {
                schema: {
                  $ref: "#/components/schemas/Error"
                }
              }
            }
          }
        }
      }
    }
  },
  components: {
    schemas: {
      TimeRecordInput: {
        type: "object",
        required: ["user", "startTime", "endTime"],
        properties: {
          user: {
            type: "string",
            example: "Victor"
          },
          startTime: {
            type: "string",
            example: "12/03/2025, 00:52:21"
          },
          endTime: {
            type: "string",
            example: "12/03/2025, 06:52:21"
          }
        }
      },
      TimeRecord: {
        type: "object",
        properties: {
          id: {
            type: "integer",
            example: 1
          },
          user: {
            type: "string",
            example: "Victor"
          },
          start_time: {
            type: "string",
            example: "12/03/2025, 00:52:21"
          },
          end_time: {
            type: "string",
            example: "12/03/2025, 06:52:21"
          },
          hours_spent: {
            type: "integer",
            example: 6
          },
          minutes_spent: {
            type: "integer",
            example: 0
          },
          seconds_spent: {
            type: "integer",
            example: 0
          },
          created_at: {
            type: "string",
            example: "2023-03-15 14:30:45"
          }
        }
      },
      Error: {
        type: "object",
        properties: {
          success: {
            type: "boolean",
            example: false
          },
          error: {
            type: "string",
            example: "Error message"
          }
        }
      }
    }
  }
};