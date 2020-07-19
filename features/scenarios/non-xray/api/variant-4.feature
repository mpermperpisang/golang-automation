@api @exclude
Feature: Variant API

  @variant-api-4
  Scenario: Variant API With Body
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
    Then response status should be "401"
    And response should have "$.errors[0].message" matching "User login tidak valid"
