#!/bin/bash
cd $(dirname "$0")
echo "Generating code..."
buf generate
echo "Done"
