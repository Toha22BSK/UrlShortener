clean:
	rm -rf gen

swagger: clean
	mkdir gen
	swagger generate server --with-flatten=minimal -t gen -f ./swagger.yaml --exclude-main -A backendCore
	sqlc generate

sqlgen:
	sqlc generate

import:
	go mod tidy
	go mod vendor

build:import
	go build -o ./service main.go

