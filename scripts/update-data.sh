#!/bin/bash
set -e  # Exit on any error

echo "Updating game data..."

cd data
git pull origin main
cd ..
git add data
git commit -m "Update data to latest version"

echo "Data updated successfully!"