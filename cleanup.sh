#!/bin/bash

# CLI Build
rm -rf runtipi-cli-go
rm -rf runtipi-cli-go.bak

# CLI Created Files
rm -rf apps
rm -rf data
rm -rf app-data
rm -rf state
rm -rf repos
rm -rf media
rm -rf traefik
rm -rf user-config
rm -rf backups
rm -rf logs
rm -rf docker-compose.yml
rm -rf VERSION
rm -rf .env

# Test Files
# /internal/constants/assets/RUNTIPI_VERSION
# /internal/constants/assets/CLI_VERSION
# /internal/constants/assets/VERSION
# /.env.local