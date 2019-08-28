@mobile
Feature: Variant Mobile Web

  @variant-mweb-1
  Scenario: Variant Mobile Web Without Examples
    Given client login as "USER" via mobile

  @variant-mweb-2
  Scenario Outline: Variant Mobile Web Without Examples
    Given client login as "<user>" via mobile

    Examples:
    | user       |
    | ENV:USER   |
    | ENV:USER_2 |
