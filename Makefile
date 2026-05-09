APP_NAME := paper_quarters
WEB_DIR := web
WASM := $(WEB_DIR)/$(APP_NAME).wasm
WASM_EXEC := $(shell go env GOROOT)/lib/wasm/wasm_exec.js
GAME_LANG ?= ru

.PHONY: help desktop wasm serve serve-build browser test clean

help:
	@echo "Available commands:"
	@echo "  make desktop     - run the desktop app"
	@echo "  make wasm        - build browser wasm files"
	@echo "  make serve       - serve existing web build and open it in a browser"
	@echo "  make serve-build - rebuild browser wasm, then serve and open it"
	@echo "  make browser     - rebuild browser wasm, then serve and open it"
	@echo "  make test        - run Go tests"
	@echo "  make clean       - remove generated outputs"

desktop:
	go run ./cmd/paper-quarters -lang $(GAME_LANG)

wasm:
	GOOS=js GOARCH=wasm go build -o $(WASM) ./cmd/paper-quarters
	cp "$(WASM_EXEC)" "$(WEB_DIR)/wasm_exec.js"

serve:
	go run ./cmd/serve -lang $(GAME_LANG)

serve-build:
	go run ./cmd/serve --build -lang $(GAME_LANG)

browser: serve-build

test:
	go test ./...

clean:
	rm -f "$(WASM)" "$(WEB_DIR)/wasm_exec.js" "$(APP_NAME)" "$(APP_NAME).exe" serve serve.exe
