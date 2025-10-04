.PHONY: run build utest-srv utest-repo itest migrate-up migrate-down clean

#run: build-ui run-app // не работает с текущим хендлером
#	./techUI
#
#build-ui:
#	go build -o techUI cmd/techUI/main.go

rerun-app:
	make stop-app && docker rm bs-nginx && make run-app

# тесты ППО
utest-srv:
	go test -v ./internal/tests/unitTests/serviceTests/

utest-repo:
	go test -v ./internal/tests/unitTests/repositoryTests/

gen-mocks:
	mockgen -source=./components/component-services/intfRepo/IBookRepo.go -destination=./internal/tests/unitTests/serviceTests/mocks/mockBookRepo.go --package=mocks
	mockgen -source=./components/component-services/intfRepo/ILibCardRepo.go -destination=./internal/tests/unitTests/serviceTests/mocks/mockLibCardRepo.go --package=mocks
	mockgen -source=./components/component-services/intfRepo/IRatingRepo.go -destination=./internal/tests/unitTests/serviceTests/mocks/mockRatingRepo.go --package=mocks
	mockgen -source=./components/component-services/intfRepo/IReaderRepo.go -destination=./internal/tests/unitTests/serviceTests/mocks/mockReaderRepo.go --package=mocks
	mockgen -source=./components/component-services/intfRepo/IReservationRepo.go -destination=./internal/tests/unitTests/serviceTests/mocks/mockReservationRepo.go --package=mocks

.PHONY: test
test:
	go test -v -shuffle on ./internal/tests/...

.PHONY: coverage
coverage:
	go tool cover -html ./coverage/cover.out

.PHONY: unitTests
unitTests:
	go test -v -shuffle on -p 10 ./internal/tests_for_testing/unitTests/

.PHONY: serveAllure
serveAllure:
	allure serve ./internal/tests_for_testing/unitTests/allure-results