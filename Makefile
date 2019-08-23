# I don't really understand Makefile and docker. Have a lot of things to learn
all: docker-clean docker-hub docker-browser

vendor-prepare:
	@go get github.com/DATA-DOG/godog/cmd/godog
	@go get -t -d github.com/tebeka/selenium
	@go get github.com/sclevine/agouti
	@go get github.com/joho/godotenv
	@go get -u github.com/logrusorgru/aurora
	@go get -u github.com/magiconair/properties
	@go get github.com/yalp/jsonpath
	@go get -u golang.org/x/lint/golint
	@go get github.com/docopt/docopt-go
	@echo "Package installed"

kill-port:
	@kill -HUP $$(lsof -t -i:4545)
	@echo "Port 4545 is killed"

docker-clean:
	@docker container rm $$(docker ps -aq) -f
	@echo "Docker successfully removed"

docker-hub:
	docker run -d -p 4545:4444 --name selenium-hub selenium/hub
	@echo "Docker selenium-hub is running"

docker-browser:
	docker run -d --link selenium-hub:hub selenium/node-chrome
	docker run -d --link selenium-hub:hub selenium/node-firefox
	@echo "Docker browser is running"
