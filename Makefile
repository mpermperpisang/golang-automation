all: run-package selenium-hub selenium-browser

build-package:
	docker build -t golang-automation-docker .
	@echo "Build golang-automation-docker successfully"

run-package:
	docker run golang-automation-docker
	@echo "Run golang-automation-docker successfully"

kill-port:
	@kill -9 $$(lsof -t -i:4545)
	@echo "Port 4545 is killed"

docker-clean:
	@docker container rm $$(docker ps -aq) -f
	@echo "Docker successfully removed"

selenium-hub:
	docker run -d -p 4545:4444 --name selenium-hub selenium/hub
	@echo "Docker selenium-hub is running"

selenium-browser:
	docker run -d --link selenium-hub:hub selenium/node-chrome
	docker run -d --link selenium-hub:hub selenium/node-firefox
	@echo "Docker browser is running"
