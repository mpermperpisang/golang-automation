@api
Feature: Variant API

  @variant-api-5
  Scenario: Variant API With Collect JSON Path
    Given client has "ENV:API_BASE_URL_5" as base api
    When client sends a GET request to "ENV:ENDPOINT_6"
    Then response status should be "200"
    And response should have "$..message"
    And response should have "$.message" matching "get called"
