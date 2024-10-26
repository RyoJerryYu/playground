#!/bin/bash
cd $(dirname "$0")
echo "Updating dependencies..."
buf build
buf dep update
echo "Done"
