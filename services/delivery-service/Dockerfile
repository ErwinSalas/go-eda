# Etapa de construcción
FROM golang:1.19 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o delivery-service main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=build /app/delivery-service .

EXPOSE 8080

CMD ["./delivery-service"]
