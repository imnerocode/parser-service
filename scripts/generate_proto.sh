#!/bin/bash

# Bash script to generate Go files from proto using namely/protoc-all

# Default parameters
PROTO_FILE="parser.proto"
PROTO_PATH="./proto"
OUTPUT_PATH="./proto/generated"

# Allow overriding parameters via command-line arguments
while getopts "f:p:o:" opt; do
  case $opt in
    f) PROTO_FILE=$OPTARG ;;
    p) PROTO_PATH=$OPTARG ;;
    o) OUTPUT_PATH=$OPTARG ;;
    *) echo "Usage: $0 [-f proto_file] [-p proto_path] [-o output_path]" && exit 1 ;;
  esac
done

# Check if Docker is installed
if ! command -v docker &> /dev/null; then
  echo "Error: Docker is not installed. Please install Docker and try again."
  exit 1
fi

# Ensure the output directory exists
mkdir -p "$OUTPUT_PATH"

# Run the Docker command
if docker run --rm -v "$(pwd)/${PROTO_PATH}:/protos" -v "$(pwd)/${OUTPUT_PATH}:/out" namely/protoc-all:1.44_1 \
    -f "${PROTO_FILE}" \
    -l go \
    -o /out \
    --with-grpc; then
  echo "Protobuf files generated successfully in $OUTPUT_PATH"
else
  echo "Error: Failed to generate protobuf files."
  exit 1
fi
