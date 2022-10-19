FROM golang:1.17.7

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

## install psql
#RUN apt-get update
#RUN apt-get -y install postgresql-client
#
## make wait-for-postgres.sh executable
#RUN chmod +x wait-for-postgres.sh

RUN go build -o service cmd/main.go

CMD [ "./service" ]