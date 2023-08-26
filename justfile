default: build

build:
	go build main.go

prod:
	@go build -ldfflags "-s -w" main.go

run: build
	./main

get:
	go mod tidy

clean:
    rm main
