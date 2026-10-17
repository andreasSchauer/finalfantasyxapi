#!/bin/bash
set -e  # Exit on any error

echo "Migrating to latest version..."

cd sql/schema
goose postgres "postgres://andreasschauer:@localhost:5432/ffxapi" up
cd ..
cd ..

echo "Migrated successfully!"