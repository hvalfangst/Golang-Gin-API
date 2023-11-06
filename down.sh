#!/bin/sh

# Exits immediately if a command exits with a non-zero status
set -e

# Run 'docker-compose down' for removing our DB container
docker-compose -f db/cars/docker-compose.yml down