#!/bin/bash
set -e  # Exit on any error

echo "Resetting database..."

cd sql/schema
goose postgres "postgres://andreasschauer:@localhost:5432/ffxapi" down-to 0
goose postgres "postgres://andreasschauer:@localhost:5432/ffxapi" up
cd ..
cd ..

echo "Migrated successfully!"