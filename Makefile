all: fmt imports

fmt:
	@echo "==> Formatting code"
	@go fmt ./...

imports:
	@echo "==> Sorting imports"
	@if exist vendor (goimports -w $(shell dir /s /b *.go 2>NUL | findstr /V /C:"\\vendor\\")) else (goimports -w $(shell dir /s /b *.go))

.PHONY: all fmt imports
