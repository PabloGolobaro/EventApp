FROM golang:1.18 AS development
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go install github.com/cespare/reflex@latest
EXPOSE 8080
CMD reflex -g '*.go' go run ./cmd/notify_bot/main.go --start-service

FROM golang:1.18 as build

COPY . /src

WORKDIR /src

RUN CGO_ENABLED=0 GOOS=linux go build -o notifier ./cmd/notify_bot

FROM gcr.io/distroless/static-debian11 AS production

COPY --from=build /src/notifier .

COPY --from=build /src/Birthdays.xlsx .

EXPOSE 8080

CMD ["/notifier"]