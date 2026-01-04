# syntax=docker/dockerfile:1
FROM golang:1.22-alpine AS build

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
# No CGO required (pure-go sqlite), keep it disabled for reproducible builds.
ENV CGO_ENABLED=0
RUN go build -o server .

FROM alpine:3.20
WORKDIR /app
COPY --from=build /app/server /app/server
COPY public /app/public

# Persist sqlite file (when using sqlite)
VOLUME ["/app/data"]
ENV DB_DIALECT=sqlite
ENV SQLITE_PATH=/app/data/workshop.db

EXPOSE 3000
CMD ["/app/server"]
