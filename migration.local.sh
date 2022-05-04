#!/bin/bash

export MIGRATION_DIR=./migrations
export DB_HOST="localhost"
export DB_PORT="5444"
export DB_NAME="links"
export DB_USER="links"
export DB_PASSWORD="jw8s0F4"
export DB_SSL=disable

export PG_DSN="host=${DB_HOST} port=${DB_PORT} dbname=${DB_NAME} user=${DB_USER} password=${DB_PASSWORD} sslmode=${DB_SSL}"

goose -dir ${MIGRATION_DIR} postgres "${PG_DSN}" up -v