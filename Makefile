# Go parameters
GOCMD = godep go
GOBUILD = $(GOCMD) build
GOTEST = FC_ENV=test LOG_LEVEL=2 $(GOCMD) test
GOFMT = $(GOCMD) fmt
PKG_NAME = github.com/FoxComm/libs

# Lists of services and packages to perform operations against.
SERVICE_LIST := announcer configs endpoints etcd_client logger spree utils

BUILD_LIST = $(foreach int, $(SERVICE_LIST), $(int)_build)
FMT_LIST = $(foreach int, $(SERVICE_LIST), $(int)_fmt)

# List actions
$(BUILD_LIST): %_build:
	cd $* && $(GOBUILD)

$(FMT_LIST): %_fmt:
	$(GOFMT) $(PKG_NAME)/$*/...

# Targets
.PHONY: all test clean setup

configure:
	cd $(GOPATH)
	go get -u github.com/tools/Godep
	go get -u github.com/pilu/fresh
	go get -u github.com/ddollar/forego

build: $(BUILD_LIST)

fmt: $(FMT_LIST)

test:
	$(GOTEST) -v ./... -cover
