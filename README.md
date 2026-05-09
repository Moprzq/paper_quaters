# Paper Quarters

Ebitengine-based card dealer helper for "Welcome To" / "Bumazhnye kvartaly".

The app embeds lightweight card images from `internal/assets/cards/` and
`internal/assets/missions/`. The heavy original scans were removed; the
checked-in images are already resized and compressed for desktop and browser
builds.

## Desktop

From the repository root:

```powershell
make desktop
```

If `make` is not installed:

```powershell
rtk go run ./cmd/paper-quarters
```

## Browser

From the repository root:

```powershell
make browser
```

If `make` is not installed:

```powershell
rtk powershell -NoLogo -NoProfile -NonInteractive -Command "`$env:GOOS='js'; `$env:GOARCH='wasm'; go build -o web\paper_quarters.wasm ./cmd/paper-quarters; Remove-Item Env:GOOS; Remove-Item Env:GOARCH"
rtk powershell -NoLogo -NoProfile -NonInteractive -Command "Copy-Item (Join-Path (go env GOROOT) 'lib\wasm\wasm_exec.js') web\wasm_exec.js"
rtk go run ./cmd/serve
```

Then open:

```text
http://localhost:8080/
```

## Commands

```text
make desktop   run the desktop app
make wasm      build browser wasm files
make serve     serve web/ at http://localhost:8080/
make browser   build wasm, then serve web/
make test      run Go tests
make clean     remove generated outputs
```

The `Makefile` uses standard POSIX commands and should work on Ubuntu and
macOS. On Windows, use Git Bash/MSYS2/WSL for `make`, or use the fallback
PowerShell commands above.

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
