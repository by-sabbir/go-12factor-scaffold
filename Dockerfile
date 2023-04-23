FROM golang:latest AS build

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -o srv main.go

# FROM gcr.io/distroless/static-debian11
FROM alpine

WORKDIR /app

COPY --from=build /app/srv .
COPY --from=build /app/migrations /app/migrations/
CMD [ "./srv", "server"]
