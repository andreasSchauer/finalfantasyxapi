#!/bin/bash
set -e  # Exit on any error

echo "Making one down migration..."

cd sql/schema
goose postgres "postgres://andreasschauer:@localhost:5432/ffxapi" down
cd ..
cd ..

echo "Migrated successfully!"