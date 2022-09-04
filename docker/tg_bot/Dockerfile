FROM golang:1.18 AS development
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go install github.com/cespare/reflex@latest

CMD reflex -g '*.go' go run ./cmd/tg_bot/main.go --start-service

FROM golang:1.18 as build

COPY . /src

WORKDIR /src

RUN CGO_ENABLED=0 GOOS=linux go build -o tg_bot ./cmd/tg_bot

FROM gcr.io/distroless/static-debian11 AS production

COPY --from=build /src/tg_bot .

COPY --from=build /src/config ./config

CMD ["/tg_bot"]