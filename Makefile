deps:
	go mod vendor

build:
	go build -o ./bin/iterator ./cmd/iterator

run: deps build
	./bin/iterator