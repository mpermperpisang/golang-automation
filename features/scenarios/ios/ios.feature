@ios
Feature: Variant iOS

  @variant-dweb-1
  Scenario: Variant iOS Without Examples
    Given there is client who wants to login as "USER" via desktop
    When client input valid data login
    Then client must be in logged home page
