go build: n2elastic

all: clean dependencies test build

re: clean test build

clean: 
	@echo "Cleanning..."
	-rm -f $(BINARY_NAME)
	-rm -f $(BINARY_UNIX)
	-find . -type d -name mocks -exec rm -rf \{} +	
	-$(GOCLEAN) -i
	@echo "Done cleanning."
	
test: 
	$(GO_VARS) $(GOTEST) `$(GO_VARS) $(GOLIST) ./... | grep -v vendor` && echo -e "\nTesting is passed."	
	$(GOTEST) -v ./...
	echo -e "Test Passed."

dependencies:
	@echo "Getting dependencies..."
	$(GOMOD) download
	$(GOMOD) vendor
	@echo "Done getting dependencies."

n2elastic:
	@echo "Building n2elastic"
	@echo "Installing vendors..."
	go install ./vendor/...
	@echo "Building..."
	$(GOBUILD) ./...
	$(GOBUILD) -o $(BINARY_NAME) -v

serve:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME) serve

update-dependencies: dependencies
	$(GOGET) -u
	
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v

GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOLIST=$(GOCMD) list
GOTEST=$(GOCMD) test
GOMOD=GO111MODULE=on $(GOCMD) mod
GOGET=GO111MODULE=on $(GOCMD) get
BINARY_NAME=n2elastic
BINARY_UNIX=$(BINARY_NAME)_unix
