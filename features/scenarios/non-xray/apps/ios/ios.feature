@ios
Feature: Variant iOS

  @variant-ios-1 @exclude
  Scenario: Variant iOS Without Examples
    Given there is client who wants to login as "USER" via ios
    When client go to login page in ios
    And client input valid data login in ios
    Then client must be in logged ios page

  @variant-ios-2 @exclude
  Scenario Outline: Variant iOS With Examples
    Given there is client who wants to login as "<user>" via ios
    When client go to login page in ios
    And client input valid data login in ios
    Then client must be in logged ios page

    Examples:
    | user           |
    | ENV:USER_PHONE |
    | ENV:USER_2     |
