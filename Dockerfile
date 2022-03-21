FROM golang:1.17.0-stretch

RUN mkdir -p  "app/src"

WORKDIR /app/src

COPY ./ /app/src

RUN go mod tidy

ENTRYPOINT [ "go", "run", "main.go" ]