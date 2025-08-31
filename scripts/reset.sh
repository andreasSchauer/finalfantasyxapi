#!/bin/bash
set -e

echo "Resetting database..."

source .env
curl -X POST http://localhost:8080/admin/reset \
  -H "X-Admin-Key: $ADMIN_API_KEY"

echo "Database reset successfully!"