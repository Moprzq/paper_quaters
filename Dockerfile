# syntax=docker/dockerfile:1

FROM golang:1.26.3-bookworm AS build

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ARG GOOS=windows
ARG GOARCH=amd64

RUN CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} go build -o /out/desktop/paper_quarters.exe ./cmd/paper-quarters
RUN CGO_ENABLED=0 GOOS=js GOARCH=wasm go build -o /out/web/paper_quarters.wasm ./cmd/paper-quarters
RUN cp "$(go env GOROOT)/lib/wasm/wasm_exec.js" /out/web/wasm_exec.js
RUN cp web/index.html /out/web/index.html

FROM scratch AS export

COPY --from=build /out/ /
