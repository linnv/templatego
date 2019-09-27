CURDIR := $(shell pwd)

GO        := go
GOBUILD   := GOPROXY=https://goproxy.cn GO111MODULE=on GOPATH=$(GOPATH) CGO_ENABLED=0 $(GO) build $(BUILD_FLAG)
GOTEST    := GOPATH=$(GOPATH) CGO_ENABLED=1 $(GO) test -p 3
OSLINUX   := GOARCH=amd64  GOOS=linux
TARGET= templatego

LDFLAGS += -X "github.com/linnv/logx/version.VERSION=$(shell git describe --tags --dirty)"
LDFLAGS += -X "github.com/linnv/logx/version.BUILDTIME=$(shell date '+%Y-%m-%d %H:%M:%S')"
LDFLAGS += -X "github.com/linnv/logx/version.GITHASH=$(shell git rev-parse HEAD)"
LDFLAGS += -X "github.com/linnv/logx/version.GITBRANCH=$(shell git rev-parse --abbrev-ref HEAD)"
VERSION = $(shell git describe --tags --dirty)
GITBRANCH=$(shell git rev-parse --abbrev-ref HEAD)

.PHONY: all linux clean pack

all: $(TARGET)

BUILDDIR=$(CURDIR)

CONFIGFILE="config/config.yaml.sample"

$(TARGET): 
	@mkdir -p $(BUILDDIR)
	@echo "--->>building $(BUILDDIR)/$@"
	$(OSLINUX) $(GOBUILD) -ldflags '$(LDFLAGS)' -v -o $(BUILDDIR)/$@  $(CURDIR)/main.go
	
linux:
	@mkdir -p $(BUILDDIR)
	@echo "--->>building linux app $(BUILDDIR)/$@"
	$(OSLINUX) $(GOBUILD) -ldflags '$(LDFLAGS)' -v -o $(BUILDDIR)/$(TARGET)  $(CURDIR)/main.go
		
clean: 
	@[ -f $(BUILDDIR)/$(TARGET) ] && rm $(BUILDDIR)/$(TARGET) || true

TARGETAPP=$(TARGET)-$(GITBRANCH)-$(VERSION)
pack:$(TARGET)
	@echo "--->>create dir [$(TARGETAPP)]"
	@mkdir -p $(TARGETAPP)
	@echo "--->>copy app from [$(TARGET)] to [$(TARGETAPP)]"
	@cp $(TARGET) $(TARGETAPP)
	@echo "--->>copy config from [$(CONFIGFILE)] to [$(TARGETAPP)]"
	@cp $(CONFIGFILE) $(TARGETAPP)
	@cp ./testSh/localtest.sh $(TARGETAPP)
	@echo "--->>compress target [$(TARGETAPP)]"
	tar -czf $(TARGETAPP).tar.gz  $(TARGETAPP)
	@echo "--->>remove tmp app [$(TARGETAPP)]"
	@rm -rf $(TARGETAPP)
	@echo "--->>remove app [$(TARGET)]"
	@rm $(TARGET)
	@echo "--->>done: new app built in [$(TARGETAPP).tar.gz]"

