FROM golang:1.17.0-stretch

WORKDIR /app/src

COPY . /app/src

RUN ls

RUN go mod tidy

ENTRYPOINT [ "go", "run", "main.go" ]