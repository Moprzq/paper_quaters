# syntax=docker/dockerfile:1

FROM golang:1.26.3-bookworm AS build

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/app/serve ./cmd/serve
RUN CGO_ENABLED=0 GOOS=js GOARCH=wasm go build -o /out/app/web/paper_quarters.wasm ./cmd/paper-quarters
RUN cp "$(go env GOROOT)/lib/wasm/wasm_exec.js" /out/app/web/wasm_exec.js
RUN cp web/index.html /out/app/web/index.html

FROM debian:bookworm-slim AS runtime

WORKDIR /app

COPY --from=build /out/app/ /app/

EXPOSE 8080

CMD ["/app/serve", "-addr", "0.0.0.0:8080", "-open=false", "-dir", "/app/web"]
