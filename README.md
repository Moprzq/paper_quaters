# Paper Quarters

Card dealer helper for **Welcome To** / **Бумажные кварталы**. Runs as a
desktop Ebitengine app or as a browser WebAssembly app.

## Features

- Shows three card/action pairs and three missions.
- Supports moving forward/back through turns, shuffling, restarting, and marking
  missions completed.
- Tracks built houses. The counter grows only on newly reached turns, survives
  shuffles, and resets on restart.
- Russian and English UI. Russian is the default.

## Controls

```text
Space / Right arrow   next turn
Left arrow            back
S                     shuffle deck
R                     restart game
1 / 2 / 3             mark mission completed
F11                   toggle fullscreen
Q / Esc               exit
```

## Run

Desktop:

```bash
make desktop
```

Browser:

```bash
make browser
```

Docker:

```bash
docker build -t paper-quarters .
docker run --rm -p 8080:8080 paper-quarters
```

Then open `http://localhost:8080/` in a browser. English UI is available at
`http://localhost:8080/?lang=eng`.

English UI:

```bash
make desktop GAME_LANG=eng
make browser GAME_LANG=eng
```

Without `make`:

```powershell
go run ./cmd/paper-quarters
go run ./cmd/paper-quarters -lang eng

go run ./cmd/serve --build
go run ./cmd/serve --build -lang eng
```

## Serve

`serve` opens the game in the default browser and serves `web/` with no-cache
headers.

```powershell
go run ./cmd/serve              # use existing build, build if missing
go run ./cmd/serve --build      # force rebuild
go run ./cmd/serve -open=false  # do not open browser
go run ./cmd/serve -addr localhost:18080
```

English browser UI uses `?lang=eng`. Without it, the UI is Russian.

## Make Commands

```text
make desktop       run the desktop app
make wasm          build browser wasm files
make serve         serve existing web build and open it in a browser
make serve-build   rebuild browser wasm, then serve and open it
make browser       rebuild browser wasm, then serve and open it
make test          run Go tests
make clean         remove generated outputs
```

## Project Layout

```text
cmd/paper-quarters/   app entrypoint
cmd/serve/            local browser server
internal/app/         game state, UI, localization
internal/assets/      embedded card and mission images
web/                  browser page and generated wasm files
```

## Requirements

- Go `1.26.3`
- `make` optional
