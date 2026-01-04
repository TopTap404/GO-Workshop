# ---------- build stage ----------
FROM golang:1.24 AS build
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server .

# ---------- runtime stage ----------
FROM alpine:3.20
WORKDIR /app

COPY --from=build /app/server /app/server
COPY public /app/public

EXPOSE 3000
CMD ["/app/server"]
