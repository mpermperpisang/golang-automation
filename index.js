var reporter = require('cucumber-html-reporter');

var options = {
  theme: 'bootstrap',
  jsonFile: 'test/report/cucumber_report.json',
  output: 'test/report/cucumber_report.html',
  reportSuiteAsScenarios: true,
  scenarioTimestamp: true,
  launchReport: true,
  brandTitle: "MperMperPisang",
  name: "Golang-Automation",
  storeScreenshots: false
};

reporter.generate(options);
