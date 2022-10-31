FROM golang:1.17.7-alpine3.15 AS builder
WORKDIR /app
COPY . .

## install psql
#RUN apt-get update
#RUN apt-get -y install postgresql-client
#
## make wait-for-postgres.sh executable
#RUN chmod +x wait-for-postgres.sh

RUN CGO_ENABLED=0 GOOS=linux go build -o build/app cmd/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder ./app .
CMD ["build/app"]