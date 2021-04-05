#!/bin/sh
set -e

psql -v ON_ERROR_STOP=1 \
  --host="$POSTGRES_HOST" \
  --port="$POSTGRES_PORT" \
  --username="$POSTGRES_USER" \
  --dbname="$POSTGRES_DB" \
  -f migrations/load.sql
