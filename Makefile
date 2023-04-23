.PHONY:

build:
	go build -o app main.go

run:
	./app server --config ./config.yaml
