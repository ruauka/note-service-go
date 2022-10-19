#!/bin/sh
# wait-for-postgres.sh

set -e

host="$1"
shift
cmd="$@"

until psql postgresql://pg:pass@database:5432/crud; do
  >&2 echo "Postgres is unavailable - sleeping"
  sleep 1
done

migrate -path ./migrate -database 'postgres://pg:pass@database:5432/crud?sslmode=disable' up

>&2 echo "Postgres is up - executing command"
exec $cmd