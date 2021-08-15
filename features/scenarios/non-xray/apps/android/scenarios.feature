@scenarios @android
Feature: Automation Android

  @android-tc-01
  Scenario: User access action bar tab
    When user click menu
      | menu            |
      | App             |
      | Action Bar      |
      | Action Bar Tabs |
    Then validate action bar tabs is displaying
