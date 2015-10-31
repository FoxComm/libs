# Go parameters
GOCMD = go
GOBUILD = $(GOCMD) build
GOINSTALL = $(GOCMD) install
GOTEST = FC_ENV=test LOG_LEVEL=2 $(GOCMD) test
GOFMT = $(GOCMD) fmt
PKG_NAME = github.com/FoxComm/libs

# Lists of services and packages to perform operations against.
SERVICE_LIST := announcer configs endpoints etcd_client logger spree utils

BUILD_LIST = $(foreach int, $(SERVICE_LIST), $(int)_build)
FMT_LIST = $(foreach int, $(SERVICE_LIST), $(int)_fmt)
INSTALL_LIST = $(foreach int, $(SERVICE_LIST), $(int)_install)

# List actions
$(BUILD_LIST): %_build:
	cd $* && $(GOBUILD)

$(FMT_LIST): %_fmt:
	$(GOFMT) $(PKG_NAME)/$*/...
	
$(INSTALL_LIST): %_install:
	cd $* && $(GOINSTALL)

# Targets
.PHONY: all test clean setup

configure:
	gpm install

build: $(BUILD_LIST)

fmt: $(FMT_LIST)

install: $(INSTALL_LIST)

test:
	$(GOTEST) -v ./... -cover
