.PHONY: deps clean build

deps:
	go get -u ./...

clean: 
	rm -rf ./TwitterAnalysis/TwitterAnalysis
	
build:
	GOOS=linux GOARCH=amd64 go build -o ./TwitterAnalysis/TwitterAnalysis ./TwitterAnalysis