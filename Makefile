build:
	@cd booklet; go build -o ../booklet.exe

fmt:
	@cd booklet; go fmt ./...

.PHONY: build fmt
