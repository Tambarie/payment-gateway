FROM golang:1.17.0-stretch

WORKDIR /app

COPY . /app

RUN go mod tidy

CMD [ "go", "run", "main.go" ]