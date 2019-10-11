@mobile
Feature: Variant Mobile Web

  @variant-mweb-1
  Scenario: Variant Mobile Web Without Examples
    Given there is client who wants to login as "USER" via mobile
    When client input valid data login in mobile
    Then client must be in logged mobile page

  @variant-mweb-2
  Scenario Outline: Variant Mobile Web With Examples
    Given there is client who wants to login as "<user>" via mobile
    When client input valid data login in mobile
    Then client must be in logged mobile page

    Examples:
    | user           |
    | ENV:USER_PHONE |
    | ENV:USER_2     |
