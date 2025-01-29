.DEFAULT_GOAL:= run
build:
	@go build -o bin/fs

run: build
	@./bin/fs -on bodyLine3

test:
	@go test ./... 