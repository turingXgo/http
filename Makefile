run: build
	./main

build:
	@echo "\033[0;32m[COMPILING]\033[0;0m"
	go build -o main cmd/main.go
	@echo "\033[0;32m[COMPILED]\033[0;0m"

clean:
	rm main