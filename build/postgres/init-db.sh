#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
	CREATE USER faculus ENCRYPTED PASSWORD 'faculus' LOGIN;
	CREATE DATABASE faculus OWNER faculus;
EOSQL

psql -v ON_ERROR_STOP=1 --username "faculus" --dbname "faculus" -f /app/sql/init-db.sql
