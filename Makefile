OUTPUT := _output
BIN := ${OUTPUT}/commitizen

GO_VERSION := $(word 3, $(shell go version))
GIT_COMMIT := $(shell git rev-parse --short HEAD)
PACKAGE_VERSION := $(shell git tag | tail -1)
VERSION_PACKAGE := github.com/robertzhangwenjie/commitizen/pkg/version

GIT_TREE_STATE="dirty"
ifeq (, $(shell git status --porcelain 2>/dev/null))
	GIT_TREE_STATE="clean"
endif
# build args
GO_BUILD_FLAGS += -ldflags "-X '${VERSION_PACKAGE}.GitCommit=${GIT_COMMIT}' \
                 -X '${VERSION_PACKAGE}.BuildTime=`date "+%Y-%m-%d %H:%M:%S"`' \
                 -X '${VERSION_PACKAGE}.GoVersion=${GO_VERSION}' \
                 -X '${VERSION_PACKAGE}.GitTreeState=${GIT_TREE_STATE}' \
                 -X '${VERSION_PACKAGE}.GitVersion=${PACKAGE_VERSION}'"

clean:
	@echo "========> Cleaning all build output"
	@rm -rvf ${OUTPUT}

install: build
	${BIN} install

.PHONY: commit
commit:
	@go run main.go

build: clean
	go build ${GO_BUILD_FLAGS} -o ${BIN}

dryRun: build
	@${BIN} --dry-run

test:
	@echo "========> Running all tests"
	@go test ./... -v

test.coverage: COVERAGE_FILE := coverage.out

test.coverage: tools.gocov
	@go test ./... -v -cover -coverprofile=${COVERAGE_FILE}
	@gocov convert ${COVERAGE_FILE} | gocov-xml > cover.xml

tools.gocov:
	@go install github.com/AlekSi/gocov-xml@latest
	@go install github.com/axw/gocov/gocov@latest