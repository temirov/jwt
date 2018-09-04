build:
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/cats cats/main.go
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/pangolins pangolins/main.go
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/authorize authorize/main.go

.PHONY: clean
clean:
	rm -rf ./bin

.PHONY: deploy
deploy: clean build
	sls deploy --verbose

