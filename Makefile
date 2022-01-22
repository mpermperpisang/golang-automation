# Parameters
GOCMD=go
GOGET=$(GOCMD) get
GOINSTALL=$(GOCMD) install
GOMOD=$(GOCMD) mod
GOENV=$(GOCMD) env
DOCKERCMD=docker
DOCKERBUILD=$(DOCKERCMD) build
DOCKERRUN=$(DOCKERCMD) run
DOCKERCONTAINER=$(DOCKERCMD) container
DOCKERRM=$(DOCKERCMD) rm
DOCKERRMI=$(DOCKERCMD) rmi

# kill selenium port
kill-port:
	@kill -9 $$(lsof -t -i:4444)
	@echo "Port 4444 is killed"

docker-rm: container-rm images-rm

# remove all docker container
container-rm:
	$(DOCKERRM) -vf $(docker ps -aq)
	@echo "Docker container successfully removed"

# remove all docker images
images-rm:
	$(DOCKERRMI) -f $(docker images -aq)
	@echo "Docker images successfully removed"

# run selenium server
selenium: run-selenium-hub run-selenium-node

run-selenium-hub:
	$(DOCKERRUN) -d -p 4444:4444 --name selenium-hub selenium/hub
	@echo "Run selenium-hub successfully"

run-selenium-node:
	$(DOCKERRUN) -d --link selenium-hub:hub selenium/node-chrome
	$(DOCKERRUN) -d --link selenium-hub:hub selenium/node-firefox
	@echo "Run selenium-hub node browser successfully"

# download dependencies
deps:
	$(GOINSTALL) github.com/cucumber/godog/cmd/godog@v0.12.0
	$(GOGET) -d github.com/onsi/ginkgo/v2/ginkgo
	$(GOINSTALL) github.com/onsi/ginkgo/v2/ginkgo
	$(GOMOD) download
	$(GOMOD) tidy
	@echo "Get package successfully"

# copy paste rename .sample
cp: env properties

env:
	cp env.sample .env

properties:
	cp capabilities-android.properties.sample capabilities-android.properties
	cp capabilities-ios.properties.sample capabilities-ios.properties
	cp capabilities-web.properties.sample capabilities-web.properties

api-godog:
	godog --tags=@api --random --format=cucumber > test/report/cucumber_report.json

api-ginkgo:
	ginkgo features/scenarios/non-xray/non-cucumber/api -p --randomize-all
