BIN := volume

.PHONY: all
all: clean build

.PHONY: build
build: deps
	go build -o build/$(BIN) ./cmd/$(BIN)

.PHONY: install
install: deps
	go install ./...

.PHONY: deps
deps:
	go get -d -v ./...

.PHONY: test
test: testdeps build
	go test -v .
	go test -v ./cmd/volume

.PHONY: testdeps
testdeps:
	go get -d -v -t ./...

.PHONY: lint
lint: lintdeps
	go vet
	golint -set_exit_status ./...

.PHONY: lintdeps
lintdeps:
	go get -d -v -t .
	command -v golint >/dev/null || go get -u golang.org/x/lint/golint

.PHONY: clean
clean:
	rm -rf build
	go clean
