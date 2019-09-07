@android
Feature: Variant Android

  @variant-android-1
  Scenario: Variant Android Without Examples
    Given there is client who wants to login as "USER" via android
    When client go to login page in android
    And client input valid data login in android
    Then client must be in logged apps page
