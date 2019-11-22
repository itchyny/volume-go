BIN := volume
export GO111MODULE=on

.PHONY: all
all: clean build

.PHONY: build
build:
	go build -o $(BIN) ./cmd/$(BIN)

.PHONY: install
install:
	go install ./...

.PHONY: test
test: build
	go test -v .
	go test -v ./cmd/$(BIN)

.PHONY: lint
lint: lintdeps
	go vet ./...
	golint -set_exit_status ./...

.PHONY: lintdeps
lintdeps:
	cd && go get golang.org/x/lint/golint

.PHONY: clean
clean:
	rm -rf $(BIN)
	go clean
