.PHONY: run-zipcode-service
run-zipcode-service:
	@./scripts/run.sh  zipcodeservice

.PHONY: run-weather-service
run-weather-service:
	@./scripts/run.sh  weatherservice

.PHONY: build
build:
	@./scripts/build.sh  pkg
	@./scripts/build.sh  zipcodeservice
	@./scripts/build.sh  weatherservice

.PHONY: update-dependencies
update-dependencies:
	@./scripts/update-dependencies.sh  pkg
	@./scripts/update-dependencies.sh  zipcodeservice
	@./scripts/update-dependencies.sh  zipcodeservice

.PHONY: install-dependencies
install-dependencies:
	@./scripts/install-dependencies.sh  pkg
	@./scripts/install-dependencies.sh  zipcodeservice
	@./scripts/install-dependencies.sh  zipcodeservice

.PHONY: docker-compose-all-up
docker-compose-all-up:
	@docker-compose -f deployments/docker-compose/docker-compose.all.yaml up --build -d	

.PHONY: docker-compose-all-down
docker-compose-all-down:
	@docker-compose -f deployments/docker-compose/docker-compose.all.yaml down	