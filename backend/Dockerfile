# Imagen base con Go
FROM golang:latest

ADD . /backend

WORKDIR /backend

# Instalar dependencias de Go
RUN go mod init github.com/aaraya0/ingsw3-final/backend
RUN go mod tidy
RUN go mod download

# Compilar la aplicacións
RUN go build -o main .

# Exponer el puerto en el que la aplicación escucha
EXPOSE 8090


# Comando para ejecutar la aplicación
CMD ["./main"]


