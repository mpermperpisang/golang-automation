@api
Feature: Variant API

  @variant-api-1
  Scenario: Variant API with non array body
    Given client has "ENV:API_BASE_URL" as base api
    When client sends a GET request to "ENV:ENDPOINT"
    Then response status should be "200"
    And response should have "$.status" matching "OK"
    And response should have "$.promo_due"
    And response should have "$.tag_page"
    And response should have "$.banner"
    And response should have "$.products"
    And response should have "$.categories"
    And response should have "$.message"
    And response should have "$.facets"
    And response should have "$.labels"
    And response should have "$.related_keywords"
    And response should have "$.recommended_products"
    And response should have "$.product_deeplink"

  @variant-api-2
  Scenario: Variant API with array body
    Given client has "ENV:API_BASE_URL_2" as base api
    When client sends a GET request to "ENV:ENDPOINT_2"
    Then response status should be "200"
    And response should have "$[0].title"
    And response should have "$[0].title" matching "Jakarta"
    And response should have "$[0].location_type"
    And response should have "$[0].woeid"
    And response should have "$[0].latt_long"
