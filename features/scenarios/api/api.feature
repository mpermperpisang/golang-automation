@api
Feature: Variant API

  @variant-api-1
  Scenario: Variant API with non array body
    Given client has "ENV:ACCOUNT_BASE_URL_2" as base api
    And client login as "USER"
    And client has "ENV:API_BASE_URL_3" as base api
    When client sends a GET request to "ENV:ENDPOINT"
    Then response status should be "200"
    And response should have "$.data..active"
    And response should have "$.data..assurance"

  @variant-api-2
  Scenario: Variant API with array body
    Given client has "ENV:API_BASE_URL_2" as base api
    When client sends a GET request to "ENV:ENDPOINT_2"
    Then response status should be "200"
    And response should have "$[0]..title"
    And response should have "$[0].title" matching "Jakarta"
    And response should have "$[0]..location_type"
    And response should have "$[0].location_type" matching "City"
    And response should have "$[0]..woeid"
    And response should have "$[0]..latt_long"
    And response should have "$[0].latt_long" matching "-6.171440,106.827820"

@variant-api-3
  Scenario: Variant API with authentication without body
    Given client has "ENV:ACCOUNT_BASE_URL_2" as base api
    And client login as "USER"
    And client has "ENV:API_BASE_URL_3" as base api
    When client sends a GET request to "ENV:ENDPOINT_3"
    Then response status should be "200"
    And response should have "$..data..id"
    And response should have "$..data..qr_code"
    And response should have "$..data..bike_code"
    And response should have "$..data..latitude"
    And response should have "$..data..longitude"
    And response should have "$..data..status"
    And response should have "$.data.status" matching "available"
    And response should have "$..data..address"
    And response should have "$..data..location"
    And response should have "$..data..distance"
    And response should have "$..data..calory"
    And response should have "$..data..fare"
    And response should have "$..data..estimated_walking_time"

@variant-api-4
  Scenario: Variant API with authentication with body
    Given client has "ENV:ACCOUNT_BASE_URL" as base api
    And client login as "USER"
    And client has "ENV:API_BASE_URL" as base api
    When client sends a POST request to "ENV:ENDPOINT_4" with body:
    """
    {
      "name": "ENV:NAME",
      "start_time": "ENV:START_TIME",
      "end_time": "ENV:END_TIME",
      "price": ENV:PRICE,
      "email_stakeholders": "ENV:EMAIL_STAKEHOLDERS",
      "event_id": ENV:EVENT_ID
    }
    """
    Then response status should be "422"
