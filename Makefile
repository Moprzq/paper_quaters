APP_NAME := paper_quarters
WEB_DIR := web
WASM := $(WEB_DIR)/$(APP_NAME).wasm
WASM_EXEC := $(shell go env GOROOT)/lib/wasm/wasm_exec.js

.PHONY: help desktop wasm serve browser test clean

help:
	@echo "Available commands:"
	@echo "  make desktop   - run the desktop app"
	@echo "  make wasm      - build browser wasm files"
	@echo "  make serve     - serve web/ at http://localhost:8080/"
	@echo "  make browser   - build wasm, then serve web/"
	@echo "  make test      - run Go tests"
	@echo "  make clean     - remove generated outputs"

desktop:
	go run ./cmd/paper-quarters

wasm:
	GOOS=js GOARCH=wasm go build -o $(WASM) ./cmd/paper-quarters
	cp "$(WASM_EXEC)" "$(WEB_DIR)/wasm_exec.js"

serve:
	go run ./cmd/serve

browser: wasm serve

test:
	go test ./...

clean:
	rm -f "$(WASM)" "$(WEB_DIR)/wasm_exec.js" "$(APP_NAME)" "$(APP_NAME).exe" serve serve.exe
