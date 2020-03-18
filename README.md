# golang-automation

**Description**<br/>
The real development of https://github.com/mpermperpisang/go-cucumber<br/>
Basic installation can be read in https://medium.com/@mpermperpisang/recipe-to-boil-web-automation-with-go-language-98b715800d70

**Setup env**<br/>
`cp env.sample .env`<br/>
`cp capabilities-web.properties.sample capabilities-web.properties`<br/>
`cp capabilities-android.properties.sample capabilities-android.properties`<br/>
`cp capabilities-ios.properties.sample capabilities-ios.properties`

**Install Package**<br/>
`$ make package`

**Setup and running Docker selenium-hub and browser**<br/>
`$ make`

**Running**<br/>
`godog --tags=@example`<br/>
`godog --tags=~@example`<br/>
`godog --tags="@example && @example-dweb"`<br/>
`godog --tags="@example && ~@example-dweb"`<br/>
`godog --tags=@example,@example-dweb`

**Linter**<br/>
`golint`

**Contact**<br/>
`mpermperpisang@gmail.com`

**Reference**<br/>
- https://medium.com/rungo/unit-testing-made-easy-in-go-25077669318
