short-test:
	@go test -race -short ./...

test:
	@go test -race ./...

lint:
	gofumpt -w ./.. ; golangci-lint run


install-go-tools:
	cat tools.go | grep _ | grep \".*\" -o | xargs -tI % go install %


MOCKGEN_DESTINATION := pkg/mocks
MOCK_SOURCES := pkg/bookly/auth.go pkg/bookly/offer.go pkg/bookly/offer.go pkg/bookly/user.go

mocks: ${MOCK_SOURCES}
	@rm -rf ${MOCKGEN_DESTINATION}/*
	@echo "Generating mocks..."
	@for file in $^; do mockgen -source=$$file -destination=${MOCKGEN_DESTINATION}/$$file; done

