# Etapa de construcción
FROM golang:1.19 AS build

# Crear y movernos al directorio de la app
WORKDIR /app

# Copiar los archivos go.mod y go.sum para instalar dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copiar el código fuente del servicio
COPY . .

# Compilar la aplicación Go en un binario
RUN CGO_ENABLED=0 GOOS=linux go build -o order-service main.go

# Imagen mínima para ejecutar la app (alpine)
FROM alpine:latest

# Crear un directorio en el contenedor para la app
WORKDIR /root/

# Copiar el binario desde la etapa de construcción
COPY --from=build /app/order-service .

# Exponer el puerto en el que correrá el servicio
EXPOSE 8080

# Comando que ejecuta la app cuando el contenedor se inicia
CMD ["./order-service"]
