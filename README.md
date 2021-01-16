# golang-automation

**Description**<br/>
Basic installation can be read in https://medium.com/@mpermperpisang/recipe-to-boil-web-automation-with-go-language-98b715800d70

[![SonarCloud](https://sonarcloud.io/images/project_badges/sonarcloud-white.svg)](https://sonarcloud.io/dashboard?id=mpermperpisang_golang-automation)

**Setup env**<br/>
`$ cp env.sample .env`<br/>
`$ cp capabilities-web.properties.sample capabilities-web.properties`<br/>
`$ cp capabilities-android.properties.sample capabilities-android.properties`<br/>
`$ cp capabilities-ios.properties.sample capabilities-ios.properties`

**Build Package**<br/>
Only for once<br/>
`$ make build-package`

**Run Package**<br/>
`$ make run-package`

**Kill Port 4545**<br/>
`$ make kill-port`

**Setup and running Docker selenium-hub and browser**<br/>
`$ make docker-clean`<br/>
`$ make selenium-hub`<br/>
`$ make selenium-browser`<br/>
or<br/>
`$ make`

**Running**<br/>
This command will not import automation run result to Xray<br/>
`$ ./godog -t "@example"`<br/>
`$ ./godog -t "~@example"`<br/>
`$ ./godog -t "@example && @example-dweb"`<br/>
`$ ./godog -t "@example && ~@exclude"`<br/>
`$ ./godog -t "@example,@example-dweb"`<br/>

**Running & Export Xray**<br/>
This command will do import automation run result to Xray<br/>
`$ ./godog -x TEST_EXECUTION_ID -t "@example"`

**Running & Open Generate HTML Report**<br/>
`$ ./godog -r -t "@example"`

**Open Generate HTML Report**<br/>
`$ node index.js`

**Linter**<br/>
`$ golint`

**Contact**<br/>
`mpermperpisang@gmail.com`

**Reference**<br/>
- https://medium.com/rungo/unit-testing-made-easy-in-go-25077669318
- https://github.com/gkushang/cucumber-html-reporter
- https://goinbigdata.com/testing-go-code-with-testify/
