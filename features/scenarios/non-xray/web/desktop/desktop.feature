@dweb
Feature: Variant Desktop Web

  @variant-dweb-1
  Scenario: Variant Desktop Web Without Examples
    Given there is client who wants to login as "USER" via desktop
    When client input valid data login
    Then client must be in logged home page

  @variant-dweb-2
  Scenario Outline: Variant Desktop Web With Examples
    Given there is client who wants to login as "<user>" via desktop
    When client input valid data login
    Then client must be in logged home page

    Examples:
    | user           |
    | ENV:USER_PHONE |
    | ENV:USER_2     |
