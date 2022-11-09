#!/usr/bin/env bash

if ! command -v docker &> /dev/null
then
	echo 'Docker is required for this script. Please install it at https://docs.docker.com/get-docker/';
	exit 1
fi;

DB_USER=postgres
DB_PASS=passw0rd!
DB_NAME=postgres

echo 'Making sure database name is available...';

docker rm -f lovelygames_db

echo 'Starting database...';

docker run --name lovelygames_db --rm -e POSTGRES_USER="${DB_USER}" -e POSTGRES_DB="${DB_NAME}" -e POSTGRES_PASSWORD="${DB_PASS}" -p 0.0.0.0:5432:5432 -d postgres

echo 'Waiting for database to be ready...';

sleep 10

echo 'Setting up database schema...';
docker cp ./schema.sql lovelygames_db:/tmp/schema.sql
docker exec -u "${DB_USER}" lovelygames_db psql "${DB_NAME}" "${DB_USER}" -f /tmp/schema.sql

echo 'Setting up database fixtures...';
docker cp ./fixtures.sql lovelygames_db:/tmp/fixtures.sql
docker exec -u "${DB_USER}" lovelygames_db psql "${DB_NAME}" "${DB_USER}" -f /tmp/fixtures.sql