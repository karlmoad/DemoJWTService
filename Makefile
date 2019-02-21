.PHONY: all ref build clean

DIST_PATH := dist

all: build

ref:
	go get -u

build: service.go
	go build -o $(DIST_PATH)/DemoJWTService ./service.go

clean:
	rm -f $(DIST_PATH)/*
