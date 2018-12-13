# Commnads declare
GOCMD=go
GOTEST=$(GOCMD) test
GOBUILD=$(GOCMD) build

# Params define
MAIN_PATH=.
PACKAGE_PATH=pkg
PACKAGE_BIN_PATH=bin
BIN=grbac-server

default: clean build

test: 
	- $(GOTEST) ./... -v

build: 
	# building
	mkdir $(PACKAGE_PATH)
	mkdir $(PACKAGE_PATH)/$(PACKAGE_BIN_PATH)
	cd $(MAIN_PATH) && CGO_ENABLE=false GOOS=darwin GOARCH=amd64 $(GOBUILD) -o $(BIN) 
	mv "$(MAIN_PATH)/$(BIN)" $(PACKAGE_PATH)/$(PACKAGE_BIN_PATH)

# pack:
# 	# packing
# 	cd $(PACKAGE_PATH) && tar -zcvf $(FILENAME) ./*

clean:
	echo "cleaning no shell command"
	rm -fr ${PACKAGE_PATH}