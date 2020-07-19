@api @exclude
Feature: Variant API

  @variant-api-2
  Scenario: Variant API With Array Body
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
