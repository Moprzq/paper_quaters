# Paper Quarters

Card dealer helper for **Welcome To** / **Bumazhnye kvartaly**.

The project is now an Ebitengine app. It can run as a desktop window or as a
browser WebAssembly build. The old terminal version was removed.

## Features

- Shuffles the card deck into three stacks.
- Shows the current three card/action pairs.
- Shows three random missions.
- Lets you go forward/back, reshuffle, restart, and mark missions as completed.
- Embeds lightweight card images, so the desktop binary does not need external
  image folders next to it.

## Project Layout

```text
cmd/
  paper-quarters/   desktop/browser app entrypoint
  serve/            local HTTP server for the browser build

internal/
  app/              game state, deck, missions, UI, Ebitengine loop
  assets/           embedded optimized card and mission images

web/
  index.html        browser page
```

There are two `main` packages because the repository builds two separate
commands: the app itself and the local web server.

## Requirements

- Go `1.26.3`
- `make` for the short commands on Ubuntu/macOS

On Windows, use Git Bash/MSYS2/WSL for `make`, or use the PowerShell fallback
commands below.

## Desktop

With `make`:

```bash
make desktop
```

Without `make`:

```powershell
go run ./cmd/paper-quarters
```

## Browser

With `make`:

```bash
make browser
```

Then open:

```text
http://localhost:8080/
```

Without `make` on Windows:

```powershell
$env:GOOS='js'
$env:GOARCH='wasm'
go build -o web\paper_quarters.wasm ./cmd/paper-quarters
Remove-Item Env:GOOS
Remove-Item Env:GOARCH

$goroot = go env GOROOT
Copy-Item $goroot\lib\wasm\wasm_exec.js web\wasm_exec.js

go run ./cmd/serve
```

Then open:

```text
http://localhost:8080/
```

## Make Commands

```text
make desktop   run the desktop app
make wasm      build browser wasm files
make serve     serve web/ at http://localhost:8080/
make browser   build wasm, then serve web/
make test      run Go tests
make clean     remove generated outputs
```

## Docker

Docker can produce both desktop and browser artifacts:

```powershell
docker build --output type=local,dest=dist .
```

The output will contain:

- `desktop/paper_quarters.exe`
- `web/index.html`
- `web/paper_quarters.wasm`
- `web/wasm_exec.js`

## Notes

- Source card images live under `internal/assets/`.
- The images are already resized/compressed; there is no asset generation step.
- The UI text size is controlled by `controlFontSize` in
  `internal/app/game.go`.
