FROM golang:1.22.3-bullseye as build

WORKDIR /app

RUN go env -w GOMODCACHE=/root/.cache/go-build

COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod go mod download

RUN ls -al /go
RUN ls -al /root

RUN ls -al /go/pkg
RUN ls -al /go/pkg/mod
RUN ls -al /root/.cache

COPY . .
RUN --mount=type=cache,target=/go/pkg/mod \ 
    --mount=type=cache,target="/root/.cache/go-build" \
    go build -trimpath -ldflags "-w -s" -o app

FROM debian:bullseye-slim as deploy

RUN apt-get update

COPY --from=build /app/app .

CMD ["./app"]