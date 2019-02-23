.PHONY: all ref build clean docker

TAG?=latest
PREFIX?=karlmoad

DIST_PATH := dist

all: build

ref:
	go get -u

build: service.go
	go build -o $(DIST_PATH)/DemoJWTService ./service.go

clean:
	rm -f $(DIST_PATH)/*

docker: clean build
	docker build --tag $(PREFIX)/demo-jwt-service:$(TAG) .



