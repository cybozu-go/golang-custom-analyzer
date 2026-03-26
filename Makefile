TARGET = custom-checker eventuallycheck restrictpkg

.PHONY: all
all: check-generate test build

.PHONY: setup
setup:
	aqua i -l

.PHONY: check-generate
check-generate:
	go mod tidy
	git diff --exit-code --name-only

.PHONY: build
build: $(TARGET)

$(TARGET):
	CGO_ENABLED=0 go build ./cmd/$@

.PHONY: clean
clean:
	-rm $(TARGET)

.PHONY: lint
lint:
	test -z "$$(gofmt -s -l . | grep -v '^vendor' | tee /dev/stderr)"
	staticcheck ./...
	go vet ./...

.PHONY: test
test: lint
	go test -race -v ./...
