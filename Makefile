.PHONY: run test build clean

run:
	go run . -length 20 -count 3

test:
	go test ./... -v

build:
	go build -o bin/password-gen .

clean:
	rm -rf bin/
