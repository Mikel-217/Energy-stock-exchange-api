FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

# TODO: Add dockerignore
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /energy-stock-exchange-api

EXPOSE 8080
CMD [ "/energy-stock-exchange-api" ]
