.DEFAULT_GOAL:= run
build:
	@go build -o bin/fs

run: build
	@./bin/fs -r

test:
	@go test ./... 