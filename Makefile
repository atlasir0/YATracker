.PHONY: build run clean

build:
	go build -o bin/main ./cmd/...

run: build
	./bin/main
	rm -rf bin/*

clean:
	rm -rf bin/*
