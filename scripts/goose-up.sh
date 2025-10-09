#!/bin/bash
set -e  # Exit on any error

echo "Making up migration..."

cd sql/schema
goose postgres "postgres://andreasschauer:@localhost:5432/ffxapi" up
cd ..
cd ..

echo "Migrated successfully!"