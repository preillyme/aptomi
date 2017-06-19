.PHONY: default
default: clean build test

.PHONY: vendor
vendor:
	glide install --strip-vendor

.PHONY: test
test:
	go test -v ./cmd/...
	go test -v ./pkg/...
	@echo "\nAll tests passed"

.PHONY: clean-run-noop
clean-run-noop:
	$$(go env GOPATH)/bin/aptomi policy apply --noop

.PHONY: smoke
smoke: test build install clean-run-noop

.PHONY: build
build:
	go build -i -o aptomi

.PHONY: install
install:
	go install

.PHONY: fmt
fmt:
	go fmt $$(go list ./... | grep -v /vendor/)

.PHONY: vet
vet:
	go tool vet -all -shadow main.go
	go tool vet -all -shadow ./cmd ./pkg

.PHONY: lint
lint:
	$$(go env GOPATH)/bin/golint $$(go list ./... | grep -v /vendor/)

.PHONY: validate
validate: fmt vet lint
	@echo "\nAll validations passed"

.PHONY: clean
clean:
	-rm -f aptomi
	go clean -r -i
