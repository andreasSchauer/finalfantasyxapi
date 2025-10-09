#!/bin/bash
set -e  # Exit on any error

VAL="$1"
echo "Migrating to version $1..."


cd sql/schema
goose postgres "postgres://andreasschauer:@localhost:5432/ffxapi" down-to $VAL
cd ..
cd ..

echo "Migrated successfully!"