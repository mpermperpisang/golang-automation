@desktop
Feature: Variant Desktop Web

  @variant-dweb-1
  Scenario: Variant Desktop Web Without Examples
    Given client login as "USER" via desktop

  @variant-dweb-2
  Scenario Outline: Variant Desktop Web Without Examples
    Given client login as "<user>" via desktop

    Examples:
    | user       |
    | ENV:USER   |
    | ENV:USER_2 |

