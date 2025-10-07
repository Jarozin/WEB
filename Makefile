.PHONY: run build utest-srv utest-repo itest migrate-up migrate-down clean

run-app: build-all
	docker compose up -d bs-app-main bs-app-inst1 bs-app-inst2 bs-app-mirror1 \
 		bs-postgres-master bs-postgres-slave bs-mongo bs-redis bs-nginx bs-pgadmin bs-react

build-all:
	docker compose build

stop-app:
	docker stop bs-app-main bs-app-inst1 bs-app-inst2 bs-app-mirror1 \
		bs-postgres-master bs-postgres-slave bs-mongo bs-redis bs-nginx bs-pgadmin bs-react

rerun-app:
	make stop-app && docker rm bs-nginx && make run-app

get-swagger:
	swag init -g cmd/app/main.go -o ./docs_swagger

gen-mocks:
	mockgen -source=./components/component-services/intfRepo/IBookRepo.go -destination=./internal/tests/unitTests/serviceTests/mocks/mockBookRepo.go --package=mocks
	mockgen -source=./components/component-services/intfRepo/ILibCardRepo.go -destination=./internal/tests/unitTests/serviceTests/mocks/mockLibCardRepo.go --package=mocks
	mockgen -source=./components/component-services/intfRepo/IRatingRepo.go -destination=./internal/tests/unitTests/serviceTests/mocks/mockRatingRepo.go --package=mocks
	mockgen -source=./components/component-services/intfRepo/IReaderRepo.go -destination=./internal/tests/unitTests/serviceTests/mocks/mockReaderRepo.go --package=mocks
	mockgen -source=./components/component-services/intfRepo/IReservationRepo.go -destination=./internal/tests/unitTests/serviceTests/mocks/mockReservationRepo.go --package=mocks

# тесты тестирования
tests:
	./run_tests.sh

migrate-up:
	./migrate.sh up

migrate-down:
	./migrate.sh down

mmigrate-up:
	cd ./components/component-repo-mongo/impl/migrations && migrate-mongo up
	cd data/mydatasets/ && python to_mongodb.py

mmigrate-down:
	cd ./components/component-repo-mongo/impl/migrations && migrate-mongo down

# docker inspect --format='{{range .NetworkSettings.Networks}}{{.MacAddress}}{{end}}' $INSTANCE_ID -- посмотреть ip адрес сервера