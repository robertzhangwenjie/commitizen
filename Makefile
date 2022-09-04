OUTPUT := _output
BIN := ${OUTPUT}/commitizen

GoVersion := $(word 3, $(shell go version))
GitCommit := $(shell git rev-parse --short HEAD)
BuildTime := $(shell date "+%Y-%m-%d %H:%M:%S")
GitVersion := $(shell git tag | tail -1)
VERSION_PACKAGE := github.com/robertzhangwenjie/commitizen/pkg/version

GIT_TREE_STATE=dirty
ifeq (, $(shell git status --porcelain 2>/dev/null))
	GIT_TREE_STATE=clean
endif

export GoVersion GitCommit PACKAGE_VERSION GIT_TREE_STATE BuildTime GitVersion

# build args
GO_BUILD_FLAGS += -ldflags "-X '${VERSION_PACKAGE}.GitCommit=${GitCommit}' \
                 -X '${VERSION_PACKAGE}.BuildTime=${BuildTime}' \
                 -X '${VERSION_PACKAGE}.GoVersion=${GoVersion}' \
                 -X '${VERSION_PACKAGE}.GitTreeState=${GIT_TREE_STATE}' \
                 -X '${VERSION_PACKAGE}.GitVersion=${GitVersion}'"


clean:
	@echo "========> Cleaning all build output"
	@rm -rvf ${OUTPUT}

install: build
	${BIN} install

.PHONY: commit
commit:
	@go run main.go

build: clean
	@echo "========> Building binary"
	go build ${GO_BUILD_FLAGS} -o ${BIN}

build.multiarch:
	@goreleaser build --snapshot --rm-dist

dryRun: build
	@${BIN} --dry-run

test:
	@echo "========> Running all tests"
	@go test ./... -v


test.coverage: 
	@go test ./... -v  -coverprofile=coverage.out
