FROM golang:1.17.7

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o service cmd/main.go

CMD [ "./service" ]