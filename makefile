build:
	go build -o main cmd/main.go

run: build
	./main