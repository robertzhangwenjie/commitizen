OUTPUT := _output
BIN := ${OUTPUT}/commitizen

clean:
	@echo "========> Cleaning all build output"
	@rm -rvf ${OUTPUT}

install:
	${BIN} install

commit:
	@go run main.go

build: clean
	go build -o ${BIN}

test:
	@echo "========> Running all tests"
	@go test ./... -v