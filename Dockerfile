FROM golang:1.24-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o cyclists .

FROM alpine:3.20

RUN apk add --no-cache ca-certificates
WORKDIR /app

COPY --from=build /app/cyclists ./cyclists
COPY --from=build /app/src/config/config.example.json ./src/config/config.example.json

EXPOSE 3000

CMD ["./cyclists"]
