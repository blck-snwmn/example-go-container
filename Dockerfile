FROM golang:1.22.3-bullseye as build

WORKDIR /app

RUN go env -w GOMODCACHE=/root/.cache/go-build

COPY go.mod go.sum ./
RUN --mount=type=cache,target="/root/.cache/go-build" \ 
    go mod download

COPY . .

FROM build as deploy0
RUN --mount=type=cache,target="/root/.cache/go-build" \
    go build -trimpath -ldflags "-w -s" -o app


FROM build as deploy1
RUN --mount=type=cache,target="/root/.cache/go-build" \
    go build -trimpath -ldflags "-w -s" -o /cmd/cmd1/cmd1 ./cmd/cmd1


FROM build as deploy2
RUN --mount=type=cache,target="/root/.cache/go-build" \
    go build -trimpath -ldflags "-w -s" -o /cmd/cmd2/cmd2 ./cmd/cmd2


FROM debian:bullseye-slim as deploy

RUN apt-get update

COPY --from=deploy0 /app/app .
COPY --from=deploy1 /cmd/cmd1/cmd1 .
COPY --from=deploy2 /cmd/cmd2/cmd2 .

CMD ["./app"]