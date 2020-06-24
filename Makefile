.PHONY: run

run:
	go build simple-rest-api/cmd/rest-api
	./rest-api
