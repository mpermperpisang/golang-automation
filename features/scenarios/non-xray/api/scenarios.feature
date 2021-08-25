@scenarios @api
Feature: Automation API

  @api-tc-01
  Scenario Outline: Hit /automation endpoint without body
    Given client sends a POST to "/logout"
    When client sends a <method> to "/automation"
    Then validate status is <status>
    And validate "$.status" is "success"
    And validate "$.method" is "<method>"
    And validate "$.endpoint" is "/automation"

    Examples:
      | method | status |
      | GET    | 200    |
      | POST   | 201    |
      | PATCH  | 200    |
      | PUT    | 200    |
      | DELETE | 200    |

  @api-tc-02
  Scenario: Hit /automation endpoint with body
    Given client sends a POST to "/logout"
    When client sends a POST to "/automation" with body:
    """
    {
      "method": "POST"
    }
    """
    Then validate status is 201
    And validate "$.status" is "success"
    And validate "$.method" is "POST"
    And validate "$.endpoint" is "/automation"

  @api-tc-03
  Scenario Outline: Hit /automation endpoint with env variable
    Given client sends a POST to "/logout"
    When client sends a <method> to "ENV:ENDPOINT_AUTOMATION"
    Then validate status is <status>
    And validate "$.status" is "success"
    And validate "$.method" is "<method>"
    And validate "$.endpoint" is "/automation"

    Examples:
      | method | status |
      | GET    | 200    |
      | POST   | 201    |
      | PATCH  | 200    |
      | PUT    | 200    |
      | DELETE | 200    |