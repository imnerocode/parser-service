#!/bin/bash

# Test script for generate_proto.sh

SCRIPT_PATH="../scripts/generate_proto.sh"

# Test cases
echo "Running tests for generate_proto.sh..."

# Test 1: Verify script execution with default parameters
if bash "$SCRIPT_PATH"; then
  echo "Test 1 passed: Script executed successfully with default parameters."
else
  echo "Test 1 failed: Script failed with default parameters."
  exit 1
fi

# Test 2: Verify custom parameters
CUSTOM_OUTPUT="./tests/generated"
if bash "$SCRIPT_PATH" -f "parser.proto" -p "./proto" -o "$CUSTOM_OUTPUT"; then
  echo "Test 2 passed: Script executed successfully with custom parameters."
  if [ -d "$CUSTOM_OUTPUT" ]; then
    echo "Custom output directory verified: $CUSTOM_OUTPUT"
    rm -rf "$CUSTOM_OUTPUT"
  else
    echo "Test 2 failed: Custom output directory not found."
    exit 1
  fi
else
  echo "Test 2 failed: Script failed with custom parameters."
  exit 1
fi

echo "All tests for generate_proto.sh passed."
