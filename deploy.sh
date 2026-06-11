#!/bin/bash

# Ensure script stops on first error
set -e

echo "======================================"
echo " Starting ShipTrack Deployment..."
echo "======================================"

# 1. Pull latest code from main branch
echo "[1/4] Pulling latest code..."
git pull origin master

# 2. Rebuild and restart Docker containers in detached mode
echo "[2/4] Rebuilding and restarting Docker containers..."
docker compose up -d --build

# 3. Remove unused Docker images to free up VPS space
echo "[3/4] Cleaning up old Docker images..."
docker image prune -f

echo "======================================"
echo " Deployment completed successfully!   "
echo "======================================"
