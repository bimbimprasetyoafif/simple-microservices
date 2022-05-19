FROM golang:1.18 as builder

WORKDIR /app

COPY . .

RUN go build -tags netgo -o main.app .

CMD ["/app/main.app"]

# TODO : create runner instead of run in builder