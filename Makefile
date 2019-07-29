.PHONY: build clean deploy

build:
	dep ensure -v
	env GOOS=linux go build -a -ldflags="-s -w" -o bin/ranks src/ranks/main.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	sls deploy --verbose
