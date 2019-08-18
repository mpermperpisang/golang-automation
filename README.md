# golang-automation

**Description**<br/>
The real development of https://github.com/mpermperpisang/go-cucumber<br/>
Basic installation can be read in https://medium.com/@mpermperpisang/recipe-to-boil-web-automation-with-go-language-98b715800d70

**Setup env**<br/>
`cp env.sample .env`<br/>
`cp capabilities-android.properties.sample capabilities-android.properties`<br/>
`cp capabilities-ios.properties.sample capabilities-ios.properties`

**Install Package**<br/>
`$ make vendor-prepare`

**Setup and running Docker selenium-hub and browser**<br/>
`$ make`

**Running**<br/>
`godog --tags=@initiate`

**Linter**<br/>
`golint`

**Contact**<br/>
`mpermperpisang@gmail.com`
