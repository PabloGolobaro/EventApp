FROM golang:1.18 AS development
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go get github.com/codegangsta/gin
RUN go install github.com/codegangsta/gin
EXPOSE 8000

CMD gin -i --appPort 8080 --port 3000 --bin notify_server --path ./cmd/notify_server run ./cmd/notify_server/main.go

FROM golang:1.18 as build

COPY . /src

WORKDIR /src

RUN CGO_ENABLED=0 GOOS=linux go build -o notify_server ./cmd/notify_server

FROM gcr.io/distroless/static-debian11 AS production

COPY --from=build /src/notify_server .

COPY --from=build /src/static ./static

COPY --from=build /src/app.env .

EXPOSE 80

CMD ["/notify_server"]