@api @exclude
Feature: Variant API

  @variant-api-1
  Scenario: Variant API With Non Array Body
    Given client has "ENV:ACCOUNT_BASE_URL_2" as base api
    And client login as "USER"
    And client has "ENV:API_BASE_URL_3" as base api
    When client sends a GET request to "ENV:ENDPOINT"
    Then response status should be "200"
    And response should have "$.data..active"
    And response should have "$.data..assurance"
