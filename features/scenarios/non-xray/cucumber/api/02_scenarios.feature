@scenarios @api
Feature: Automation API

  @api-tc-03
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
