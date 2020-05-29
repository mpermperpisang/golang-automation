@example
Feature: Example Automation
  Automation with golang and cucumber

  @example-dweb
  Scenario: Example DWEB
    Given visit dweb

  @example-mweb
  Scenario: Example MWEB
    Given visit mweb

  @example-android
  Scenario: Example Android
    Given visit android

  @example-ios
  Scenario: Example iOS
    Given visit ios

  @example-api
  Scenario: Example API
    Given client has "ENV:API_BASE_URL_4" as base api
    When client sends a GET request to "ENV:ENDPOINT"
    Then response status should be "200"

  @example-unit
  Scenario: Example Unit
    Given user has a name "Banana"
    When Testivus meet user
    Then Testivus say "Hello Banana!"

  @example-data
  Scenario: Set and Get Data
    Given some data
    When set data
    Then get data

  @example-yaml
  Scenario: Read Yaml
    Given yaml file
    When read file
    Then print contents

  @example-mapping
  Scenario: Call Function
    Given function mapping

  @example-page-struct
  Scenario: Page Object Using Struct
    Given visit dweb
    When access login url
    Then validate login page
