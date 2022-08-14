FROM golang:1.18 as build

COPY . /src

WORKDIR /src

RUN CGO_ENABLED=0 GOOS=linux go build -o notifier ./cmd/notify_bot

FROM gcr.io/distroless/static-debian11

COPY --from=build /src/notifier .

COPY --from=build /src/Birthdays.xlsx .

EXPOSE 8080

CMD ["/notifier"]