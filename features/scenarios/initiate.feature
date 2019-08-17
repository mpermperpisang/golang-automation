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
    Given client has "ENV:API_BASE_URL" as base api
    When client sends a GET request to "ENV:ENDPOINT"
    Then response status should be "200"

  @initiate-unit
  Scenario: Initiate Unit Testing
    Given user has a name "Banana"
    When Testivus meet user
    Then Testivus say "Hello Banana!"
