@example
Feature: Example Automation
  Automation with golang and cucumber

  @example-dweb @dweb
  Scenario: Example DWEB
    Given visit dweb

  @example-mweb @mweb
  Scenario: Example MWEB
    Given visit mweb

  @example-android @android
  Scenario: Example Android
    Given visit android

  @example-android-2 @android
  Scenario: Example Android
    Given visit "LOGIN_URL" deeplink

  @example-android-3 @android
  Scenario: Example Android
    Given mitra is in home page

  @example-ios @ios @exclude
  Scenario: Example iOS
    Given visit ios

  @example-api @exclude
  Scenario: Example API
    Given client has "ENV:API_BASE_URL_4" as base api
    When client sends a GET request to "ENV:ENDPOINT"
    Then response status should be "200"

  @example-unit @exclude
  Scenario: Example Unit
    Given user has a name "Banana"
    When Testivus meet user
    Then Testivus say "Hello Banana!"

  @example-data @exclude
  Scenario: Set and Get Data
    Given some data
    When set data
    Then get data

  @example-yaml @exclude
  Scenario: Read Yaml
    Given yaml file
    When read file
    Then print contents

  @example-mapping @exclude
  Scenario: Call Function
    Given function mapping
