@scenarios @dweb
Feature: Automation Desktop Web

  @dweb-tc-01
  Scenario: Non login user follow mpermperpisang.medium.com
    Given user visit "mpermperpisang.medium.com" link
    When user click "Follow" button
    Then validate login warning is displaying
