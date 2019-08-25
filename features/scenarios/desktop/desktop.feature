@desktop
Feature: Variant Desktop Web

  @variant-dweb-1
  Scenario Outline: Logged user fill suggestion for BukaBike new location from DWEB
    Given client has been logged in with "<user>" user
    And visit BukaBike page
    When user click "yes" button for BukaBike
    And user input "valid" location suggestion form
    And user click "suggestion" button for BukaBike
    Then user must see suggestion success

    Examples:
      | user   |
      | user_1 |
      | user_2 |
