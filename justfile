default: build

build:
	go build .

prod:
	@go build -ldfflags "-s -w" .

run: build
	./matrix-go

get:
	go mod tidy

clean:
    rm matrix-go
