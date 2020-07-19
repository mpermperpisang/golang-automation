@api @exclude
Feature: Variant API

  @variant-api-3
  Scenario: Variant API Without Body
    Given client has "ENV:ACCOUNT_BASE_URL_2" as base api
    And client login as "USER"
    And client has "ENV:API_BASE_URL_3" as base api
    When client sends a GET request to "ENV:ENDPOINT_3"
    Then response status should be "200"
    And response should have "$..data..id"
    And response "$.data.id" should be float64
    And response should have "$..data..qr_code"
    And response should have "$..data..bike_code"
    And response should have "$..data..latitude"
    And response should have "$..data..longitude"
    And response should have "$..data..status"
    And response should have "$.data.status" matching "inactive"
    And response should have "$..data..address"
    And response should have "$..data..location"
    And response should have "$..data..distance"
    And response should have "$..data..calory"
    And response should have "$..data..fare"
    And response should have "$..data..estimated_walking_time"
