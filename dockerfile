FROM golang:1.21

RUN go install github.com/cosmtrek/air@latest

RUN git config --global --add safe.directory /app 

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

CMD ["air", "-c", ".air.toml"]
