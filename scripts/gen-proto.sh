#!/bin/bash

# Set the current directory to the script's directory
cd "$(dirname "${BASH_SOURCE[0]}")"

# Paths to the Protobuf files (assuming the script is in the "scripts" directory)
PROTO_FILES=("../proto/collector/collector.proto" "../proto/post/post.proto")

# Output directory for generated Go files
OUTPUT_DIR_1="../internal/collector/port/genproto"
OUTPUT_DIR_2="../internal/post/port/genproto"
OUTPUT_DIR_3="../internal/apigateway/genproto"

# Ensure the output directories exist
mkdir -p "$OUTPUT_DIR_1"
mkdir -p "$OUTPUT_DIR_2"
mkdir -p "$OUTPUT_DIR_3"

# Generate Go files using protoc for each Protobuf file
for PROTO_FILE in "${PROTO_FILES[@]}"; do
    # Get the directory containing the Protobuf file
    PROTO_DIR=$(dirname "$PROTO_FILE")

    # Generate Go files using protoc with specified proto_path for OUTPUT_DIR_1
    protoc --proto_path="../proto" --go_out="$OUTPUT_DIR_1" --go-grpc_out="$OUTPUT_DIR_1" "$PROTO_DIR"/*.proto

    # Generate Go files using protoc with specified proto_path for OUTPUT_DIR_2
    protoc --proto_path="../proto" --go_out="$OUTPUT_DIR_2" --go-grpc_out="$OUTPUT_DIR_2" "$PROTO_DIR"/*.proto

     # Generate Go files using protoc with specified proto_path for OUTPUT_DIR_2
    protoc --proto_path="../proto" --go_out="$OUTPUT_DIR_3" --go-grpc_out="$OUTPUT_DIR_3" "$PROTO_DIR"/*.proto
done