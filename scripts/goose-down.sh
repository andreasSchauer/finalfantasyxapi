#!/bin/bash
set -e  # Exit on any error


echo "Migrating to version 0..."


cd sql/schema
goose postgres "postgres://andreasschauer:@localhost:5432/ffxapi" down-to 0
cd ..
cd ..

echo "Migrated successfully!"