FROM golang:1.17.7-alpine3.15 AS builder
WORKDIR /app
COPY . .

## install psql
#RUN apt-get update
#RUN apt-get -y install postgresql-client
#
## make wait-for-postgres.sh executable
#RUN chmod +x wait-for-postgres.sh

ARG CGO_ENABLED=0
ARG GOOS=linux
RUN go build -o build/main cmd/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder ./app/config.yml .
COPY --from=builder ./app/build .
COPY --from=builder ./app/migrations ./migrations
CMD ["./main"]