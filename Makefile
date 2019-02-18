.PHONY: all ref build clean

DIST_PATH := dist
SITE_DIST := $(DIST_PATH)/site

all: clean build

ref:
	go get -u

build: clean service.go
	go build -o $(DIST_PATH)/DemoJWTService ./service.go

clean:
	rm -f $(DIST_PATH)/WorkspaceService
	rm -rf $(SITE_DIST)
	mkdir -p $(DIST_PATH)
	mkdir -p $(SITE_DIST)
