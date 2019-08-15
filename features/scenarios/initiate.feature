@initiate
Feature: Initiate Automation
  Try automation with golang and cucumber

  @initiate-dweb
  Scenario: Initiate DWEB
    Given visit dweb

  @initiate-mweb
  Scenario: Initiate MWEB
    Given visit mweb

  @initiate-android
  Scenario: Initiate Android
    Given visit android

  @initiate-ios
  Scenario: Initiate iOS
    Given visit ios

  @initiate-api
  Scenario: Initiate API
    When client sends a GET request to "ENV:ENDPOINT"
    Then response status should be "200"
